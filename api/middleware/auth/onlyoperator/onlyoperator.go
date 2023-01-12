package onlyoperator

import (
	"net/http"

	"github.com/MyriadFlow/storefront_gateway/config/smartcontract"
	storefront "github.com/MyriadFlow/storefront_gateway/config/storefront"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// Used along with jwt middleware
func OnlyOperator(c *gin.Context) {
	walletAddress := c.GetString("walletAddress")
	client, err := smartcontract.GetClient()
	if err != nil {
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}
	instance, err := storefront.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "STOREFRONT", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}

	operatorRole, err := instance.STOREFRONTOPERATORROLE(nil)
	if err != nil {
		logwrapper.Errorf("Failed to get %v, error: %v", "STOREFRONTOPERATORROLE", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}
	hasRole, err := instance.HasRole(nil, operatorRole, common.HexToAddress(walletAddress))
	if err != nil {
		logwrapper.Errorf("failed to call %v smart contract function HasRole , error: %v", "STOREFRONT", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}

	if !hasRole {
		httphelper.ErrResponse(c, http.StatusForbidden, "only operator can access this API")
		c.Abort()
		return
	}
	c.Next()
}
