package bootstrap

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"

	"github.com/labstack/echo/v4"
)

// SetupApp initializes the application with the given port, admin key, and database.
func SetupApp(
	port int,
	db *database.Database,
) error {
	// get the current migration's admin key
	mig, err := db.GetLastMigration()
	if err != nil || mig == nil {
		return fmt.Errorf("failed to get last migration: %w", err)
	}

	// create a new Echo instance and register endpoints
	app := handler.Handler{
		AdminKey: mig.AdminKey,
		DB:       db,
	}.RegisterEndpoints(echo.New())

	// start the server
	log.Printf("server is running on port %d\n", port)
	if err := app.Start(fmt.Sprintf(":%d", port)); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}

	return nil
}
