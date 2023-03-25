package envconfig

import (
	"log"

	"github.com/caarlos0/env/v6"
)

type config struct {
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
	POLYGON_RPC                  string `env:"POLYGON_RPC,required"`
	MNEMONIC                     string `env:"MNEMONIC,required"`
	NFT_STORAGE_API_KEY          string `env:"NFT_STORAGE_API_KEY,required" `

	PASETO_SIGNED_BY           string `env:"PASETO_SIGNED_BY,required"`
	PASETO_EXPIRATION_IN_HOURS string `env:"PASETO_EXPIRATION_IN_HOURS,required" `

	ORG_NAME         string `env:"ORG_NAME,required"`
	HOME_TITLE       string `env:"HOME_TITLE,required"`
	HOME_DESCRIPTION string `env:"HOME_DESCRIPTION,required"`
	FOOTER           string `env:"FOOTER,required"`

	ALLOWED_WALLET_ADDRESS []string `env:"ALLOWED_WALLET_ADDRESS,required" envSeparator:","`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
