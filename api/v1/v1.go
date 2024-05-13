package apiv1

import (
	"fmt"

	authenticate "github.com/MyriadFlow/storefront-gateway/api/v1/authenticate"
	creatorrole "github.com/MyriadFlow/storefront-gateway/api/v1/creatorrole"
	delegateassetcreation "github.com/MyriadFlow/storefront-gateway/api/v1/delegateassetcreation"
	"github.com/MyriadFlow/storefront-gateway/api/v1/launchpad"
	"github.com/MyriadFlow/storefront-gateway/api/v1/storefront"
	"github.com/MyriadFlow/storefront-gateway/api/v1/subgraph"
	"github.com/MyriadFlow/storefront-gateway/api/v1/webapp"
	"github.com/MyriadFlow/storefront-gateway/api/v1/website"
	suiRouter "github.com/MyriadFlow/storefront-gateway/hackathon/sui_snl/ApiRouting"
	voyagerrouting "github.com/MyriadFlow/storefront-gateway/hackathon/voyager/voyagerRouting"

	"github.com/MyriadFlow/storefront-gateway/api/v1/healthcheck"
	"github.com/MyriadFlow/storefront-gateway/api/v1/highlights"
	"github.com/MyriadFlow/storefront-gateway/api/v1/likes"
	"github.com/MyriadFlow/storefront-gateway/api/v1/nftstorage"
	"github.com/MyriadFlow/storefront-gateway/api/v1/profile"

	"github.com/MyriadFlow/storefront-gateway/api/v1/wishlist"
	"github.com/gin-gonic/gin"
)

// ApplyRoutes Use the given Routes
func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1.0")
	{
		authenticate.ApplyRoutes(v1)
		profile.ApplyRoutes(v1)

		creatorrole.ApplyRoutes(v1)
		delegateassetcreation.ApplyRoutes(v1)
		nftstorage.ApplyRoutes(v1)

		healthcheck.ApplyRoutes(v1)
		highlights.ApplyRoutes(v1)
		likes.ApplyRoutes(v1)
		wishlist.ApplyRoutes(v1)
		launchpad.ApplyRoutes(v1)
		storefront.ApplyRoutes(v1)
		subgraph.ApplyRoutes(v1)
		webapp.ApplyRoutes(v1)
		website.ApplyRoutes(v1)
		fmt.Println("************* HACKATHONE APIS ***************")
		voyagerrouting.VoyagerApplyRoutes(v1)
		suiRouter.SuiApplyRoutes(v1)
	}
}
