package routes

import "github.com/labstack/echo/v4"

type HealthGroup struct{}

func (h *HealthGroup) HealthCheck(c echo.Context) error {
	return c.String(200, "OK")
}
