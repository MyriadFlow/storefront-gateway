package webapp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/webapp")
	{
		g.POST("", DeployWebapp)
	}
}
func DeployWebapp(c *gin.Context) {
	storefrontId := c.Request.Header["storefrontId"]

	c.JSON(http.StatusOK, gin.H{})
}
