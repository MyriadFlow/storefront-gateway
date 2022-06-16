package onlyoperator

import (
	"net/http"

	"github.com/TheLazarusNetwork/marketplace-engine/config/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// Used along with jwt middleware
func OnlyOperator(c *gin.Context) {
	walletAddress := c.GetString("walletAddress")
	client, err := smartcontract.GetClient()
	if err != nil {
		httphelper.InternalServerError(c)
		return
	}
	instance, err := creatify.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "CREATIFY", err.Error())
		httphelper.InternalServerError(c)
		return
	}

	operatorRole, err := instance.CREATIFYOPERATORROLE(nil)
	if err != nil {
		logwrapper.Errorf("Failed to get %v, error: %v", "CREATIFYOPERATORROLE", err.Error())
		httphelper.InternalServerError(c)
		return
	}
	hasRole, err := instance.HasRole(nil, operatorRole, common.HexToAddress(walletAddress))
	if err != nil {
		logwrapper.Errorf("failed to call %v smart contract function HasRole , error: %v", "CREATIFY", err.Error())
		httphelper.InternalServerError(c)
		return
	}

	if !hasRole {
		httphelper.ErrResponse(c, http.StatusForbidden, "only operator can access this API")
		return
	}
	c.Next()
}
