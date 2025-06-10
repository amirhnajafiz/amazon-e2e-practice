package main

import (
	"fmt"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"
	"github.com/amirhnajafiz/aep/backend/internal/logger"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	// load configs
	cfg := configs.Load("config.yaml")

	// initialize logger
	logr := logger.NewLogger(cfg.Logger.Level)

	// create a new database instance
	db, err := database.NewDatabase(cfg.Storage.URI())
	if err != nil {
		logr.Fatal("failed to initialize database", zap.Error(err))
	}

	// initialize handler
	hd := handler.Handler{
		DB:     db,
		JWT:    jwt.New(cfg.JWT.PrivateKey, cfg.JWT.ExpireTime),
		Logger: logr,
	}

	// create a new Echo instance
	app := echo.New()

	// register endpoints
	hd.RegisterEndpoints(app)

	logr.Info("starting server", zap.Int("port", cfg.Port))

	// start the server
	if err := app.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		hd.Logger.Fatal("failed to start server", zap.Error(err))
	}
}
