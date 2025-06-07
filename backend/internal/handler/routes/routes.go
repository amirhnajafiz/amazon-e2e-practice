package routes

import (
	"github.com/go-pg/pg/v10"
	"go.uber.org/zap"
)

// Routes is a struct that groups all route handlers.
type Routes struct {
	Auth   *AuthGroup
	Health *HealthGroup
	Users  *UsersGroup
}

// NewRoutes initializes and returns a new Routes instance with its sub-groups.
func NewRoutes(logr *zap.Logger, db *pg.DB) *Routes {
	return &Routes{
		Auth: &AuthGroup{
			DB:     db,
			Logger: logr.Named("auth"),
		},
		Health: &HealthGroup{},
		Users: &UsersGroup{
			DB:     db,
			Logger: logr.Named("users"),
		},
	}
}
