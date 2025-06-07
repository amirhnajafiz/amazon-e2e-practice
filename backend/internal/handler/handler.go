package handler

import (
	"github.com/amirhnajafiz/caaas/internal/handler/middlewares"
	"github.com/amirhnajafiz/caaas/internal/handler/routes"
	"github.com/amirhnajafiz/caaas/internal/telemetry/metrics"
	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/go-pg/pg/v10"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// Handler struct contains the dependencies required for handling requests.
type Handler struct {
	DB      *pg.DB
	JWT     *jwt.Auth
	Logger  *zap.Logger
	Metrics *metrics.Metrics
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h *Handler) RegisterEndpoints(app *echo.Echo) {
	// create a new routes instance
	routes := routes.NewRoutes()

	// register endpoints
	app.GET("/health", groups.Health.HealthCheck)

	api := app.Group("/api")

	// register authentication endpoints
	auth := api.Group("/auth")
	auth.POST("/auth/signin", groups.Auth.Signin)
	auth.POST("/auth/signup", groups.Auth.Signup)

	// register user management endpoints
	users := api.Group("/users", middlewares.JWT(h.JWT), middlewares.RoleCheck("admin"))
	users.GET("", groups.Users.ListUsers)
	users.GET("/:username", groups.Users.GetUser)
	users.PUT("/:username", groups.Users.UpdateUser)
	users.DELETE("/:username", groups.Users.DeleteUser)
}
