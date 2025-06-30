package bootstrap

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/configs"
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/pkg/hashing"
)

// SetupDatabase initializes the database with the given revision, admin key, and URLs.
func SetupDatabase(
	revision int,
	key string,
	db *database.Database,
	urls []configs.URLEntry,
) error {
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

	// delete the last migration entry if it exists
	if lastMigration != nil {
		if err := db.DeleteMigration(lastMigration.Revision); err != nil {
			return fmt.Errorf("failed to delete last migration entry: %w", err)
		}
		log.Printf("deleted last migration entry with revision %d\n", lastMigration.Revision)
	}

	// insert the migration entry
	if err := db.InsertMigration(int64(revision), hashing.MD5Hash(key)); err != nil {
		return fmt.Errorf("failed to insert migration entry: %w", err)
	}

	log.Printf("database initialized with %d URLs and migration revision %d\n", len(urls), revision)

	return nil
}
