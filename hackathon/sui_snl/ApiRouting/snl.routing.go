package suiRouter

import (
	"github.com/gin-gonic/gin"
)

func SuiApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("snl")
	{
		v1.GET("")
	}
}
