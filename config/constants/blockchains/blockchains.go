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

var Mainnets map[string]Blockchain

var Testnets map[string]Blockchain

func InitChains() {
	Testnets = map[string]Blockchain{
		"polygon": {
			Name:                "polygon",
			Type:                "testnet",
			ChainId:             80001,
			RpcHttp:             envconfig.EnvVars.POLYGON_TESTNET_HTTP,
			RpcWss:              envconfig.EnvVars.POLYGON_TESTNET_WSS,
			GraphPort:           envconfig.EnvVars.POLYGON_TESTNET_NODE_PORT,
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
			SubgraphNetworkName: "sepolia",
		},
		"arbitrum": {
			Name:                "arbitrum",
			Type:                "testnet",
			ChainId:             421613,
			RpcHttp:             envconfig.EnvVars.ARBITRUM_TESTNET_HTTP,
			RpcWss:              envconfig.EnvVars.ARBITRUM_TESTNET_WSS,
			GraphPort:           envconfig.EnvVars.ARBITRUM_TESTNET_NODE_PORT,
			IpfsPort:            envconfig.EnvVars.ARBITRUM_TESTNET_IPFS_PORT,
			GraphHttpsUrl:       envconfig.EnvVars.ARBITRUM_TESTNET_GRAPH_HTTPS,
			DeploymentName:      "arbGoerli",
			SubgraphNetworkName: "arbitrum-goerli",
		},
		"optimism": {
			Name:                "optimism",
			Type:                "testnet",
			ChainId:             420,
			RpcHttp:             envconfig.EnvVars.OPTIMISM_TESTNET_HTTP,
			RpcWss:              envconfig.EnvVars.OPTIMISM_TESTNET_WSS,
			GraphPort:           envconfig.EnvVars.OPTIMISM_TESTNET_NODE_PORT,
			IpfsPort:            envconfig.EnvVars.OPTIMISM_TESTNET_IPFS_PORT,
			GraphHttpsUrl:       envconfig.EnvVars.OPTIMISM_TESTNET_GRAPH_HTTPS,
			DeploymentName:      "optGoerli",
			SubgraphNetworkName: "arbitrum-goerli",
		},
		"base": {
			Name:                "base",
			Type:                "testnet",
			ChainId:             84531,
			RpcHttp:             envconfig.EnvVars.BASE_TESTNET_HTTP,
			RpcWss:              envconfig.EnvVars.BASE_TESTNET_WSS,
			GraphPort:           envconfig.EnvVars.BASE_TESTNET_NODE_PORT,
			IpfsPort:            envconfig.EnvVars.BASE_TESTNET_IPFS_PORT,
			GraphHttpsUrl:       envconfig.EnvVars.BASE_TESTNET_GRAPH_HTTPS,
			DeploymentName:      "baseGoerli",
			SubgraphNetworkName: "base-testnet",
		},
	}
}
