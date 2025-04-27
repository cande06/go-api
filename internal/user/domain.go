package user

import "time"

// User represents a system user with metadata for auditing and versioning.
type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	NickName  string    `json:"nickname"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Version   int       `json:"version"`
}

// UpdateFields represents the optional fields for updating a User.
// Puntero nil = no hay cambios en ese campo
type UpdateFields struct {
	Name     *string `json:"name"`
	Address  *string `json:"address"`
	NickName *string `json:"nickname"`
}
