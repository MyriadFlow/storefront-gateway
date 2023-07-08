package launchpad

import (
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
	}
}
func GetContracts(c *gin.Context) {

}
