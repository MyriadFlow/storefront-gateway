package app

import (
	"time"

	"github.com/MyriadFlow/storefront_gateway/api"
	"github.com/MyriadFlow/storefront_gateway/global"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
	"github.com/gin-contrib/cors"

	"github.com/MyriadFlow/storefront_gateway/config/dbconfig/dbinit"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/config/storefront"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init() {
	envconfig.InitEnvVars()
	gin.SetMode(envconfig.EnvVars.GIN_MODE)
	logwrapper.Init()
	dbinit.Init()
	global.InitGlobal()
	storefront.InitRolesId()
	GinApp = gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     envconfig.EnvVars.ALLOWED_ORIGIN})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
}
