package main

import (
	"fmt"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"
	"github.com/amirhnajafiz/aep/backend/internal/storage"
	"github.com/amirhnajafiz/aep/backend/internal/telemetry/logger"
	"github.com/amirhnajafiz/aep/backend/internal/telemetry/metrics"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	// load configs
	cfg := configs.Load("config.yaml")

	// initialize logger
	logr := logger.NewLogger(cfg.Logger.Level)

	// open storage connection
	stg, err := storage.NewConnection(cfg.Storage.URI())
	if err != nil {
		logr.Fatal("failed to connect to database", zap.Error(err))
	}

	// create a new database instance
	db, err := database.NewDatabase(stg)
	if err != nil {
		logr.Fatal("failed to initialize database", zap.Error(err))
	}

	// initialize JWT auth
	auth := jwt.New(cfg.JWT.PrivateKey, cfg.JWT.ExpireTime)

	// initialize metrics
	mtx := metrics.NewMetrics()
	if cfg.Metrics.Enabled {
		logr.Info("metrics are enabled", zap.Int("port", cfg.Metrics.Port))
		metrics.StartServer(cfg.Metrics.Port)
	}

	// initialize handler
	hd := handler.Handler{
		DB:      db,
		JWT:     auth,
		Logger:  logr,
		Metrics: mtx,
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
