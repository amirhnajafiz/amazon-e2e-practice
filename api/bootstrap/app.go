package bootstrap

import (
	"fmt"
	"log"

	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// SetupApp initializes the application with the given port, admin key, and database.
func SetupApp(
	port int,
	db *database.Database,
) error {
	// get the current migration's admin key
	mig, err := db.GetLastMigration()
	if err != nil || mig == nil {
		log.Fatal("failed to get last migration", zap.Error(err))
	}

	// create a new Echo instance and register endpoints
	app := handler.Handler{
		DB: db,
	}.RegisterEndpoints(mig.AdminKey, echo.New())

	// start the server
	log.Printf("server is running on port %d\n", port)
	if err := app.Start(fmt.Sprintf(":%d", port)); err != nil {
		log.Fatal("failed to start server", zap.Error(err))
	}

	return nil
}
