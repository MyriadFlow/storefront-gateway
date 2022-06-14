package details

import (
	"github.com/TheLazarusNetwork/marketplace-engine/models/Org"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/details")
	{
		g.GET("", getDetails)
	}
}

func getDetails(c *gin.Context) {
	org, err := Org.GetOrgDetails()
	if err != nil {
		httphelper.InternalServerError(c)
		return
	}
	httphelper.SuccessResponse(c, "details successfully fetched", org)
}
