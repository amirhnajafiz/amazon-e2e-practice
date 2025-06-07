package model

import "time"

// Each user has a unique username and a password.
type User struct {
	Username  string    `pg:"username,unique"`
	Password  string    `pg:"password"`
	Age       int       `pg:"age"`
	Email     string    `pg:"email,unique"`
	Phone     string    `pg:"phone,unique"`
	Address   string    `pg:"address"`
	ZIPCode   string    `pg:"zip_code"`
	FirstName string    `pg:"first_name"`
	LastName  string    `pg:"last_name"`
	Salary    float64   `pg:"salary"`
	JobTitle  string    `pg:"job_title"`
	CreatedAt time.Time `pg:"created_at,default:now()"`
	UpdatedAt time.Time `pg:"updated_at,default:now()"`
	DeletedAt time.Time `pg:"deleted_at"`
}
