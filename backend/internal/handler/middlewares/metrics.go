package middlewares

import (
	"fmt"
	"time"

	"github.com/amirhnajafiz/aep/backend/internal/telemetry/metrics"

	"github.com/labstack/echo/v4"
)

// MetricsMiddleware logs status, latency, method, and path for each request.
func ObserveMetrics(mtx *metrics.Metrics) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			err := next(c)
			stop := time.Now()

			latency := stop.Sub(start)
			status := c.Response().Status
			method := c.Request().Method
			path := c.Request().URL.Path

			mtx.AddRequest(path, method)
			mtx.ObserveLatency(path, method, float64(latency.Milliseconds()))
			mtx.AddStatusCode(path, method, fmt.Sprintf("%d", status))

			return err
		}
	}
}
