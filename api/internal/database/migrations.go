package database

import (
	"context"

	"github.com/amirhnajafiz/aep/backend/internal/models"
)

// InsertMigration inserts a new migration entry into the database with the given revision and admin key.
func (db *Database) InsertMigration(revision int64, adminKey string) error {
	ctx := context.Background()

	_, err := db.conn.NewInsert().Model(&models.Migration{
		Revision: revision,
		AdminKey: adminKey,
	}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetLastMigration retrieves the last migration entry from the database, ordered by revision in descending order.
func (db *Database) GetLastMigration() (*models.Migration, error) {
	ctx := context.Background()

	var migration models.Migration
	err := db.conn.NewSelect().
		Model(&migration).
		OrderExpr("revision DESC").
		Where("deleted_at IS NULL").
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return &migration, nil
}

// DeleteMigration marks a migration entry as deleted by setting the deleted_at timestamp to the current time.
func (db *Database) DeleteMigration(revision int64) error {
	ctx := context.Background()

	_, err := db.conn.NewUpdate().
		Model(&models.Migration{}).
		Set("deleted_at = NOW()").
		Where("revision = ?", revision).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
