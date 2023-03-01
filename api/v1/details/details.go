package details

import (
	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/onlyoperator"
	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/models/Org"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/details")
	{
		g.GET("", getDetails)
		g.Use(paseto.PASETO)
		g.Use(onlyoperator.OnlyOperator)
		//g.POST("", postDetails)
		g.PATCH("", patchDetails)
	}
}

// func postDetails(c *gin.Context) {
func patchDetails(c *gin.Context) {
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
