package routes

import "github.com/labstack/echo/v4"

type AuthGroup struct{}

func (a *AuthGroup) Signin(c echo.Context) error {
	// Implement the signin logic here
	return nil
}

func (a *AuthGroup) Signup(c echo.Context) error {
	// Implement the signup logic here
	return nil
}
