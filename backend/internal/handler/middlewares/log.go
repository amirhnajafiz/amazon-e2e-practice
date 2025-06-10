package middlewares

import (
	"log"

	"github.com/labstack/echo/v4"
)

// Log is a middleware function that logs the details of each request.
func Log() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// call the next handler in the chain
			if err := next(c); err != nil {
				c.Error(err)
			}

			// log the request details
			method := c.Request().Method
			path := c.Request().URL.Path
			status := c.Response().Status

			log.Printf("Method: %s, Status: %d, Path: %s,\n", method, status, path)

			return nil
		}
	}
}
