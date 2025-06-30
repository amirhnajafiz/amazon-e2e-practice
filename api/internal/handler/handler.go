package handler

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"

	"github.com/labstack/echo/v4"
)

// Handler struct contains the dependencies required for handling requests.
type Handler struct {
	AdminKey string
	DB       *database.Database
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h Handler) RegisterEndpoints(app *echo.Echo) *echo.Echo {
	// create a new API group with metrics middleware
	api := app.Group("/api", h.log("API"))

	// register the endpoints for the API group
	api.GET("/urls", h.getUrls)
	api.PUT("/urls", h.insertSession)
	api.GET("/urls/:url_id", h.getSessionsCount)
	api.GET("/stats", h.getTopUrls)

	// register the admin endpoints
	admin := app.Group("/admin", h.admin(), h.log("Admin"))

	admin.POST("/urls", h.createUrl)
	admin.DELETE("/urls/:url_id", h.deleteUrl)

	return app
}
