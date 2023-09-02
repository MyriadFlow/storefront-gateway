package subgraph

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subgraph")
	{
		g.Use(paseto.PASETO)
		g.GET("", GetSubgraphsByAddress)
	}
}

func GetSubgraphsByAddress(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var subgraphs []models.Subgraph
	err := db.Model(&models.Subgraph{}).Where("wallet_address = ?", walletAddress).Find(&subgraphs)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", subgraphs)
}
