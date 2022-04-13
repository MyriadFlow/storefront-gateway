package main

import (
	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/envutil"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
)

func main() {
	app.Init(".env", "logs")
	logwrapper.Log.Info("Starting app")
	err := app.GinApp.Run(":" + envutil.MustGetEnv("APP_PORT"))
	if err != nil {
		logwrapper.Log.Fatal(err)
	}
}
