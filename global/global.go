package global

import (
	"strings"

	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
)

var AllowedWalletAddresses []string

func InitGlobal() {
	envAllowedWalletAddress := envutil.MustGetEnv("ALLOWED_WALLET_ADDRESS")
	AllowedWalletAddresses = strings.Split(envAllowedWalletAddress, ",")
}
