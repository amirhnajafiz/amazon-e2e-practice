package database

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v10"
)

// openConnection establishes a connection to the PostgreSQL database using the provided URI.
func openConnection(uri string) (*pg.DB, error) {
	// parse the URI to get connection options
	opt, err := pg.ParseURL(uri)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// configure the connection options
	opt.PoolSize = 10    // set the maximum number of connections in the pool
	opt.ReadTimeout = 30 // set the read timeout to 30 seconds

	// create a new database connection and ping the database to ensure it's reachable
	db := pg.Connect(opt)
	ctx := context.Background()
	if err := db.Ping(ctx); err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return db, nil
}
