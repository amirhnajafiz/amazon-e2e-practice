package main

import (
	"log"

	"github.com/amirhnajafiz/aep/backend/bootstrap"
	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"

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

	// initialize database entries with predefined URLs
	if err := bootstrap.SetupDatabase(cfg.Revision, cfg.AdminKey, db, cfg.URLs); err != nil {
		log.Fatal("failed to initialize URL entries", zap.Error(err))
	}

	// get the current migration's admin key
	mig, err := db.GetLastMigration()
	if err != nil || mig == nil {
		log.Fatal("failed to get last migration", zap.Error(err))
	}

	// bootstrap the application
	if err := bootstrap.SetupApp(cfg.Port, mig.AdminKey, db); err != nil {
		log.Fatal("failed to bootstrap application", zap.Error(err))
	}
}
