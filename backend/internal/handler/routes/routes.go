package routes

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
	"github.com/amirhnajafiz/aep/backend/pkg/jwt"
)

// Routes is a struct that groups all route handlers.
type Routes struct {
	Auth   *AuthGroup
	Health *HealthGroup
}

// NewRoutes initializes and returns a new Routes instance with its sub-groups.
func NewRoutes(db *database.Database, j *jwt.Auth) *Routes {
	return &Routes{
		Auth: &AuthGroup{
			DB:  db,
			JWT: j,
		},
		Health: &HealthGroup{},
	}
}
