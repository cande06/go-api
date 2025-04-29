package sale

import "time"

type Sale struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"`
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

// Metadata me permite tener un orden definitivo para mostrar los datos
type Metadata struct {
	Quantity    int     `json:"quantity"`
	Approved    int     `json:"approved"`
	Rejected    int     `json:"rejected"`
	Pending     int     `json:"pending"`
	TotalAmount float32 `json:"total_amount"`
}
