package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

// openConnection establishes a connection to the PostgreSQL database using the provided URI.
func openConnection(uri string) (*bun.DB, error) {
	// open a connection to the database
	db := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(uri)))

	// build a new Bun DB instance
	conn := bun.NewDB(db, pgdialect.New())
	if err := conn.Ping(); err != nil {
		return nil, err
	}

	return conn, nil
}
