package storage

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// NewConnection to the give Postgres database.
func NewConnection(uri string) (*pg.DB, error) {
	opt, err := pg.ParseURL(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	db := pg.Connect(opt)

	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}

// CreateTable creates tables for the given models in the database.
func CreateTables(db *pg.DB, models []interface{}) error {
	for _, model := range models {
		if err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			return fmt.Errorf("failed to create table: %v", err)
		}
	}

	return nil
}
