package routes

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler/requests"
	"github.com/amirhnajafiz/aep/backend/internal/models"
	"github.com/amirhnajafiz/aep/backend/pkg/hashing"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"

	"github.com/labstack/echo/v4"
)

// AuthGroup handles authentication-related routes.
type AuthGroup struct {
	DB  *database.Database
	JWT *jwt.Auth
}

func (a *AuthGroup) Signup(c echo.Context) error {
	// parse the request body json into a struct
	var req requests.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request format"})
	}

	// validate the request
	if err := req.Validate(); err != nil {
		return c.JSON(400, map[string]string{"error": "Validation failed"})
	}

	// hash the password
	hash := hashing.MD5Hash(req.Password)

	// store the user in the database
	user := models.User{
		Username: req.Username,
		Password: hash,
	}
	if _, err := a.DB.CreateUser(&user); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	}

	return c.JSON(201, map[string]string{"message": "User created successfully"})
}

func (a *AuthGroup) Signin(c echo.Context) error {
	// parse the request body json into a struct
	var req requests.UserRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request format"})
	}

	// validate the request
	if err := req.Validate(); err != nil {
		return c.JSON(400, map[string]string{"error": "Validation failed"})
	}

	// hash the password
	hash := hashing.MD5Hash(req.Password)

	// check if the user exists in the database
	user, err := a.DB.GetUserByUsername(req.Username)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to get user"})
	}

	// check if the password matches
	if user.Password != hash {
		return c.JSON(401, map[string]string{"error": "Invalid password"})
	}

	// build the JWT token
	token, err := a.JWT.GenerateJWT(user.Username)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to generate JWT"})
	}

	return c.JSON(200, map[string]string{"message": "Signin successful", "token": token})
}

// TokenCheck checks the validity of the JWT token.
func (a *AuthGroup) TokenCheck(c echo.Context) error {
	return c.JSON(200, map[string]string{"message": "Token is valid", "username": c.Get("user").(string)})
}
