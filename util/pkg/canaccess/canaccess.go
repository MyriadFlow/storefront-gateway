package canaccess

import (
	"strings"

	"github.com/TheLazarusNetwork/marketplace-engine/global"
)

func CanAccess(walletAddress string) bool {
	if global.AllowedWalletAddresses[0] == "*" {
		return true
	}
	for _, v := range global.AllowedWalletAddresses {
		if strings.EqualFold(v, walletAddress) {
			return true
		}
	}
	return false
}
