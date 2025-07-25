// Copyright 2024 The Erigon Authors
// This file is part of Erigon.
//
// Erigon is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Erigon is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with Erigon. If not, see <http://www.gnu.org/licenses/>.

package jsonrpc

import (
	"context"
	"fmt"
	"math/big"
	"testing"
	"time"

	"github.com/holiman/uint256"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/erigontech/erigon-db/rawdb"
	"github.com/erigontech/erigon-lib/chain"
	"github.com/erigontech/erigon-lib/common"
	"github.com/erigontech/erigon-lib/common/hexutil"
	"github.com/erigontech/erigon-lib/crypto"
	txpool "github.com/erigontech/erigon-lib/gointerfaces/txpoolproto"
	"github.com/erigontech/erigon-lib/kv"
	"github.com/erigontech/erigon-lib/kv/kvcache"
	"github.com/erigontech/erigon-lib/kv/rawdbv3"
	"github.com/erigontech/erigon-lib/log/v3"
	"github.com/erigontech/erigon-lib/trie"
	"github.com/erigontech/erigon-lib/types"
	"github.com/erigontech/erigon/cmd/rpcdaemon/rpcdaemontest"
	"github.com/erigontech/erigon/core"
	"github.com/erigontech/erigon/core/state"
	"github.com/erigontech/erigon/eth/ethconfig"
	"github.com/erigontech/erigon/execution/stages/mock"
	"github.com/erigontech/erigon/rpc"
	"github.com/erigontech/erigon/rpc/ethapi"
	"github.com/erigontech/erigon/rpc/rpccfg"
	"github.com/erigontech/erigon/rpc/rpchelper"
)

func TestEstimateGas(t *testing.T) {
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	stateCache := kvcache.New(kvcache.DefaultCoherentConfig)
	ctx, conn := rpcdaemontest.CreateTestGrpcConn(t, mock.Mock(t))
	mining := txpool.NewMiningClient(conn)
	ff := rpchelper.New(ctx, rpchelper.DefaultFiltersConfig, nil, nil, mining, func() {}, m.Log)
	api := NewEthAPI(NewBaseApi(ff, stateCache, m.BlockReader, false, rpccfg.DefaultEvmCallTimeout, m.Engine, m.Dirs, nil), m.DB, nil, nil, nil, 5000000, ethconfig.Defaults.RPCTxFeeCap, 100_000, false, 100_000, 128, log.New())
	var from = common.HexToAddress("0x71562b71999873db5b286df957af199ec94617f7")
	var to = common.HexToAddress("0x0d3ab14bbad3d99f4203bd7a11acb94882050e7e")
	if _, err := api.EstimateGas(context.Background(), &ethapi.CallArgs{
		From: &from,
		To:   &to,
	}, nil, nil); err != nil {
		t.Errorf("calling EstimateGas: %v", err)
	}
}

func TestEthCallNonCanonical(t *testing.T) {
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	stateCache := kvcache.New(kvcache.DefaultCoherentConfig)
	api := NewEthAPI(NewBaseApi(nil, stateCache, m.BlockReader, false, rpccfg.DefaultEvmCallTimeout, m.Engine, m.Dirs, nil), m.DB, nil, nil, nil, 5000000, ethconfig.Defaults.RPCTxFeeCap, 100_000, false, 100_000, 128, log.New())
	var from = common.HexToAddress("0x71562b71999873db5b286df957af199ec94617f7")
	var to = common.HexToAddress("0x0d3ab14bbad3d99f4203bd7a11acb94882050e7e")
	if _, err := api.Call(context.Background(), ethapi.CallArgs{
		From: &from,
		To:   &to,
	}, rpc.BlockNumberOrHashWithHash(common.HexToHash("0x3fcb7c0d4569fddc89cbea54b42f163e0c789351d98810a513895ab44b47020b"), true), nil); err != nil {
		if fmt.Sprintf("%v", err) != "hash 3fcb7c0d4569fddc89cbea54b42f163e0c789351d98810a513895ab44b47020b is not currently canonical" {
			t.Errorf("wrong error: %v", err)
		}
	}
}

func TestEthCallToPrunedBlock(t *testing.T) {
	pruneTo := uint64(3)
	ethCallBlockNumber := rpc.BlockNumber(2)

	m, bankAddress, contractAddress := chainWithDeployedContract(t)
	doPrune(t, m.DB, pruneTo)
	api := NewEthAPI(newBaseApiForTest(m), m.DB, nil, nil, nil, 5000000, ethconfig.Defaults.RPCTxFeeCap, 100_000, false, 100_000, 128, log.New())

	callData := hexutil.MustDecode("0x2e64cec1")
	callDataBytes := hexutil.Bytes(callData)

	if _, err := api.Call(context.Background(), ethapi.CallArgs{
		From: &bankAddress,
		To:   &contractAddress,
		Data: &callDataBytes,
	}, rpc.BlockNumberOrHashWithNumber(ethCallBlockNumber), nil); err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetProof(t *testing.T) {
	var maxGetProofRewindBlockCount = 1 // Note, this is unsafe for parallel tests, but, this test is the only consumer for now

	m, bankAddr, contractAddr := chainWithDeployedContract(t)
	api := NewEthAPI(newBaseApiForTest(m), m.DB, nil, nil, nil, 5000000, ethconfig.Defaults.RPCTxFeeCap, 100_000, false, maxGetProofRewindBlockCount, 128, log.New())

	key := func(b byte) hexutil.Bytes {
		result := common.Hash{}
		result[31] = b
		return result.Bytes()
	}

	tests := []struct {
		name        string
		blockNum    uint64
		addr        common.Address
		storageKeys []hexutil.Bytes
		stateVal    uint64
		expectedErr string
	}{
		{
			name:     "currentBlockNoState",
			addr:     contractAddr,
			blockNum: 3,
		},
		{
			name:     "currentBlockEOA",
			addr:     bankAddr,
			blockNum: 3,
		},
		{
			name:     "currentBlockNoAccount",
			addr:     common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddead0"),
			blockNum: 3,
		},
		{
			name:        "currentBlockWithState",
			addr:        contractAddr,
			blockNum:    3,
			storageKeys: []hexutil.Bytes{key(0), key(4), key(8), key(10)},
			stateVal:    2,
		},
		{
			name:        "currentBlockWithMissingState",
			addr:        contractAddr,
			storageKeys: []hexutil.Bytes{hexutil.FromHex("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead")},
			blockNum:    3,
			stateVal:    0,
		},
		{
			name:        "currentBlockEOAMissingState",
			addr:        bankAddr,
			storageKeys: []hexutil.Bytes{hexutil.FromHex("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead")},
			blockNum:    3,
			stateVal:    0,
		},
		{
			name:        "currentBlockNoAccountMissingState",
			addr:        common.HexToAddress("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddead0"),
			storageKeys: []hexutil.Bytes{hexutil.FromHex("0xdeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddeaddead")},
			blockNum:    3,
			stateVal:    0,
		},
		// {
		// 	name:        "olderBlockWithState",
		// 	addr:        contractAddr,
		// 	blockNum:    2,
		// 	storageKeys: []common.Hash{key(1), key(5), key(9), key(13)},
		// 	stateVal:    1,
		// },
		// {
		// 	name:        "tooOldBlock",
		// 	addr:        contractAddr,
		// 	blockNum:    1,
		// 	expectedErr: "requested block is too old, block must be within 1 blocks of the head block number (currently 3)",
		// },
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			proof, err := api.GetProof(
				context.Background(),
				tt.addr,
				tt.storageKeys,
				rpc.BlockNumberOrHashWithNumber(rpc.BlockNumber(tt.blockNum)),
			)
			if tt.expectedErr != "" {
				require.EqualError(t, err, tt.expectedErr)
				require.Nil(t, proof)
				return
			}
			require.NoError(t, err)
			require.NotNil(t, proof)

			tx, err := m.DB.BeginTemporalRo(context.Background())
			require.NoError(t, err)
			defer tx.Rollback()
			header, err := api.headerByRPCNumber(context.Background(), rpc.BlockNumber(tt.blockNum), tx)
			require.NoError(t, err)

			require.Equal(t, tt.addr, proof.Address)
			err = trie.VerifyAccountProof(header.Root, proof)
			require.NoError(t, err)

			require.Len(t, proof.StorageProof, len(tt.storageKeys))
			for _, storageKey := range tt.storageKeys {
				found := false
				for _, storageProof := range proof.StorageProof {
					var proofKeyHash, storageKeyHash common.Hash
					proofKeyHash.SetBytes(hexutil.FromHex(storageProof.Key))
					storageKeyHash.SetBytes(uint256.NewInt(0).SetBytes(storageKey).Bytes())
					if proofKeyHash != storageKeyHash {
						continue
					}
					found = true
					require.Equal(t, tt.stateVal, (*big.Int)(storageProof.Value).Uint64())
					err = trie.VerifyStorageProof(proof.StorageHash, storageProof)
					require.NoError(t, err)
				}
				require.True(t, found, "did not find storage proof for key=%x", storageKey)
			}
		})
	}
}

func TestGetBlockByTimestampLatestTime(t *testing.T) {
	ctx := context.Background()
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	tx, err := m.DB.BeginTemporalRo(ctx)
	if err != nil {
		t.Errorf("fail at beginning tx")
	}
	defer tx.Rollback()
	api := NewErigonAPI(newBaseApiForTest(m), m.DB, nil)

	latestBlock, err := m.BlockReader.CurrentBlock(tx)
	require.NoError(t, err)
	response, err := ethapi.RPCMarshalBlockDeprecated(latestBlock, true, false)

	if err != nil {
		t.Error("couldn't get the rpc marshal block")
	}

	if err == nil && rpc.BlockNumber(latestBlock.NumberU64()) == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}

	block, err := api.GetBlockByTimestamp(ctx, rpc.Timestamp(latestBlock.Header().Time), false)
	if err != nil {
		t.Errorf("couldn't retrieve block %v", err)
	}

	if block["timestamp"] != response["timestamp"] || block["hash"] != response["hash"] {
		t.Errorf("Retrieved the wrong block.\nexpected block hash: %s expected timestamp: %d\nblock hash retrieved: %s timestamp retrieved: %d", response["hash"], response["timestamp"], block["hash"], block["timestamp"])
	}
}

func TestGetBlockByTimestampOldestTime(t *testing.T) {
	ctx := context.Background()
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	tx, err := m.DB.BeginTemporalRo(ctx)
	if err != nil {
		t.Errorf("failed at beginning tx")
	}
	defer tx.Rollback()
	api := NewErigonAPI(newBaseApiForTest(m), m.DB, nil)

	oldestBlock, err := m.BlockReader.BlockByNumber(m.Ctx, tx, 0)
	if err != nil {
		t.Error("couldn't retrieve oldest block")
	}

	response, err := ethapi.RPCMarshalBlockDeprecated(oldestBlock, true, false)

	if err != nil {
		t.Error("couldn't get the rpc marshal block")
	}

	if err == nil && rpc.BlockNumber(oldestBlock.NumberU64()) == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}

	block, err := api.GetBlockByTimestamp(ctx, rpc.Timestamp(oldestBlock.Header().Time), false)
	if err != nil {
		t.Errorf("couldn't retrieve block %v", err)
	}

	if block["timestamp"] != response["timestamp"] || block["hash"] != response["hash"] {
		t.Errorf("Retrieved the wrong block.\nexpected block hash: %s expected timestamp: %d\nblock hash retrieved: %s timestamp retrieved: %d", response["hash"], response["timestamp"], block["hash"], block["timestamp"])
	}
}

func TestGetBlockByTimeHigherThanLatestBlock(t *testing.T) {
	ctx := context.Background()
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	tx, err := m.DB.BeginTemporalRo(ctx)
	if err != nil {
		t.Errorf("fail at beginning tx")
	}
	defer tx.Rollback()
	api := NewErigonAPI(newBaseApiForTest(m), m.DB, nil)

	latestBlock, err := m.BlockReader.CurrentBlock(tx)
	require.NoError(t, err)

	response, err := ethapi.RPCMarshalBlockDeprecated(latestBlock, true, false)

	if err != nil {
		t.Error("couldn't get the rpc marshal block")
	}

	if err == nil && rpc.BlockNumber(latestBlock.NumberU64()) == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}

	block, err := api.GetBlockByTimestamp(ctx, rpc.Timestamp(latestBlock.Header().Time+999999999999), false)
	if err != nil {
		t.Errorf("couldn't retrieve block %v", err)
	}

	if block["timestamp"] != response["timestamp"] || block["hash"] != response["hash"] {
		t.Errorf("Retrieved the wrong block.\nexpected block hash: %s expected timestamp: %d\nblock hash retrieved: %s timestamp retrieved: %d", response["hash"], response["timestamp"], block["hash"], block["timestamp"])
	}
}

func TestGetBlockByTimeMiddle(t *testing.T) {
	ctx := context.Background()
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	tx, err := m.DB.BeginTemporalRo(ctx)
	if err != nil {
		t.Errorf("fail at beginning tx")
	}
	defer tx.Rollback()
	api := NewErigonAPI(newBaseApiForTest(m), m.DB, nil)

	currentHeader := rawdb.ReadCurrentHeader(tx)
	oldestHeader, err := api._blockReader.HeaderByNumber(ctx, tx, 0)
	if err != nil {
		t.Errorf("error getting the oldest header %s", err)
	}
	if oldestHeader == nil {
		t.Error("couldn't find oldest header")
	}

	middleNumber := (currentHeader.Number.Uint64() + oldestHeader.Number.Uint64()) / 2
	middleBlock, err := m.BlockReader.BlockByNumber(m.Ctx, tx, middleNumber)
	if err != nil {
		t.Error("couldn't retrieve middle block")
	}

	response, err := ethapi.RPCMarshalBlockDeprecated(middleBlock, true, false)

	if err != nil {
		t.Error("couldn't get the rpc marshal block")
	}

	if err == nil && rpc.BlockNumber(middleBlock.NumberU64()) == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}

	block, err := api.GetBlockByTimestamp(ctx, rpc.Timestamp(middleBlock.Header().Time), false)
	if err != nil {
		t.Errorf("couldn't retrieve block %v", err)
	}
	if block["timestamp"] != response["timestamp"] || block["hash"] != response["hash"] {
		t.Errorf("Retrieved the wrong block.\nexpected block hash: %s expected timestamp: %d\nblock hash retrieved: %s timestamp retrieved: %d", response["hash"], response["timestamp"], block["hash"], block["timestamp"])
	}
}

func TestGetBlockByTimestamp(t *testing.T) {
	ctx := context.Background()
	m, _, _ := rpcdaemontest.CreateTestSentry(t)
	tx, err := m.DB.BeginTemporalRo(ctx)
	if err != nil {
		t.Errorf("fail at beginning tx")
	}
	defer tx.Rollback()
	api := NewErigonAPI(newBaseApiForTest(m), m.DB, nil)

	highestBlockNumber := rawdb.ReadCurrentHeader(tx).Number
	pickedBlock, err := m.BlockReader.BlockByNumber(m.Ctx, tx, highestBlockNumber.Uint64()/3)
	if err != nil {
		t.Errorf("couldn't get block %v", pickedBlock.Number())
	}

	if pickedBlock == nil {
		t.Error("couldn't retrieve picked block")
	}
	response, err := ethapi.RPCMarshalBlockDeprecated(pickedBlock, true, false)

	if err != nil {
		t.Error("couldn't get the rpc marshal block")
	}

	if err == nil && rpc.BlockNumber(pickedBlock.NumberU64()) == rpc.PendingBlockNumber {
		// Pending blocks need to nil out a few fields
		for _, field := range []string{"hash", "nonce", "miner"} {
			response[field] = nil
		}
	}

	block, err := api.GetBlockByTimestamp(ctx, rpc.Timestamp(pickedBlock.Header().Time), false)
	if err != nil {
		t.Errorf("couldn't retrieve block %v", err)
	}

	if block["timestamp"] != response["timestamp"] || block["hash"] != response["hash"] {
		t.Errorf("Retrieved the wrong block.\nexpected block hash: %s expected timestamp: %d\nblock hash retrieved: %s timestamp retrieved: %d", response["hash"], response["timestamp"], block["hash"], block["timestamp"])
	}
}

// contractHexString is the output of compiling the following solidity contract:
//
// pragma solidity ^0.8.0;
//
//	contract Box {
//	    uint256 private _value0x0;
//	    uint256 private _value0x1;
//	    uint256 private _value0x2;
//	    uint256 private _value0x3;
//	    uint256 private _value0x4;
//	    uint256 private _value0x5;
//	    uint256 private _value0x6;
//	    uint256 private _value0x7;
//	    uint256 private _value0x8;
//	    uint256 private _value0x9;
//	    uint256 private _value0xa;
//	    uint256 private _value0xb;
//	    uint256 private _value0xc;
//	    uint256 private _value0xd;
//	    uint256 private _value0xe;
//	    uint256 private _value0xf;
//	    uint256 private _value0x10;
//
//	    // Emitted when the stored value changes
//	    event ValueChanged(uint256 value);
//
//	    // Stores a new value in the contract
//	    function store(uint256 value) public {
//	        _value0x0 = value;
//	        _value0x1 = value;
//	        _value0x2 = value;
//	        _value0x3 = value;
//	        _value0x4 = value;
//	        _value0x5 = value;
//	        _value0x6 = value;
//	        _value0x7 = value;
//	        _value0x8 = value;
//	        _value0x9 = value;
//	        _value0xa = value;
//	        _value0xb = value;
//	        _value0xc = value;
//	        _value0xd = value;
//	        _value0xe = value;
//	        _value0xf = value;
//	        _value0x10 = value;
//	        emit ValueChanged(value);
//	    }
//
//	    // Reads the last stored value
//	    function retrieve() public view returns (uint256) {
//	        return _value0x0;
//	    }
//	}
//
// You may produce this hex string by saving the contract into a file
// Box.sol and invoking
//
//	solc Box.sol --bin --abi --optimize
//
// This contract is a slight modification of Box.sol to use more storage nodes
// and ensure the contract storage will contain at least 1 non-leaf node (by
// storing 17 values).
const contractHexString = "0x608060405234801561001057600080fd5b5061013f806100206000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80632e64cec11461003b5780636057361d14610050575b600080fd5b60005460405190815260200160405180910390f35b61006361005e3660046100f0565b610065565b005b6000819055600181905560028190556003819055600481905560058190556006819055600781905560088190556009819055600a819055600b819055600c819055600d819055600e819055600f81905560108190556040518181527f93fe6d397c74fdf1402a8b72e47b68512f0510d7b98a4bc4cbdf6ac7108b3c599060200160405180910390a150565b60006020828403121561010257600080fd5b503591905056fea2646970667358221220031e17f1bd1d1dcbee088287a905b152410b180064c149763590a0bbc516d95e64736f6c63430008130033"

var contractFuncSelector = crypto.Keccak256([]byte("store(uint256)"))[:4]

// contractInvocationData returns data suitable for invoking the 'store'
// function of the contract in contractHexString, note
func contractInvocationData(val byte) []byte {
	return hexutil.MustDecode(fmt.Sprintf("0x%x00000000000000000000000000000000000000000000000000000000000000%02x", contractFuncSelector, val))
}

func chainWithDeployedContract(t *testing.T) (*mock.MockSentry, common.Address, common.Address) {
	var (
		signer      = types.LatestSignerForChainID(nil)
		bankKey, _  = crypto.HexToECDSA("b71c71a67e1177ad4e901695e1b4b9ee17ae16c6668d313eac2f96dbcda3f291")
		bankAddress = crypto.PubkeyToAddress(bankKey.PublicKey)
		bankFunds   = big.NewInt(1e9)
		contract    = hexutil.MustDecode(contractHexString)
		gspec       = &types.Genesis{
			Config: chain.TestChainConfig,
			Alloc:  types.GenesisAlloc{bankAddress: {Balance: bankFunds}},
		}
	)
	m := mock.MockWithGenesis(t, gspec, bankKey, false)
	db := m.DB

	var contractAddr common.Address

	chain, err := core.GenerateChain(m.ChainConfig, m.Genesis, m.Engine, m.DB, 3, func(i int, block *core.BlockGen) {
		nonce := block.TxNonce(bankAddress)
		switch i {
		case 0:
			tx, err := types.SignTx(types.NewContractCreation(nonce, new(uint256.Int), 1e6, new(uint256.Int), contract), *signer, bankKey)
			require.NoError(t, err)
			block.AddTx(tx)
			contractAddr = crypto.CreateAddress(bankAddress, nonce)
		case 1:
			txn, err := types.SignTx(types.NewTransaction(nonce, contractAddr, new(uint256.Int), 900000, new(uint256.Int), contractInvocationData(1)), *signer, bankKey)
			require.NoError(t, err)
			block.AddTx(txn)
		case 2:
			txn, err := types.SignTx(types.NewTransaction(nonce, contractAddr, new(uint256.Int), 900000, new(uint256.Int), contractInvocationData(2)), *signer, bankKey)
			require.NoError(t, err)
			block.AddTx(txn)
		}
	})
	if err != nil {
		t.Fatalf("generate blocks: %v", err)
	}

	err = m.InsertChain(chain)
	require.NoError(t, err)

	tx, err := db.BeginTemporalRo(context.Background())
	if err != nil {
		t.Fatalf("read only db tx to read state: %v", err)
	}
	defer tx.Rollback()

	stateReader, err := rpchelper.CreateHistoryStateReader(tx, 1, 0, rawdbv3.TxNums)
	require.NoError(t, err)
	st := state.New(stateReader)
	require.NoError(t, err)
	exist, err := st.Exist(contractAddr)
	require.NoError(t, err)
	assert.False(t, exist, "Contract should not exist at block #1")

	stateReader, err = rpchelper.CreateHistoryStateReader(tx, 2, 0, rawdbv3.TxNums)
	require.NoError(t, err)
	st = state.New(stateReader)
	require.NoError(t, err)
	exist, err = st.Exist(contractAddr)
	require.NoError(t, err)
	assert.True(t, exist, "Contract should exist at block #2")

	return m, bankAddress, contractAddr
}

func doPrune(t *testing.T, db kv.RwDB, pruneTo uint64) {
	ctx := context.Background()
	//logger := testlog.Logger(t, log.LvlCrit)
	tx, err := db.BeginRw(ctx)
	require.NoError(t, err)

	logEvery := time.NewTicker(20 * time.Second)

	err = rawdb.PruneTableDupSort(tx, kv.TblAccountVals, "", pruneTo, logEvery, ctx)
	require.NoError(t, err)

	err = rawdb.PruneTableDupSort(tx, kv.StorageChangeSetDeprecated, "", pruneTo, logEvery, ctx)
	require.NoError(t, err)

	//err = rawdb.PruneTable(tx, kv.RCacheDomain, pruneTo, ctx, math.MaxInt32, time.Hour, logger, "")
	//require.NoError(t, err)

	err = tx.Commit()
	require.NoError(t, err)
}
