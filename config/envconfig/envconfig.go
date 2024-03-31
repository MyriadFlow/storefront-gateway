package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
	APP_ID             string   `env:"APP_ID,required"`
	APP_NAME           string   `env:"APP_NAME,required"`
	APP_PORT           int      `env:"APP_PORT,required"`
	APP_MODE           string   `env:"APP_MODE,required"`
	APP_ALLOWED_ORIGIN []string `env:"APP_ALLOWED_ORIGIN,required" envSeparator:","`

	DB_HOST     string `env:"DB_HOST,required"`
	DB_USERNAME string `env:"DB_USERNAME,required"`
	DB_PASSWORD string `env:"DB_PASSWORD,required"`
	DB_NAME     string `env:"DB_NAME,required"`
	DB_PORT     int    `env:"DB_PORT,required"`

	AUTH_EULA                    string `env:"AUTH_EULA,required"`
	MARKETPLACE_CONTRACT_ADDRESS string `env:"MARKETPLACE_CONTRACT_ADDRESS,required"`
	STOREFRONT_CONTRACT_ADDRESS  string `env:"STOREFRONT_CONTRACT_ADDRESS"`
	POLYGON_RPC_HTTP             string `env:"POLYGON_RPC_HTTP,required"`
	MNEMONIC                     string `env:"MNEMONIC,required"`
	NFT_STORAGE_API_KEY          string `env:"NFT_STORAGE_API_KEY,required" `

	PASETO_SIGNED_BY           string `env:"PASETO_SIGNED_BY,required"`
	PASETO_FOOTER              string `env:"PASETO_FOOTER,required"`
	PASETO_EXPIRATION_IN_HOURS string `env:"PASETO_EXPIRATION_IN_HOURS,required"`

	ALLOWED_WALLET_ADDRESS        []string `env:"ALLOWED_WALLET_ADDRESS,required" envSeparator:","`
	SMARTCONTRACT_API_URL         string   `env:"SMARTCONTRACT_API_URL,required"`
	SUBSCRIPTION_CONTRACT_ADDRESS string   `env:"SUBSCRIPTION_CONTRACT_ADDRESS,required"`
	POLYGON_RPC_WSS               string   `env:"POLYGON_RPC_WSS,required"`
	DEFAULT_SUBSCRIPTION_IMAGE    string   `env:"DEFAULT_SUBSCRIPTION_IMAGE,required"`
	SUBGRAPH_SERVER_URL           string   `env:"SUBGRAPH_SERVER_URL,required"`
	NODECTL_SERVER_URL            string   `env:"NODECTL_SERVER_URL,required"`
	NODECTL_SERVER_PORT           string   `env:"NODECTL_SERVER_PORT,required"`
	BASE_URL_GATEWAY              string   `env:"BASE_URL_GATEWAY,required"`
	IPFS_GATEWAY                  string   `env:"IPFS_GATEWAY,required"`
	RAINBOWKIT_PROJECT_ID         string   `env:"RAINBOWKIT_PROJECT_ID,required"`
	ALCHEMY_ID                    string   `env:"ALCHEMY_ID,required"`
	POLYGON_TESTNET_HTTP          string   `env:"POLYGON_TESTNET_HTTP,required"`
	POLYGON_TESTNET_WSS           string   `env:"POLYGON_TESTNET_WSS,required"`
	POLYGON_TESTNET_NODE_PORT     string   `env:"POLYGON_TESTNET_NODE_PORT,required"`
	POLYGON_TESTNET_IPFS_PORT     string   `env:"POLYGON_TESTNET_IPFS_PORT,required"`
	POLYGON_TESTNET_GRAPH_HTTPS   string   `env:"POLYGON_TESTNET_GRAPH_HTTPS,required"`
	ETHEREUM_TESTNET_HTTP         string   `env:"ETHEREUM_TESTNET_HTTP,required"`
	ETHEREUM_TESTNET_WSS          string   `env:"ETHEREUM_TESTNET_WSS,required"`
	ETHEREUM_TESTNET_NODE_PORT    string   `env:"ETHEREUM_TESTNET_NODE_PORT,required"`
	ETHEREUM_TESTNET_IPFS_PORT    string   `env:"ETHEREUM_TESTNET_IPFS_PORT,required"`
	ETHEREUM_TESTNET_GRAPH_HTTPS  string   `env:"ETHEREUM_TESTNET_GRAPH_HTTPS,required"`
	ARBITRUM_TESTNET_HTTP         string   `env:"ARBITRUM_TESTNET_HTTP,required"`
	ARBITRUM_TESTNET_WSS          string   `env:"ARBITRUM_TESTNET_WSS,required"`
	ARBITRUM_TESTNET_NODE_PORT    string   `env:"ARBITRUM_TESTNET_NODE_PORT,required"`
	ARBITRUM_TESTNET_IPFS_PORT    string   `env:"ARBITRUM_TESTNET_IPFS_PORT,required"`
	ARBITRUM_TESTNET_GRAPH_HTTPS  string   `env:"ARBITRUM_TESTNET_GRAPH_HTTPS,required"`
	OPTIMISM_TESTNET_HTTP         string   `env:"OPTIMISM_TESTNET_HTTP,required"`
	OPTIMISM_TESTNET_WSS          string   `env:"OPTIMISM_TESTNET_WSS,required"`
	OPTIMISM_TESTNET_NODE_PORT    string   `env:"OPTIMISM_TESTNET_NODE_PORT,required"`
	OPTIMISM_TESTNET_IPFS_PORT    string   `env:"OPTIMISM_TESTNET_IPFS_PORT,required"`
	OPTIMISM_TESTNET_GRAPH_HTTPS  string   `env:"OPTIMISM_TESTNET_GRAPH_HTTPS,required"`
	BASE_TESTNET_HTTP             string   `env:"BASE_TESTNET_HTTP,required"`
	BASE_TESTNET_WSS              string   `env:"BASE_TESTNET_WSS,required"`
	BASE_TESTNET_NODE_PORT        string   `env:"BASE_TESTNET_NODE_PORT,required"`
	BASE_TESTNET_IPFS_PORT        string   `env:"BASE_TESTNET_IPFS_PORT,required"`
	BASE_TESTNET_GRAPH_HTTPS      string   `env:"BASE_TESTNET_GRAPH_HTTPS,required"`
	WALLET_PRIVATE_KEY            string   `env:"WALLET_PRIVATE_KEY,required"`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
