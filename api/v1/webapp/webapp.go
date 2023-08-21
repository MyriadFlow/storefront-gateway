package webapp

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
	g := r.Group("/webapp")
	{
		g.Use(paseto.PASETO)
		g.POST("", DeployWebapp)
	}
}
func DeployWebapp(c *gin.Context) {
	db := dbconfig.GetDb()
	storefrontId := c.GetHeader("StorefrontId")
	var subgraph models.Subgraph
	err := db.Where("storefront_id = ?", storefrontId).First(&subgraph).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	res := WebappResponse{
		SubgraphUrl: subgraph.SubgraphUrl,
	}
	c.JSON(http.StatusOK, res)
}
