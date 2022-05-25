package delegateartifactcreation

import (
	"net/http"

	"github.com/TheLazarusNetwork/marketplace-engine/api/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/marketplace-engine/config/smartcontract/rawtrasaction"
	gcreatify "github.com/TheLazarusNetwork/marketplace-engine/generated/smartcontract/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/canaccess"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
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
	abiS := gcreatify.CreatifyABI

	tx, err := rawtrasaction.SendRawTrasac(abiS, "delegateArtifactCreation", creatorAddr, request.MetaDataHash)

	if err != nil {
		httphelper.NewInternalServerError(c, "", "failed to call %v of %v, error: %v", "delegateArtifactCreation", "Creatify", err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	payload := DelegateArtifactCreationPayload{
		TransactionHash: transactionHash,
	}
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	httphelper.SuccessResponse(c, "request successfully send, artififact will be delegated soon", payload)
}
