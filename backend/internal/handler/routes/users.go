package routes

import (
	"time"

	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/internal/handler/requests"
	"github.com/amirhnajafiz/aep/backend/internal/handler/responses"
	"github.com/amirhnajafiz/aep/backend/pkg/hashing"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

// UsersGroup handles user management routes.
type UsersGroup struct {
	DB     *database.Database
	Logger *zap.Logger
}

func (u *UsersGroup) ListUsers(c echo.Context) error {
	// get all users from the database
	users, err := u.DB.GetAllUsers()
	if err != nil {
		u.Logger.Error("failed to get users", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to retrieve users"})
	}

	// convert users to a response format
	var userResponses []responses.UserResponse
	for _, user := range users {
		userResponses = append(userResponses, responses.UserResponse{
			Username:  user.Username,
			Role:      user.Role,
			CreatedAt: user.CreatedAt.String(),
			UpdatedAt: user.UpdatedAt.String(),
			DeletedAt: user.DeletedAt.String(),
		})
	}

	return c.JSON(200, userResponses)
}

func (u *UsersGroup) UpdateUser(c echo.Context) error {
	// parse the request body json into a struct
	var req requests.UserUpdateRequest
	if err := c.Bind(&req); err != nil {
		u.Logger.Error("failed to bind request", zap.Error(err))
		return c.JSON(400, map[string]string{"error": "Invalid request format"})
	}

	// hash the password
	var hash string
	if req.Password != "" {
		hash = hashing.MD5Hash(req.Password)
	}

	// get the user by username
	user, err := u.DB.GetUserByUsername(req.Username)
	if err != nil {
		u.Logger.Error("failed to get user", zap.Error(err))
		return c.JSON(404, map[string]string{"error": "User not found"})
	}

	// update the user details
	user.Password = hash
	user.Role = req.Role
	user.UpdatedAt = time.Now()

	// save the updated user back to the database
	if count, err := u.DB.UpdateUser(user); err != nil {
		u.Logger.Error("failed to update user", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to update user"})
	} else {
		u.Logger.Info("number of users updated", zap.Int64("count", count))
	}

	return c.JSON(201, map[string]string{"message": "User updated successfully"})
}

func (u *UsersGroup) DeleteUser(c echo.Context) error {
	// get the username from the path parameter
	username := c.Param("username")
	if username == "" {
		u.Logger.Warn("username is required")
		return c.JSON(400, map[string]string{"error": "Username is required"})
	}

	// delete the user from the database
	if count, err := u.DB.DeleteUser(username); err != nil {
		u.Logger.Error("failed to delete user", zap.Error(err))
		return c.JSON(500, map[string]string{"error": "Failed to delete user"})
	} else {
		u.Logger.Info("number of users deleted", zap.Int64("count", count))
	}

	return c.JSON(201, map[string]string{"message": "User deleted successfully"})
}
