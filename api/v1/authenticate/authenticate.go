package authenticate

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/models/claims"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"

	"github.com/MyriadFlow/storefront-gateway/util/pkg/auth"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/cryptosign"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/authenticate")
	{
		g.POST("", authenticate)
	}
}

func authenticate(c *gin.Context) {

	db := dbconfig.GetDb()
	var req AuthenticateRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	//Get flowid type
	var flowIdData models.FlowId
	err = db.Model(&models.FlowId{}).Where("flow_id = ?", req.FlowId).First(&flowIdData).Error
	if err != nil {
		logwrapper.Errorf("failed to get flowId, error %v", err)
		httphelper.ErrResponse(c, http.StatusNotFound, "flow id not found")
		return
	}

	if flowIdData.FlowIdType != models.AUTH {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Flow id not created for auth")
		return
	}

	if err != nil {
		logwrapper.Error(err)
		httphelper.ErrResponse(c, 500, "Unexpected error occured")
		return
	}
	userAuthEULA := envconfig.EnvVars.AUTH_EULA
	message := userAuthEULA + req.FlowId
	walletAddress, isCorrect, err := cryptosign.CheckSign(req.Signature, req.FlowId, message)

	if err == cryptosign.ErrFlowIdNotFound {
		httphelper.ErrResponse(c, http.StatusNotFound, "Flow Id not found")
		return
	}

	if err != nil {
		logwrapper.Errorf("failed to CheckSignature, error %v", err.Error())
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	if isCorrect {
		customClaims := claims.New(walletAddress)
		pasetoToken, err := auth.GenerateTokenPaseto(customClaims)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to generate token, error %v", err.Error())
			return
		}
		err = db.Where("flow_id = ?", req.FlowId).Delete(&models.FlowId{}).Error
		if err != nil {
			httphelper.NewInternalServerError(c, "", "failed to delete flowId, error %v", err.Error())
			return
		}
		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		httphelper.SuccessResponse(c, "Token generated successfully", payload)
	} else {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}
