package handler

import (
	"log"

	"github.com/amirhnajafiz/aep/backend/pkg/hashing"

	"github.com/labstack/echo/v4"
)

// admin middleware checks if the request has a valid admin key in the header.
func (h Handler) admin(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// get admin key from the request header
			adminKey := c.Request().Header.Get("X-Admin-Key")

			// check if the admin key matches the expected key
			if hashing.MD5Hash(adminKey) != key {
				log.Printf("unauthorized access attempt with key: %s\n", adminKey)
				return c.String(403, "forbidden")
			}

			return next(c)
		}
	}
}

// log middleware logs the request method, status, path, and request body size.
func (h Handler) log(handler string) echo.MiddlewareFunc {
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

			log.Printf("[%s] method: %s, status: %d, path: %s, request size: %d\n", handler, method, status, path, reqBodySize)

			return nil
		}
	}
}
