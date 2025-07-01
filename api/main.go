package main

import (
	"log"

	"github.com/amirhnajafiz/aep/backend/bootstrap"
	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
)

func main() {
	// load configs
	cfg := configs.Load("config.yaml")

	// create a new database instance
	db, err := database.NewDatabase(cfg.Storage.URI())
	if err != nil {
		log.Fatalf("failed to initialize database: %v", err)
	}

	// initialize database entries with predefined URLs
	if err := bootstrap.SetupDatabase(cfg.Revision, cfg.AdminKey, db, cfg.URLs); err != nil {
		log.Fatalf("failed to initialize URL entries: %v", err)
	}

	// bootstrap the application
	if err := bootstrap.SetupApp(cfg.Port, db); err != nil {
		log.Fatalf("failed to bootstrap application: %v", err)
	}
}
