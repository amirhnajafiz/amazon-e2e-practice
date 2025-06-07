package models

import "time"

// Each user has a unique username and a password.
type User struct {
	Username  string    `pg:"username,unique"`
	Password  string    `pg:"password"`
	Role      string    `pg:"role,default:'user'"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
	DeletedAt time.Time `pg:"deleted_at"`
}
