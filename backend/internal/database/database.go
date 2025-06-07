package database

import (
	"fmt"

	"github.com/amirhnajafiz/aep/backend/pkg/models"

	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

// Database struct holds the database connection and provides methods to interact with it.
type Database struct {
	conn *pg.DB
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
		&models.Role{},
	}
	for _, model := range interfaces {
		if err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		}); err != nil {
			return nil, fmt.Errorf("failed to create table: %v", err)
		}
	}

	return &Database{conn: db}, nil
}
