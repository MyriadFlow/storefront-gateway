package blockchains

import "github.com/MyriadFlow/storefront-gateway/config/envconfig"

type Blockchain struct {
	Name                string
	Type                string
	ChainId             int
	RpcHttp             string
	RpcWss              string
	GraphPort           string
	IpfsPort            string
	GraphHttpsUrl       string
	DeploymentName      string
	SubgraphNetworkName string
}

var Testnets map[string]Blockchain = map[string]Blockchain{
	"polygon": {
		Name:                "polygon",
		Type:                "testnet",
		ChainId:             80001,
		RpcHttp:             envconfig.EnvVars.POLYGON_TESTNET_HTTP,
		RpcWss:              envconfig.EnvVars.POLYGON_TESTNET_WSS,
		GraphPort:           envconfig.EnvVars.ETHEREUM_TESTNET_NODE_PORT,
		IpfsPort:            envconfig.EnvVars.POLYGON_TESTNET_IPFS_PORT,
		GraphHttpsUrl:       envconfig.EnvVars.POLYGON_TESTNET_GRAPH_HTTPS,
		DeploymentName:      "maticmum",
		SubgraphNetworkName: "mumbai"},
	"ethereum": {
		Name:                "ethereum",
		Type:                "testnet",
		ChainId:             11155111,
		RpcHttp:             envconfig.EnvVars.ETHEREUM_TESTNET_HTTP,
		RpcWss:              envconfig.EnvVars.ETHEREUM_TESTNET_WSS,
		GraphPort:           envconfig.EnvVars.ETHEREUM_TESTNET_NODE_PORT,
		IpfsPort:            envconfig.EnvVars.ETHEREUM_TESTNET_IPFS_PORT,
		GraphHttpsUrl:       envconfig.EnvVars.ETHEREUM_TESTNET_GRAPH_HTTPS,
		DeploymentName:      "sepolia",
		SubgraphNetworkName: "sepolia"},
	// "binance": {
	// 	Name:           "binance",
	// 	Type:           "testnet",
	// 	ChainId:        97,
	// 	RpcHttp:        envconfig.EnvVars.BINANCE_TESTNET_HTTP,
	// 	RpcWss:         envconfig.EnvVars.BINANCE_TESTNET_WSS,
	// 	GraphPort:      envconfig.EnvVars.BINANCE_TESTNET_NODE_PORT,
	// 	IpfsPort:       envconfig.EnvVars.FEVM_TESTNET_IPFS_PORT,
	// 	GraphHttpsUrl:  envconfig.EnvVars.BINANCE_TESTNET_GRAPH_HTTPS,
	// 	DeploymentName: "bnbTest"},
	// "fevm": {
	// 	Name:           "fevm",
	// 	Type:           "testnet",
	// 	ChainId:        314159,
	// 	RpcHttp:        envconfig.EnvVars.FEVM_TESTNET_HTTP,
	// 	RpcWss:         envconfig.EnvVars.FEVM_TESTNET_WSS,
	// 	GraphPort:      envconfig.EnvVars.FEVM_TESTNET_NODE_PORT,
	// 	IpfsPort:       envconfig.EnvVars.FEVM_TESTNET_IPFS_PORT,
	// 	GraphHttpsUrl:  envconfig.EnvVars.FEVM_TESTNET_GRAPH_HTTPS,
	// 	DeploymentName: "filecoinCalibaration"},
}
