package product

import "time"

type (
	Product struct {
		ID          int32
		Name        string
		Price       float64
		Description string
		Category    string
		CreatedAt   time.Time
		UpdatedAt   time.Time
	}
)
