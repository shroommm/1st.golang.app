package entities

import "time"

// Comment represents a comment in database table.
type Comment struct {
	ID          int       `db:"id"`
	UserID      string    `db:"user_id"`
	PostID      string    `db:"post_id"`
	Description string    `db:"description"`
	Status      bool      `db:"status"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
