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
			label := fmt.Sprintf("%s-%s", method, path)

			mtx.AddRequest(label)
			mtx.ObserveLatency(label, float64(latency.Milliseconds()))
			if status >= 300 {
				mtx.AddFailedCall(label)
			}

			return err
		}
	}
}
