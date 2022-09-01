package envconfig

import (
	"log"
	"time"

	"github.com/caarlos0/env/v6"
)

type config struct {
	JWT_PRIVATE_KEY           string        `env:"JWT_PRIVATE_KEY,required"`
	JWT_EXPIRATION            time.Duration `env:"JWT_EXPIRATION,required"`
	APP_PORT                  int           `env:"APP_PORT,required"`
	AUTH_EULA                 string        `env:"AUTH_EULA,required"`
	CREATOR_EULA              string        `env:"CREATOR_EULA,required"`
	APP_NAME                  string        `env:"APP_NAME,required"`
	GIN_MODE                  string        `env:"GIN_MODE,required"`
	DB_HOST                   string        `env:"DB_HOST,required"`
	DB_USERNAME               string        `env:"DB_USERNAME,required"`
	DB_PASSWORD               string        `env:"DB_PASSWORD,required"`
	DB_NAME                   string        `env:"DB_NAME,required"`
	DB_PORT                   int           `env:"DB_PORT,required"`
	CREATIFY_CONTRACT_ADDRESS string        `env:"CREATIFY_CONTRACT_ADDRESS,required"`
	POLYGON_RPC               string        `env:"POLYGON_RPC,required"`
	MNEMONIC                  string        `env:"MNEMONIC,required"`
	IPFS_NODE_URL             string        `env:"IPFS_NODE_URL,required"`
	ALLOWED_ORIGIN            []string      `env:"ALLOWED_ORIGIN,required" envSeparator:","`
	SIGNED_BY                 string        `env:"SIGNED_BY,required"`
	ALLOWED_WALLET_ADDRESS    []string      `env:"ALLOWED_WALLET_ADDRESS,required" envSeparator:","`

	ORG_NAME                     string   `env:"ORG_NAME,required"`
	HOME_TITLE                   string   `env:"HOME_TITLE,required"`
	HOME_DESCRIPTION             string   `env:"HOME_DESCRIPTION,required"`
	GRAPH_URL                    string   `env:"GRAPH_URL,required"`
	CREATIFY_ADDRESS             string   `env:"CREATIFY_ADDRESS,required"`
	MARKETPLACE_CONTRACT_ADDRESS string   `env:"MARKETPLACE_CONTRACT_ADDRESS,required"`
	TOP_HIGHLIGHTS               []string `env:"TOP_HIGHLIGHTS,required" envSeparator:","`
	TRENDINGS                    []string `env:"TRENDINGS,required" envSeparator:","`
	TOP_BIDS                     []string `env:"TOP_BIDS,required" envSeparator:","`
	FOOTER                       string   `env:"FOOTER,required"`
}

var EnvVars config = config{}

func InitEnvVars() {

	if err := env.Parse(&EnvVars); err != nil {
		log.Fatalf("failed to parse EnvVars: %s", err)
	}
}
