package delegateassetcreation

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract/rawtransaction"
	"github.com/MyriadFlow/storefront-gateway/generated/smartcontract/signatureSeries"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/delegateAssetCreation")
	{
		g.GET("", getAssets)
		g.POST("", deletegateAssetCreation)
		g.Use(paseto.PASETO)
		g.POST("/store", storeAsset)
	}
}

func deletegateAssetCreation(c *gin.Context) {
	var request DelegateAssetCreationRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	// if !canaccess.CanAccess(request.CreatorAddress) {
	// 	httphelper.ErrResponse(c, http.StatusForbidden, "this api not open to current wallet")
	// 	return
	// }
	creatorAddr := common.HexToAddress(request.CreatorAddress)
	abiS := signatureSeries.SignatureSeriesABI

	tx, err := rawtransaction.SendRawTransaction(abiS, "delegateAssetCreation", creatorAddr, request.MetaDataHash, request.RoyaltyPercentBasisPoint)

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

func storeAsset(c *gin.Context) {
	var request AssetStoreRequest
	err := c.BindJSON(&request)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	db := dbconfig.GetDb()
	asset := models.DelegateAsset{
		MetaDataHash:             request.MetaDataHash,
		RoyaltyPercentBasisPoint: request.RoyaltyPercentBasisPoint,
		StorefrontId:             request.StorefrontId,
		ContractAddress:          request.ContractAddress,
	}
	err = db.Model(&models.DelegateAsset{}).Create(&asset).Error
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to create asset in db", "failed to create asset in db: ", err.Error())
		return
	}
	httphelper.SuccessResponse(c, "asset created", nil)

}
func getAssets(c *gin.Context) {
	var assets []models.DelegateAsset
	contractAddress := c.Query("contractAddress")
	storefrontId := c.Query("storefrontId")
	if storefrontId == "" || contractAddress == "" {
		logwrapper.Error("Bad query params")
		c.JSON(http.StatusBadRequest, gin.H{"message": "bad request"})
		return

	}
	db := dbconfig.GetDb()

	err := db.Model(&models.DelegateAsset{}).Where("storefront_id = ? AND contract_address = ?", storefrontId, contractAddress).Find(&assets).Error
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to create asset in db", "failed to create asset in db: ", err.Error())
		return
	}
	httphelper.SuccessResponse(c, "assets fetched", assets)
}
