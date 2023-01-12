package global

import (
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
)

var AllowedWalletAddresses []string

func InitGlobal() {
	AllowedWalletAddresses = envconfig.EnvVars.ALLOWED_WALLET_ADDRESS
}
