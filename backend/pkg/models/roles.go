package models

import "time"

// Role reperents a relationship between each user and its role.
type Role struct {
	Username  string    `pg:"username,unique:userrole"`
	Role      string    `pg:"role,unique:userrole"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
}
