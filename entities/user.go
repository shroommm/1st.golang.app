package entities

import "time"

type User struct {
	ID        int       `json:"id" db:"id"`
	FirstName string    `json:"firstname" db:"firstname"`
	LastName  string    `json:"lastname" db:"lastname"`
	Email     string    `json:"email" db:"email"`
	Status    bool      `json:"status" db:"status"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}
