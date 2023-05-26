package order

import (
	"time"
)

type (
	Payment struct {
		ID        string    `json:"id"`
		VaNumber  string    `json:"va_number"`
		Amount    float64   `json:"amount"`
		Status    string    `json:"status"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	OrderItem struct {
		ID        int32     `json:"id"`
		ProductID int32     `json:"product_id"`
		Quantity  int32     `json:"quantity"`
		Name      string    `json:"name"`
		Price     float64   `json:"price"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	Order struct {
		ID         int32       `json:"id"`
		Payment    Payment     `json:"payment"`
		OrderItems []OrderItem `json:"order_items"`
		Status     string      `json:"status"`
		CreatedAt  time.Time   `json:"created_at"`
		UpdatedAt  time.Time   `json:"updated_at"`
	}
)
