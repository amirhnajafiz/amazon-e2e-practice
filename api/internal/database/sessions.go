package database

import (
	"context"

	"github.com/amirhnajafiz/aep/backend/internal/models"
)

// InsertSession inserts a new session into the database with the given URL ID.
func (db *Database) InsertSession(urlID int64) error {
	ctx := context.Background()

	_, err := db.conn.NewInsert().Model(&models.Session{
		UrlID: urlID,
	}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}

// GetTopUrlsBySessionCount retrieves the top URLs by session count, limited to the specified number.
func (db *Database) GetTopUrlsBySessionCount(limit int) ([]*models.Url, error) {
	ctx := context.Background()

	var urls []*models.Url
	err := db.conn.NewSelect().
		Model(&urls).
		Column("url.*").
		Join("JOIN sessions ON sessions.url_id = url.id").
		Group("url.id").
		OrderExpr("COUNT(sessions.id) DESC").
		Limit(limit).
		Scan(ctx)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

// GetSessionsCountByUrlID retrieves the count of sessions for a specific URL ID.
func (db *Database) GetSessionsCountByUrlID(urlID int64) (int, error) {
	ctx := context.Background()

	count, err := db.conn.NewSelect().
		Model(&models.Session{}).
		Where("url_id = ?", urlID).
		Count(ctx)
	if err != nil {
		return 0, err
	}

	return int(count), nil
}

// ClearSessions deletes all session records from the database.
func (db *Database) ClearSessions() error {
	ctx := context.Background()

	_, err := db.conn.NewDelete().Model(&models.Session{}).Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
