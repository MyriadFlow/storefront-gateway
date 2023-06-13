package authenticate

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/models/claims"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/auth"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/cryptosign"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/flowid"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common/hexutil"
	log "github.com/sirupsen/logrus"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/auth")
	{
		g.GET("/web3", getFlowId)
		g.POST("/web3", authenticate)
		g.POST("/web2", Web2Auth)
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

		payload := AuthenticatePayload{
			Token: pasetoToken,
		}
		httphelper.SuccessResponse(c, "Token generated successfully", payload)
	} else {
		httphelper.ErrResponse(c, http.StatusForbidden, "Wallet Address is not correct")
		return
	}
}

func getFlowId(c *gin.Context) {
	walletAddress := c.Query("walletAddress")

	if walletAddress == "" {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Wallet address (walletAddress) is required")
		return
	}
	_, err := hexutil.Decode(walletAddress)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusBadRequest, "Wallet address (walletAddress) is not valid")
		return
	}
	flowId, err := flowid.GenerateFlowId(walletAddress, models.AUTH, "")
	if err != nil {
		log.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	userAuthEULA := envconfig.EnvVars.AUTH_EULA
	payload := GetFlowIdPayload{
		FlowId: flowId,
		Eula:   userAuthEULA,
	}
	httphelper.SuccessResponse(c, "Flowid successfully generated", payload)
}

func Web2Auth(c *gin.Context) {

	db := dbconfig.GetDb()
	var req = web2AuthRequest{}

	if err := c.ShouldBindJSON(&req); err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	} else {

		var userID, email string
		var supabaseErr error

		req.User_Type = "web2"

		if req.Provider == "supabase" {

			userID, email, supabaseErr = auth.FromJWTSupabaseTokenGetData(c, req.Token)

			if supabaseErr != nil || userID == "" || email == "" {
				logwrapper.Infof("userId : %v email : %v", userID, email)
				httphelper.ErrResponse(c, http.StatusForbidden, supabaseErr.Error())
				return
			}

			newUser := models.User{
				Email:    email,
				UserType: req.User_Type,
			}

			var count int64
			result := db.Model(&models.User{}).Where("email = ?", email).Count(&count)
			if result.Error != nil {
				logwrapper.Errorf("failed to get user details >> Web2Auth, error %v", err)
				httphelper.ErrResponse(c, http.StatusForbidden, "Web2Auth failed to count user")
				return
			}

			if count == 1 && email != "" {
				result := db.Where("email = ?", newUser.Email).First(&models.User{})
				if result.Error != nil {
					logwrapper.Errorf("failed to create new user, error %v for email %v", err, email)
					httphelper.ErrResponse(c, http.StatusNotFound, "Web2Auth failed to create new user for email : "+email)
					return
				}

				customClaims := claims.NewUser(newUser.Email)
				pasetoToken, err := auth.GenerateTokenPasetoWeb2User(customClaims)
				if err != nil {
					httphelper.NewInternalServerError(c, "failed to generate token, error %v", err.Error())
					return
				}

				payload := AuthenticatePayload{
					Token: pasetoToken,
				}
				httphelper.SuccessResponse(c, "Token generated successfully", payload)

			} else if count == 0 && email != "" {
				result := db.Create(&newUser)
				if result.Error != nil {
					logwrapper.Errorf("failed to create new user, error %v", err)
					httphelper.ErrResponse(c, http.StatusBadGateway, "Web2Auth failed to create new user for email : "+email)
					return
				} else {
					customClaims := claims.NewUser(newUser.Email)
					pasetoToken, err := auth.GenerateTokenPasetoWeb2User(customClaims)
					if err != nil {
						httphelper.NewInternalServerError(c, "failed to generate token, error %v", err.Error())
						return
					}
					payload := AuthenticatePayload{
						Token: pasetoToken,
					}
					httphelper.SuccessResponse(c, "Token generated successfully", payload)
				}
			} else if email == "" {
				httphelper.ErrResponse(c, http.StatusNotFound, "Web2Auth failed to Authenticate jwt Token: "+email)
				return
			}
		}

	}
}
