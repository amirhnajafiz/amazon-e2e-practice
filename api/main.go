package main

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// initEntries initializes the database with predefined URL entries from the configuration.
func initEntries(db *database.Database, urls []configs.URLEntry) error {
	// clear existing URLs in the database
	if err := db.ClearUrls(); err != nil {
		return err
	}

	for _, url := range urls {
		if err := db.InsertUrl(url.Name, url.URL, url.Description); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	// load configs
	cfg := configs.Load("config.yaml")

	// create a new database instance
	db, err := database.NewDatabase(cfg.Storage.URI())
	if err != nil {
		log.Fatal("failed to initialize database", err)
	}

	// create a new Echo instance and register endpoints
	app := handler.Handler{
		DB: db,
	}.RegisterEndpoints(echo.New())

	// initialize database entries with predefined URLs
	if cfg.Migrate {
		if err := initEntries(db, cfg.URLs); err != nil {
			log.Fatal("failed to initialize URL entries", zap.Error(err))
		}
	}

	// start the server
	log.Printf("server is running on port %d\n", cfg.Port)
	if err := app.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
