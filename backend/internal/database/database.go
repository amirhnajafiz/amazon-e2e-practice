package database

import (
	"context"
	"fmt"

	"github.com/amirhnajafiz/aep/backend/pkg/models"

	"github.com/uptrace/bun"
)

// Database struct holds the database connection and provides methods to interact with it.
type Database struct {
	conn *bun.DB
}

// NewDatabase initializes a new Database instance and creates the necessary tables if they do not exist.
func NewDatabase(uri string) (*Database, error) {
	// establish a connection to the PostgreSQL database
	db, err := openConnection(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to open database connection: %v", err)
	}

	// create the necessary tables if they do not exist
	interfaces := []interface{}{
		&models.User{},
	}
	for _, model := range interfaces {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(context.Background()); err != nil {
			return nil, fmt.Errorf("failed to create table for model %T: %v", model, err)
		}
	}

	return &Database{conn: db}, nil
}
