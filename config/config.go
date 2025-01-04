package config

import (
	"os"

	"gopkg.in/yaml.v2"

	"github.com/ethereum/go-ethereum/log"
)

type Server struct {
	Port string `yaml:"port"`
}

// type RPC struct {
// 	RPCURL  string `yaml:"rpc_url"`
// 	RPCUser string `yaml:"rpc_user"`
// 	RPCPass string `yaml:"rpc_pass"`
// }

type Node struct {
	RpcUrl       string `json:"rpc_url"`
	RpcUser      string `json:"rpc_user"`
	RpcPass      string `json:"rpc_pass"`
	DataApiUrl   string `json:"data_api_url" yaml:"data_api_url"`
	DataApiKey   string `json:"data_api_key" yaml:"data_api_key"`
	DataApiToken string `json:"data_api_token" yaml:"data_api_token"`
	TimeOut      uint64 `json:"timeout" yaml:"timeout"`
}

type WalletNode struct {
	Eth     Node `yaml:"eth"`
	Arbi    Node `yaml:"arbi"`
	Op      Node `yaml:"op"`
	Zksync  Node `yaml:"zksync"`
	Bsc     Node `yaml:"bsc"`
	Heco    Node `yaml:"heco"`
	Avax    Node `yaml:"avax"`
	Polygon Node `yaml:"polygon"`
	Tron    Node `yaml:"tron"`
	Sol     Node `yaml:"sol"`
	Cosmos  Node `yaml:"cosmos"`
	Aptos   Node `yaml:"aptos"`
	Mantle  Node `yaml:"mantle"`
	Scroll  Node `yaml:"scroll"`
	Base    Node `yaml:"base"`
	Linea   Node `yaml:"linea"`
	Sui     Node `yaml:"sui"`
	Ton     Node `yaml:"ton"`
}

type Config struct {
	Server     Server     `yaml:"server"`
	WalletNode WalletNode `yaml:"wallet_node"`
	NetWork    string     `yaml:"network"`
	Chains     []string   `yaml:"chains"`
}

func New(path string) (*Config, error) {
	// config global config instance
	var config = new(Config)
	h := log.NewTerminalHandler(os.Stdout, true)
	log.SetDefault(log.NewLogger(h))

	data, err := os.ReadFile(path)
	if err != nil {
		log.Error("read config file error", "err", err)
		return nil, err
	}

	err = yaml.Unmarshal(data, config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

const UnsupportedChain = "Unsupport chain"
const UnsupportedOperation = UnsupportedChain
