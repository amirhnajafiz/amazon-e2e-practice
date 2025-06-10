package handler

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler/middlewares"
	"github.com/amirhnajafiz/aep/backend/internal/handler/routes"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// Handler struct contains the dependencies required for handling requests.
type Handler struct {
	DB  *database.Database
	JWT *jwt.Auth
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h *Handler) RegisterEndpoints(app *echo.Echo) {
	// create a new routes instance
	routes := routes.NewRoutes(h.DB, h.JWT)

	// register endpoints
	app.GET("/health", routes.Health.HealthCheck)

	// create a new API group with metrics middleware
	api := app.Group("/api")

	// register authentication endpoints
	auth := api.Group("/auth")
	auth.POST("/signin", routes.Auth.Signin)
	auth.POST("/signup", routes.Auth.Signup)

	// register user management endpoints
	users := api.Group("/users", middlewares.JWT(h.JWT), middlewares.RoleCheck("admin"))
	users.GET("", routes.Users.ListUsers)
	users.PUT("", routes.Users.UpdateUser)
	users.DELETE("/:username", routes.Users.DeleteUser)
}
