package main

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"
	"github.com/amirhnajafiz/aep/backend/pkg/hashing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// initEntries initializes the database with predefined URL entries from the configuration.
func initEntries(revision int, db *database.Database, key string, urls []configs.URLEntry) error {
	// get the last migration entry
	lastMigration, err := db.GetLastMigration()
	if err != nil {
		return fmt.Errorf("failed to get last migration: %w", err)
	}

	// check if the last migration revision matches the current revision
	if lastMigration != nil && lastMigration.Revision >= int64(revision) {
		log.Printf("database is already initialized with the latest migration (revision %d)\n", revision)
		return nil
	}

	// clear existing URLs in the database
	if err := db.ClearUrls(); err != nil {
		return fmt.Errorf("failed to clear URLs: %w", err)
	}

	// clear existing sessions in the database
	if err := db.ClearSessions(); err != nil {
		return fmt.Errorf("failed to clear sessions: %w", err)
	}

	for _, url := range urls {
		if err := db.InsertUrl(url.Name, url.URL, url.Description); err != nil {
			return fmt.Errorf("failed to insert URL: %w", err)
		}
	}

	// insert the migration entry
	if err := db.InsertMigration(int64(revision), hashing.MD5Hash(key)); err != nil {
		return fmt.Errorf("failed to insert migration entry: %w", err)
	}

	log.Printf("database initialized with %d URLs and migration revision %d\n", len(urls), revision)

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
	if err := initEntries(cfg.Revision, db, cfg.AdminKey, cfg.URLs); err != nil {
		log.Fatal("failed to initialize URL entries", zap.Error(err))
	}

	// start the server
	log.Printf("server is running on port %d\n", cfg.Port)
	if err := app.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}
}
