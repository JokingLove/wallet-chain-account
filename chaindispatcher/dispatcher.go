package chaindispatcher

import (
	"context"
	"runtime/debug"
	"strings"

	"github.com/JokingLove/wallet-chain-account/chain"
	"github.com/JokingLove/wallet-chain-account/chain/ethereum"
	"github.com/JokingLove/wallet-chain-account/config"
	"github.com/JokingLove/wallet-chain-account/rpc/account"
	"github.com/JokingLove/wallet-chain-account/rpc/common"
	"github.com/ethereum/go-ethereum/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CommonRequest interface {
	GetChain() string
}

type CommonReply = account.SupportChainsResponse

type ChainType = string

type ChainDispatcher struct {
	registry map[ChainType]chain.IChainsAdaptor
}

func New(conf *config.Config) (*ChainDispatcher, error) {
	dispatcher := ChainDispatcher{
		registry: make(map[ChainType]chain.IChainsAdaptor),
	}
	chainAdaptorFactoryMap := map[string]func(conf *config.Config) (chain.IChainsAdaptor, error){
		// add support chain
		ethereum.ChainName: ethereum.NewChainAdaptor,
	}

	supportedChains := []string{
		// addd support chain name
		ethereum.ChainName,
	}

	for _, c := range conf.Chains {
		if factory, ok := chainAdaptorFactoryMap[c]; ok {
			adaptor, err := factory(conf)
			if err != nil {
				log.Crit("failed to setup chain", "chain", c, "error", err)
			}
			dispatcher.registry[c] = adaptor
		} else {
			log.Error("unsupported chain", "chain", c, " supportedChains", supportedChains)
		}
	}
	return &dispatcher, nil
}

func (d *ChainDispatcher) Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	defer func() {
		if e := recover(); e != nil {
			log.Error("panic error ", "msg", e)
			log.Debug(string(debug.Stack()))
			err = status.Errorf(codes.Internal, "Panic err: %v", e)
		}
	}()

	pos := strings.LastIndex(info.FullMethod, "/")
	method := info.FullMethod[pos+1]
	chainName := req.(CommonRequest).GetChain()
	log.Info(string(method), "chain", chainName, "req", req)

	resp, err = handler(ctx, req)
	log.Debug("Finish handling", "resp", resp, "err", err)

	return
}

func (d *ChainDispatcher) preHandler(req interface{}) (resp *CommonReply) {
	chainName := req.(CommonRequest).GetChain()
	if _, ok := d.registry[chainName]; !ok {
		return &CommonReply{
			Code:    common.ReturnCode_ERROR,
			Msg:     config.UnsupportedOperation,
			Support: false,
		}
	}
	return nil
}

func (d *ChainDispatcher) GetSupportChains(ctx context.Context, request *account.SupportChainsRequest) (*account.SupportChainsResponse, error) {
	resp := d.preHandler(request)
	if resp != nil {
		return &account.SupportChainsResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  config.UnsupportedOperation,
		}, nil
	}
	return d.registry[request.Chain].GetSupportChains(request)
}

func (d *ChainDispatcher) ConvertAddress(ctx context.Context, req *account.ConvertAddressRequest) (*account.ConvertAddressResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.ConvertAddressResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "convert address fail",
		}, nil
	}
	return d.registry[req.Chain].ConvertAddress(req)
}

func (d *ChainDispatcher) ValidAddress(ctx context.Context, req *account.ValidAddressRequest) (*account.ValidAddressResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.ValidAddressResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "valid address fail",
		}, nil
	}
	return d.registry[req.Chain].ValidAddress(req)
}

func (d *ChainDispatcher) GetBlockByNumber(ctx context.Context, req *account.BlockNumberRequest) (*account.BlockResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.BlockResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block by number fail",
		}, nil
	}
	return d.registry[req.Chain].GetBlockByNumber(req)
}

func (d *ChainDispatcher) GetBlockByHash(ctx context.Context, req *account.BlockHashRequest) (*account.BlockResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.BlockResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block by hash fail",
		}, nil
	}
	return d.registry[req.Chain].GetBlockByHash(req)
}

func (d *ChainDispatcher) GetBlockHeaderByHash(ctx context.Context, req *account.BlockHeaderHashRequest) (*account.BlockHeaderResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.BlockHeaderResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block header by hash fail",
		}, nil
	}
	return d.registry[req.Chain].GetBlockHeaderByHash(req)
}

func (d *ChainDispatcher) GetBlockHeaderByNumber(ctx context.Context, req *account.BlockHeaderNumberRequest) (*account.BlockHeaderResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.BlockHeaderResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block header by number fail",
		}, nil
	}
	return d.registry[req.Chain].GetBlockHeaderByNumber(req)
}

func (d *ChainDispatcher) GetBlockHeaderByRange(ctx context.Context, req *account.BlockByRangeRequest) (*account.BlockByRangeResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.BlockByRangeResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get block header by range fail",
		}, nil
	}
	return nil, nil
}

func (d *ChainDispatcher) GetAccount(ctx context.Context, req *account.AccountRequest) (*account.AccountResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.AccountResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get account fail",
		}, nil
	}
	return d.registry[req.Chain].GetAccount(req)
}

func (d *ChainDispatcher) GetFee(ctx context.Context, req *account.FeeRequest) (*account.FeeResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.FeeResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get fee fail",
		}, nil
	}
	return d.registry[req.Chain].GetFee(req)
}

func (d *ChainDispatcher) SendTx(ctx context.Context, req *account.SendTxRequest) (*account.SendTxResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.SendTxResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "send tx fail",
		}, nil
	}
	return d.registry[req.Chain].SendTx(req)
}

func (d *ChainDispatcher) GetTxByAddress(ctx context.Context, req *account.TxAddressRequest) (*account.TxAddressResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.TxAddressResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get tx by address fail",
		}, nil
	}
	return d.registry[req.Chain].GetTxByAddress(req)
}

func (d *ChainDispatcher) GetTxByHash(ctx context.Context, req *account.TxHashRequest) (*account.TxHashResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.TxHashResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get tx by hash fail",
		}, nil
	}
	return d.registry[req.Chain].GetTxByHash(req)
}

func (d *ChainDispatcher) CreateUnSignTransaction(ctx context.Context, req *account.UnSignTransactionRequest) (*account.UnSignTransactionResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.UnSignTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "create unsigned transaction fail",
		}, nil
	}
	return d.registry[req.Chain].CreateUnSignTransaction(req)
}

func (d *ChainDispatcher) BuildSignedTransaction(ctx context.Context, req *account.SignedTransactionRequest) (*account.SignedTransactionResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.SignedTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "build signed transaction fail",
		}, nil
	}
	return d.registry[req.Chain].BuildSignedTransaction(req)
}

func (d *ChainDispatcher) DecodeTransaction(ctx context.Context, req *account.DecodeTransactionRequest) (*account.DecodeTransactionResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.DecodeTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "decode transaction fail",
		}, nil
	}
	return d.registry[req.Chain].DecodeTransaction(req)
}

func (d *ChainDispatcher) VerifySignedTransaction(ctx context.Context, req *account.VerifyTransactionRequest) (*account.VerifyTransactionResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.VerifyTransactionResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "verify signed transaction fail",
		}, nil
	}
	return d.registry[req.Chain].VerifySignedTransaction(req)
}

func (d *ChainDispatcher) GetExtraData(ctx context.Context, req *account.ExtraDataRequest) (*account.ExtraDataResponse, error) {
	resp := d.preHandler(req)
	if resp != nil {
		return &account.ExtraDataResponse{
			Code: common.ReturnCode_ERROR,
			Msg:  "get extra data fail",
		}, nil
	}
	return d.registry[req.Chain].GetExtraData(req)
}
