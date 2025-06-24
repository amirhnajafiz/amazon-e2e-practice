package routes

import (
	"github.com/amirhnajafiz/aep/backend/internal/database"
)

// Routes is a struct that groups all route handlers.
type Routes struct {
	Health *HealthGroup
}

// NewRoutes initializes and returns a new Routes instance with its sub-groups.
func NewRoutes(_ *database.Database) *Routes {
	return &Routes{
		Health: &HealthGroup{},
	}
}
