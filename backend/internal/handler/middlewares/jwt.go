package middlewares

import (
	"net/http"

	"github.com/amirhnajafiz/caaas/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// JWT is a middleware function that validates JWT tokens in incoming requests.
func JWT(auth *jwt.Auth) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// extract the JWT token from the request header
			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "missing or invalid JWT token")
			}

			// validate the JWT token
			user, role, err := auth.ParseJWT(token)
			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "invalid JWT token: "+err.Error())
			}

			// store user and role in the context for later use
			c.Set("user", user)
			c.Set("role", role)

			// croceed to the next handler if the token is valid
			return next(c)
		}
	}
}
