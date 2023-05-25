package cart_item

type (
	AddCartItem struct {
		ProductID int32 `json:"product_id" validate:"required"`
		Quantity  int32 `json:"quantity" validate:"required"`
	}
)
