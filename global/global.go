package global

import (
	"github.com/TheLazarusNetwork/marketplace-engine/config/envconfig"
)

var AllowedWalletAddresses []string

func InitGlobal() {
	AllowedWalletAddresses = envconfig.EnvVars.ALLOWED_WALLET_ADDRESS
}
