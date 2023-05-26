package order

import "time"

const (
	OrderStatusPending    OrderStatus = "Pending"
	OrderStatusProcessing OrderStatus = "Processing"
	OrderStatusShipped    OrderStatus = "Shipped"
	OrderStatusDelivered  OrderStatus = "Delivered"
	OrderStatusExpired    OrderStatus = "Expired"
)

type (
	OrderStatus string

	Order struct {
		ID        int32
		UserID    int32
		PaymentID string
		Status    OrderStatus
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddOrder struct {
		UserID    int32
		PaymentID string
	}

	OrderItem struct {
		ID        int32
		OrderID   int32
		ProductID int32
		Quantity  int32
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddOrderItem struct {
		OrderID   int32
		ProductID int32
		Quantity  int32
	}
)
