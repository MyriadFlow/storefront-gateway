package apiv1

import (
	authenticate "github.com/MyriadFlow/storefront_gateway/api/v1/authenticate"
	claimrole "github.com/MyriadFlow/storefront_gateway/api/v1/claimRole"
	delegateartifactcreation "github.com/MyriadFlow/storefront_gateway/api/v1/delegateArtifactCreation"
	"github.com/MyriadFlow/storefront_gateway/api/v1/details"
	flowid "github.com/MyriadFlow/storefront_gateway/api/v1/flowid"
	"github.com/MyriadFlow/storefront_gateway/api/v1/healthcheck"
	"github.com/MyriadFlow/storefront_gateway/api/v1/profile"
	roleid "github.com/MyriadFlow/storefront_gateway/api/v1/roleId"
	"github.com/MyriadFlow/storefront_gateway/api/v1/uploadtoipfs"
	"github.com/MyriadFlow/storefront_gateway/api/v1/products"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		flowid.ApplyRoutes(v1)
		authenticate.ApplyRoutes(v1)
		profile.ApplyRoutes(v1)
		roleid.ApplyRoutes(v1)
		claimrole.ApplyRoutes(v1)
		delegateartifactcreation.ApplyRoutes(v1)
		uploadtoipfs.ApplyRoutes(v1)
		details.ApplyRoutes(v1)
		healthcheck.ApplyRoutes(v1)
		products.ApplyRoutes(v1)
	}
}
