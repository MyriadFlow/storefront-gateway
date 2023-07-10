package launchpad

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/launchpad")
	{
		g.POST("/AccessMaster", DeployAccessMaster)
		g.POST("/TradeHub", DeployTradeHub)
		g.POST("/FusionSeries", DeployFusionSeries)
		g.POST("/SignatureSeries", DeploySignatureSeries)
		g.POST("/InstaGen", DeployInstaGen)
		g.POST("/EternumPass", DeployEternumPass)
		g.GET("/contracts", GetContracts)
		g.GET("/contracts/:contractName", GetContractsByName)
	}
}
func GetContracts(c *gin.Context) {
	db := dbconfig.GetDb()
	var contracts []models.Contract
	if result := db.Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)
}

func GetContractsByName(c *gin.Context) {
	contractName := c.Param("contractName")
	db := dbconfig.GetDb()
	var contracts []models.Contract
	if result := db.Where("contract_name = ?", contractName).Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)
}
