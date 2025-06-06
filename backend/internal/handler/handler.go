package handler

import (
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
	// Register your endpoints here
	// For example:
	// app.GET("/api/v1/resource", h.GetResource)
	// app.POST("/api/v1/resource", h.CreateResource)
	// app.PUT("/api/v1/resource/:id", h.UpdateResource)
	// app.DELETE("/api/v1/resource/:id", h.DeleteResource)

}
