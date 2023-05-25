package cart_item

import (
	httpProductCommon "synapsis-challange/common/http/product"
	"time"
)

type (
	AddCartItem struct {
		ProductID int32 `json:"product_id" validate:"required"`
		Quantity  int32 `json:"quantity" validate:"required"`
	}

	FullCartItem struct {
		ID        int32                     `json:"id"`
		Product   httpProductCommon.Product `json:"product"`
		Quantity  int32                     `json:"quantity"`
		CreatedAt time.Time                 `json:"created_at"`
		UpdatedAt time.Time                 `json:"updated_at"`
	}
)
