package webapp

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/webapp")
	{
		g.GET("/:id", DeployWebapp)
		g.GET("/contracts/:id", GetContractAddresses)
	}
}
func DeployWebapp(c *gin.Context) {
	db := dbconfig.GetDb()
	storefrontId := c.Param("id")
	var storefront models.Storefront
	err := db.Model(&models.Storefront{}).Where("id = ?", storefrontId).First(&storefront).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	var tradehub Contract
	if result := db.Model(models.Contract{}).Where("storefront_id = ? AND contract_name = ?", storefrontId, "TradeHub").Find(&tradehub); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	var accessMaster Contract
	if result := db.Model(models.Contract{}).Where("storefront_id = ? AND contract_name = ?", storefrontId, "AccessMaster").Find(&accessMaster); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	var profile models.User
	if result := db.Model(models.User{}).Where("wallet_address", storefront.WalletAddress).Find(&profile); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	res := WebappResponse{
		Storefront:          storefront,
		TradehubAddress:     tradehub.ContractAddress,
		AccessMasterAddress: accessMaster.ContractAddress,
		BaseUrlGateway:      envconfig.EnvVars.BASE_URL_GATEWAY,
		IpfsGateway:         envconfig.EnvVars.IPFS_GATEWAY,
		Profile:             profile,
	}
	c.JSON(http.StatusOK, res)
}

func GetContractAddresses(c *gin.Context) {
	storefrontId := c.Param("id")
	db := dbconfig.GetDb()
	var contracts []Contract
	if result := db.Model(models.Contract{}).Where("storefront_id = ?", storefrontId).Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)

}
