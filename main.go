package main

import (
	"fmt"

	"github.com/MyriadFlow/storefront-gateway/app"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app.Init()
	logwrapper.Log.Info("Starting app")
	addr := fmt.Sprintf(":%d", envconfig.EnvVars.APP_PORT)
	err := app.GinApp.Run(addr)
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
