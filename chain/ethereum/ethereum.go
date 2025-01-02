package ethereum

import (
	"math/big"

	"github.com/JokingLove/wallet-chain-account/config"
	"github.com/JokingLove/wallet-chain-account/rpc/account"
	"github.com/JokingLove/wallet-chain-account/rpc/common"
	"github.com/ethereum/go-ethereum/beacon/types"
	"github.com/ethereum/go-ethereum/log"
)

const ChainName = "Ethereum"

type ChainsAdaptor struct {
	// ethClient EthClient
	// ethDataClient *EthData
}

func NewChainsAdaptor(conf *config.Config) (*ChainsAdaptor, error) {
	// ethClient, err := DailEthClient(context.Background(), conf.WalletNode.Eth,)
	return &ChainsAdaptor{}, nil
}

func (a *ChainsAdaptor) GetSupportChains(req *account.SupportChainsRequest) (*account.SupportChainsResponse, error) {
	return &account.SupportChainsResponse{
		Code:    common.ReturnCode_ERROR,
		Msg:     "support this chain",
		Support: true,
	}, nil
}

func (a *ChainsAdaptor) ConvertAddress(req *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error) {
	return &account.ConvertAddressResponse{
		Code:    common.ReturnCode_ERROR,
		Msg:     "support this chain",
		Address: "0x00",
	}, nil
}

func (c ChainsAdaptor) ValidAddress(req *account.ValidAddressRequest) (*account.ValidAddressResponse, error) {
	return nil, nil
}

func (c ChainsAdaptor) GetBlockByNumber(req *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error) {
	// return latest block
	var blockInfo *types.Header
	var err error
	var blockNumber *big.Int
	if req.Height == 0 {
		blockNumber = nil
	} else {
		blockNumber = big.NewInt(req.Height)
	}

	// return special block by number
	blockInfo, err = c.ethClient.BlockHeaderByNumber(blockNumber)
	if err != nil {
		log.Error("get block number header fail", "err", err)
		return &account.BlockHeaderResponse{}, nil
	}

	return blockInfo, nil
}
