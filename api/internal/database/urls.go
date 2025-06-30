package database

import (
	"context"

	"github.com/amirhnajafiz/aep/backend/internal/models"
)

// InsertUrl inserts a new URL into the database.
func (db *Database) InsertUrl(short, address, description string) error {
	ctx := context.Background()

	_, err := db.conn.NewInsert().Model(&models.Url{
		ShortenedID: short,
		Address:     address,
		Description: description,
	}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// DeleteUrl deletes a URL from the database by its ID.
func (db *Database) DeleteUrl(urlID int64) error {
	ctx := context.Background()

	_, err := db.conn.NewDelete().Model(&models.Url{}).
		Where("id = ?", urlID).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetUrls retrieves all URLs from the database.
func (db *Database) GetUrls() ([]*models.Url, error) {
	ctx := context.Background()

	var urls []*models.Url
	err := db.conn.NewSelect().Model(&urls).Scan(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// ClearUrls deletes all URLs from the database.
func (db *Database) ClearUrls() error {
	ctx := context.Background()

	_, err := db.conn.NewDelete().Model(&models.Url{}).Where("1=1").Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
