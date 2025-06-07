package handler

import (
	"github.com/amirhnajafiz/caaas/internal/handler/middlewares"
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

func (h *Handler) healthCheck(c echo.Context) error {
	return c.String(200, "OK")
}

// RegisterEndpoints registers the API endpoints with the Echo framework.
func (h *Handler) RegisterEndpoints(app *echo.Echo) {
	// register endpoints
	app.GET("/health", h.healthCheck)

	api := app.Group("/api")

	// register authentication endpoints
	auth := api.Group("/auth")
	auth.POST("/auth/signin", h.signin)
	auth.POST("/auth/signup", h.signup)

	// register stats endpoint
	stats := api.Group("/stats", middlewares.JWT(h.JWT))
	stats.GET("", h.getStats)

	// register user management endpoints
	users := api.Group("/users", middlewares.JWT(h.JWT), middlewares.RoleCheck("admin"))
	users.GET("", h.listUsers)
	users.GET("/:username", h.getUser)
	users.PUT("/:username", h.updateUser)
	users.DELETE("/:username", h.deleteUser)
}
