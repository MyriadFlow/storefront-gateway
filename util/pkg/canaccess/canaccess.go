package canaccess

import (
	"strings"

	"github.com/MyriadFlow/storefront-gateway/global"
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
