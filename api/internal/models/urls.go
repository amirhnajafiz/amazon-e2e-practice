package models

// Url represents a URL entry in the database.
type Url struct {
	ID          int64  `bun:"id,pk,autoincrement"`
	ShortenedID string `bun:"shortened_id,notnull"`
	Address     string `bun:"address,notnull"`
	Description string `bun:"description"`
	CreatedAt   string `bun:"created_at,default:current_timestamp"`
	UpdatedAt   string `bun:"updated_at,default:current_timestamp"`
	DeletedAt   string `bun:"deleted_at,default:null"`
}
