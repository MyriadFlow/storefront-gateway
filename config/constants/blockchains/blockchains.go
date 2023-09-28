package blockchains

import "github.com/MyriadFlow/storefront-gateway/config/envconfig"

type Blockchain struct {
	Name    string `json:"name"`
	Type    string `json:"type"`
	ChainId int    `json:"chainId"`
	RpcHttp string `json:"rpcHttp"`
	RpcWss  string `json:"rpcWss"`
}

var Testnets map[string]Blockchain = map[string]Blockchain{
	"polygon":  {Name: "polygon", Type: "testnet", ChainId: 80001, RpcHttp: envconfig.EnvVars.POLYGON_TESTNET_HTTP, RpcWss: envconfig.EnvVars.POLYGON_TESTNET_WSS},
	"ethereum": {Name: "ethereum", Type: "testnet", ChainId: 11155111, RpcHttp: envconfig.EnvVars.ETHEREUM_TESTNET_HTTP, RpcWss: envconfig.EnvVars.ETHEREUM_TESTNET_WSS},
	"binance":  {Name: "binance", Type: "testnet", ChainId: 97, RpcHttp: envconfig.EnvVars.BINANCE_TESTNET_HTTP, RpcWss: envconfig.EnvVars.BINANCE_TESTNET_WSS},
	"fevm":     {Name: "fevm", Type: "testnet", ChainId: 314159, RpcHttp: envconfig.EnvVars.FEVM_TESTNET_HTTP, RpcWss: envconfig.EnvVars.FEVM_TESTNET_WSS},
}
