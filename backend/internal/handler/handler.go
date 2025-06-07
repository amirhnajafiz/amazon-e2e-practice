package handler

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler/middlewares"
	"github.com/amirhnajafiz/aep/backend/internal/handler/routes"
	"github.com/amirhnajafiz/aep/backend/internal/telemetry/metrics"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Handler struct contains the dependencies required for handling requests.
type Handler struct {
	DB      *database.Database
	JWT     *jwt.Auth
	Logger  *zap.Logger
	Metrics *metrics.Metrics
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h *Handler) RegisterEndpoints(app *echo.Echo) {
	// create a new routes instance
	routes := routes.NewRoutes(h.Logger, h.DB)

	// register endpoints
	app.GET("/health", routes.Health.HealthCheck)

	// create a new API group with metrics middleware
	api := app.Group("/api", middlewares.ObserveMetrics(h.Metrics))

	// register authentication endpoints
	auth := api.Group("/auth")
	auth.POST("/auth/signin", routes.Auth.Signin)
	auth.POST("/auth/signup", routes.Auth.Signup)

	// register user management endpoints
	users := api.Group("/users", middlewares.JWT(h.JWT), middlewares.RoleCheck("admin"))
	users.GET("", routes.Users.ListUsers)
	users.GET("/:username", routes.Users.GetUser)
	users.PUT("/:username", routes.Users.UpdateUser)
	users.DELETE("/:username", routes.Users.DeleteUser)
}
