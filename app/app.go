package app

import (
	"time"

	"github.com/TheLazarusNetwork/marketplace-engine/api"
	"github.com/TheLazarusNetwork/marketplace-engine/config"
	"github.com/TheLazarusNetwork/marketplace-engine/global"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/gin-contrib/cors"

	"github.com/TheLazarusNetwork/marketplace-engine/config/creatify"
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig/dbinit"
	"github.com/gin-gonic/gin"
)

var GinApp *gin.Engine

func Init(envPath string, logBasePath string) {
	config.Init(envPath)
	logwrapper.Init(logBasePath)
	dbinit.Init()
	global.InitGlobal()
	creatify.InitRolesId()
	GinApp = gin.Default()

	corsM := cors.New(cors.Config{AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
		AllowOrigins:     []string{envutil.MustGetEnv("ALLOWED_ORIGIN")}})
	GinApp.Use(corsM)
	api.ApplyRoutes(GinApp)
}
