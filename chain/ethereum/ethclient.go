package ethereum

import (
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

const (
	defaultDialTimeout    = 5 * time.Second
	defaultDialAttempts   = 5
	defaultRequestTimeout = 5 * time.Second
)

type TransactionList struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Hash  string `json:"hash"`
	Value string `json:"value"`
}

type RpcBlock struct {
	Hash         common.Hash       `json:"hash"`
	Height       uint64            `json:"height"`
	Transactions []TransactionList `json:"transactions"`
	BaseFee      string            `json:"baseFeePerGas"`
}

type clnt struct {
	rpc RPC
}

type Logs struct {
	Logs          []types.Log
	ToBlockHeader *types.Header
}

type EthClient interface {
	BlockHeaderByNumber(*big.Int) (*types.Header, error)
	BlockHeaderByHash(common.Hash) (*types.Header, error)
	BlockHeaderByRange(*big.Int, *big.Int, uint) ([]types.Header, error)
	BlockByNumber(*big.Int) (*RpcBlock, error)
	BlockByHash(common.Hash) (*RpcBlock, error)

	LatestSafeBlockHeader() (*types.Header, error)
	LatestFinalizedBlockHeader() (*types.Header, error)

	TxCountByAddress(common.Address) (hexutil.Uint64, error)
	SuggestGasPrice() (*big.Int, error)
	SuggestGasTipCap() (*big.Int, error)

	SendRawTransaction(rawTx string) (*common.Hash, error)

	TxByHash(common.Hash) (*types.Transaction, error)
	TxReceiptByHash(common.Hash) (*types.Receipt, error)

	StorageHash(common.Address, *big.Int) (common.Hash, error)
	EthGetCode(common.Address) (string, error)
	GetBalance(address common.Address) (*big.Int, error)
	FilterLogs(filterQuery ethereum.FilterQuery, chainId uint) (Logs, error)
	Close()
}
