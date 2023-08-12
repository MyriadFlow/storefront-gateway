package webapp

import "github.com/gin-gonic/gin"

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/webapp")
	{
		g.POST("", DeployWebapp)
	}
}
func DeployWebapp(c *gin.Context) {

}
