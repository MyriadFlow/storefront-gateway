package products

import (
	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
)


// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/products")
	{
		g.Use(paseto.PASETO)
		g.GET("/trending", getTrending)
		g.GET("/highlights", getHighlights)
		g.GET("/marketplaceCounts", getMarketPlaceCounts)
	}
}

func getTrending(c *gin.Context) {

	payload:=Trending
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}

func getHighlights(c *gin.Context) {
	payload:=Highlights
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}

func getMarketPlaceCounts(c *gin.Context) {
	payload:=[]MarketPlaceCounts{{7878,4545,5546}}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}