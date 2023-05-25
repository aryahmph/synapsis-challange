package cart_item

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

type Usecase interface {
	CreateCartItem(ctx context.Context, arg cartItemModel.AddCartItem) (int32, error)
}
