package handler

import (
	"log"

	"github.com/labstack/echo/v4"
)

func (h Handler) Log() echo.MiddlewareFunc {
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
			reqBodySize := c.Request().ContentLength

			log.Printf("[REQ] method: %s, status: %d, path: %s, request size: %d\n", method, status, path, reqBodySize)
			log.Printf("[HED] headers: %v\n", c.Request().Header)

			return nil
		}
	}
}
