package models

import "time"

// Each user has a unique username and a password.
type User struct {
	ID        int64     `bun:"id,pk,autoincrement"`
	Username  string    `bun:"username,unique"`
	Password  string    `bun:"password"`
	Role      string    `bun:"role,default:'user'"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
	UpdatedAt time.Time `bun:"updated_at,default:current_timestamp"`
	DeletedAt time.Time `bun:"deleted_at,default:null"`
}
