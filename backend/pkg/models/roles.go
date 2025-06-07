package models

import "time"

// Role reperents a relationship between each user and its role.
type Role struct {
	Role        string    `pg:"role,unique"`
	Description string    `pg:"description"`
	CreatedAt   time.Time `pg:"created_at,default:now()"`
	UpdatedAt   time.Time `pg:"updated_at,default:now()"`
}

// UserRole reperents a relationship between each user and its role.
type UserRole struct {
	Username  string    `pg:"username,unique:userrole"`
	Role      string    `pg:"role,unique:userrole"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
	DeletedAt time.Time `pg:"deleted_at"`
}
