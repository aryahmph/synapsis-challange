package cart_item

import "time"

type (
	CartItem struct {
		ID        int32
		ProductID int32
		UserID    int32
		Quantity  int32
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddCartItem struct {
		ProductID int32
		UserID    int32
		Quantity  int32
	}
)
