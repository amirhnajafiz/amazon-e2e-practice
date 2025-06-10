package main

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func main() {
	// load configs
	cfg := configs.Load("config.yaml")

	// create a new database instance
	db, err := database.NewDatabase(cfg.Storage.URI())
	if err != nil {
		log.Fatal("failed to initialize database", err)
	}

	// initialize handler
	hd := handler.Handler{
		DB:  db,
		JWT: jwt.New(cfg.JWT.PrivateKey, cfg.JWT.ExpireTime),
	}

	// create a new Echo instance
	app := echo.New()

	// register endpoints
	hd.RegisterEndpoints(app)

	log.Printf("server is running on port %d\n", cfg.Port)

	// start the server
	if err := app.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
