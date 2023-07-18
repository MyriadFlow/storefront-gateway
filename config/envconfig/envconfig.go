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
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
