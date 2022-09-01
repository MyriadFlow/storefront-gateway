package main

import (
	"fmt"

	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/config/envconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
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
