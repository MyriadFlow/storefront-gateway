package webapp

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/webapp")
	{
		g.Use(paseto.PASETO)
		g.POST("", DeployWebapp)
	}
}
func DeployWebapp(c *gin.Context) {
	storefrontId := c.Request.Header["storefrontId"]

	c.JSON(http.StatusOK, gin.H{})
}
