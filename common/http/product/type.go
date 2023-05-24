package product

import "time"

type (
	Product struct {
		ID          int32     `json:"id"`
		Name        string    `json:"name"`
		Price       float64   `json:"price"`
		Description string    `json:"description"`
		Category    string    `json:"category"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}
)
