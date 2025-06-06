package model

// Metadata reperents a relationship between each user and its metadata.
type Metadata struct {
	Username string `pg:"username,unique:userrole"`
	Role     string `pg:"role,unique:userrole"`
}
