package app

import (
	"time"

	"github.com/MyriadFlow/storefront-gateway/api"
	"github.com/MyriadFlow/storefront-gateway/global"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/auth"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/static"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig/dbinit"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/config/storefront"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	gin.SetMode(envconfig.EnvVars.APP_MODE)
	auth.Init()
	logwrapper.Init()
	dbinit.Init()
	global.InitGlobal()
	storefront.InitRolesId()
	GinApp = gin.Default()

	corsM := cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		ExposeHeaders:    []string{"*"},
		AllowOrigins:     []string{"*"},
	})

	// serve static files
	GinApp.Use(static.Serve("/", static.LocalFile("./web", false)))
	GinApp.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"status": 404, "message": "Invalid Endpoint Request"})
	})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
}
