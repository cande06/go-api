package user

import "time"

// User represents a system user with metadata for auditing and versioning.
type Sale struct {
	ID        string    `json:"id"`
	User_id   string    `json:"user_id"`
	Amount    float32   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

// UpdateFields se usa para actualizar campos del struct
type UpdateFields struct {
	Status *string `json:"status"`
}
