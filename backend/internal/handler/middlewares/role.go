package middlewares

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

// RoleCheck is a middleware function that checks if the user's role is allowed to access the endpoint.
func RoleCheck(allowedRoles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			// retrieve the role from the context
			role := c.Get("role")
			if role == nil {
				return echo.NewHTTPError(http.StatusForbidden, "role not found in context")
			}

			// check if the role is allowed
			for _, allowedRole := range allowedRoles {
				if role == allowedRole {
					return next(c)
				}
			}

			return echo.NewHTTPError(http.StatusForbidden, "access denied for role: "+role.(string))
		}
	}
}
