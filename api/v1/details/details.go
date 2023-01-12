package details

import (
	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/jwt"
	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/onlyoperator"
	"github.com/MyriadFlow/storefront_gateway/models/Org"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/details")
	{
		g.GET("", getDetails)
		g.Use(jwt.JWT)
		g.Use(onlyoperator.OnlyOperator)
		g.POST("", postDetails)
	}
}

func postDetails(c *gin.Context) {
	var org Org.Org
	err := c.BindJSON(&org)
	if err != nil {
		logwrapper.Errorf("failed to parse body, err: %s", err)
		return
	}
	err = Org.UpdateOrg(org)

	if err != nil {
		logwrapper.Errorf("failed to update org, err: %s", err)
		httphelper.InternalServerError(c)
		return
	}

	httphelper.SuccessResponse(c, "successfully updated org", nil)
}

func getDetails(c *gin.Context) {
	org, err := Org.GetOrgDetails()
	if err != nil {
		httphelper.InternalServerError(c)
		return
	}
	httphelper.SuccessResponse(c, "details successfully fetched", org)
}
