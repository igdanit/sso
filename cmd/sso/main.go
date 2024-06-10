package main

import (
	"os"

	appGRPC "github.com/igdanit/sso/internal/app"
	"github.com/igdanit/sso/internal/auth/logger"
	"github.com/igdanit/sso/internal/config"
	"google.golang.org/grpc"
)

func main() {
	// Initialize config
	appCfg := config.GetConfig()

	// Initialize Logger
	appLogger := logger.New(appCfg.Env, os.Stdout)

	appLogger.Info("Config initialized", "cfg", appCfg)
	appLogger.Info("App initialized")

	// Starting application
	app := appGRPC.New(appCfg.GRPC.Port, appLogger, grpc.NewServer())
	app.Run()
	appLogger.Warn("CHECK GRPC DOC to SERVER TYPE SERVE METHOD DESCIRPTION FOR PROPRELY SETUP KEEP-ALIVE PARAMETER")
}
