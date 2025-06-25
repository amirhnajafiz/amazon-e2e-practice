package handler

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"

	"github.com/labstack/echo/v4"
)

// Handler struct contains the dependencies required for handling requests.
type Handler struct {
	DB *database.Database
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h Handler) RegisterEndpoints(app *echo.Echo) *echo.Echo {
	// create a new API group with metrics middleware
	_ = app.Group("/api", h.log())

	return app
}
