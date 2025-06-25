package models

// Session represents a session entry in the database.
type Session struct {
	ID        int64  `bun:"id,pk,autoincrement"`
	UrlID     int64  `bun:"url_id,notnull"`
	CreatedAt string `bun:"created_at,default:current_timestamp"`
	UpdatedAt string `bun:"updated_at,default:current_timestamp"`
	DeletedAt string `bun:"deleted_at,default:null"`
}
