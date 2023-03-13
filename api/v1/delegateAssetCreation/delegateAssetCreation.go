package delegateartifactcreation

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract/rawtransaction"
	"github.com/MyriadFlow/storefront-gateway/generated/smartcontract/storefront"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/canaccess"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateAssetCreation")
	{
		g.Use(paseto.PASETO)
		g.POST("", deletegateAssetCreation)
	}
}

func deletegateAssetCreation(c *gin.Context) {
	var request DelegateAssetCreationRequest
	walletAddressGin := c.GetString("walletAddress")
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	if !canaccess.CanAccess(walletAddressGin) {
		httphelper.ErrResponse(c, http.StatusForbidden, "this api not open to current wallet")
		return
	}
	creatorAddr := common.HexToAddress(request.CreatorAddress)
	abiS := storefront.StoreABI

	tx, err := rawtransaction.SendRawTransaction(abiS, "delegateAssetCreation", creatorAddr, request.MetaDataHash)

	if err != nil {
		httphelper.NewInternalServerError(c, "", "failed to call %v of %v, error: %v", "delegateAssetCreation", "StoreFront", err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	payload := DelegateAssetCreationPayload{
		TransactionHash: transactionHash,
	}
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	httphelper.SuccessResponse(c, "request successfully send, asset will be delegated soon", payload)
}
