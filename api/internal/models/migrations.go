package models

// Migration represents a migration entry in the database.
type Migration struct {
	ID        int64  `bun:"id,pk,autoincrement"`
	Version   string `bun:"version,notnull"`
	AdminKey  string `bun:"admin_key,notnull"`
	CreatedAt string `bun:"created_at,default:current_timestamp"`
	UpdatedAt string `bun:"updated_at,default:current_timestamp"`
	DeletedAt string `bun:"deleted_at,default:null"`
}
