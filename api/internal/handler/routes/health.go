package routes

import "github.com/labstack/echo/v4"

// HealthGroup handles health check routes.
type HealthGroup struct{}

func (h *HealthGroup) HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
