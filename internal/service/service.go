package service

import (
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"

	"github.com/DmitryPostolenko/XM_EX/internal/configuration"
)

// Configuration file
var configFile = "internal/configuration/configuration.yml"

// Start service
func Start() {
	// Echo initialization
	e := echo.New()
	e.Debug = true

	// Load configuration
	cfg, err := configuration.Load(configFile)
	if err != nil {
		panic(err)
	}

	// Routes
	initMiddleware(e, cfg)

	// Routes
	initRoutes(e, cfg)

	// Start listening
	e.Logger.Fatal(e.Start(":" + setUpPort(cfg.Server.Port)))
}
