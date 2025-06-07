package routes

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler/requests"
	"github.com/amirhnajafiz/aep/backend/pkg/hashing"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"
	"github.com/amirhnajafiz/aep/backend/pkg/models"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// AuthGroup handles authentication-related routes.
type AuthGroup struct {
	DB     *database.Database
	JWT    *jwt.Auth
	Logger *zap.Logger
}

func (a *AuthGroup) Signup(c echo.Context) error {
	// parse the request body json into a struct
	var req requests.UserRequest
	if err := c.Bind(&req); err != nil {
		a.Logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid request format"})
	}

	// validate the request
	if err := c.Validate(&req); err != nil {
		a.Logger.Warn("validation failed", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Validation failed"})
	}

	// hash the password
	hash := hashing.MD5Hash(req.Password)

	// store the user in the database
	user := models.User{
		Username: req.Username,
		Password: hash,
	}
	if count, err := a.DB.CreateUser(&user); err != nil {
		a.Logger.Error("failed to create user", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to create user"})
	} else {
		a.Logger.Info("number of users created", zap.Int("count", count))
	}

	return nil
}

func (a *AuthGroup) Signin(c echo.Context) error {
	// parse the request body json into a struct
	var req requests.UserRequest
	if err := c.Bind(&req); err != nil {
		a.Logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid request format"})
	}

	// validate the request
	if err := c.Validate(&req); err != nil {
		a.Logger.Warn("validation failed", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Validation failed"})
	}

	// hash the password
	hash := hashing.MD5Hash(req.Password)

	// check if the user exists in the database
	user, err := a.DB.GetUserByUsername(req.Username)
	if err != nil {
		a.Logger.Error("failed to get user", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to get user"})
	}
	if user == nil {
		a.Logger.Warn("user not found", zap.String("username", req.Username))
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	// check if the password matches
	if user.Password != hash {
		a.Logger.Warn("invalid password", zap.String("username", req.Username))
		return c.JSON(401, map[string]string{"error": "Invalid password"})
	}

	// build the JWT token
	token, err := a.JWT.GenerateJWT(user.Username, user.Role)
	if err != nil {
		a.Logger.Error("failed to generate JWT", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to generate JWT"})
	}

	return c.JSON(200, map[string]string{"message": "Signin successful", "token": token})
}
