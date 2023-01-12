package delegateartifactcreation

import (
	"net/http"

	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/jwt"
	"github.com/MyriadFlow/storefront_gateway/config/smartcontract/rawtransaction"
	"github.com/MyriadFlow/storefront_gateway/generated/smartcontract/storefront"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/canaccess"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateArtifactCreation")
	{
		g.Use(jwt.JWT)
		g.POST("", deletegateArtifactCreation)
	}
}

func deletegateArtifactCreation(c *gin.Context) {
	var request DelegateArtifactCreationRequest
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
	abiS := storefront.StorefrontABI

	tx, err := rawtransaction.SendRawTransaction(abiS, "delegateArtifactCreation", creatorAddr, request.MetaDataHash)

	if err != nil {
		httphelper.NewInternalServerError(c, "", "failed to call %v of %v, error: %v", "delegateArtifactCreation", "StoreFront", err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	payload := DelegateArtifactCreationPayload{
		TransactionHash: transactionHash,
	}
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	httphelper.SuccessResponse(c, "request successfully send, artififact will be delegated soon", payload)
}
