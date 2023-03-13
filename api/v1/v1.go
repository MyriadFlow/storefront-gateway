package apiv1

import (
	authenticate "github.com/MyriadFlow/storefront-gateway/api/v1/authenticate"
	claimrole "github.com/MyriadFlow/storefront-gateway/api/v1/claimRole"
	delegateassetcreation "github.com/MyriadFlow/storefront-gateway/api/v1/delegateAssetCreation"
	"github.com/MyriadFlow/storefront-gateway/api/v1/details"
	flowid "github.com/MyriadFlow/storefront-gateway/api/v1/flowid"
	"github.com/MyriadFlow/storefront-gateway/api/v1/healthcheck"
	"github.com/MyriadFlow/storefront-gateway/api/v1/marketplace"
	"github.com/MyriadFlow/storefront-gateway/api/v1/profile"
	roleid "github.com/MyriadFlow/storefront-gateway/api/v1/roleId"
	//"github.com/MyriadFlow/storefront-gateway/api/v1/uploadtoipfs"
	"github.com/MyriadFlow/storefront-gateway/api/v1/nftstorage"
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
		delegateassetcreation.ApplyRoutes(v1)
		//uploadtoipfs.ApplyRoutes(v1)
		nftstorage.ApplyRoutes(v1)
		details.ApplyRoutes(v1)
		healthcheck.ApplyRoutes(v1)
		marketplace.ApplyRoutes(v1)
	}
}
