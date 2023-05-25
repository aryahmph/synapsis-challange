package cart_item

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

type Usecase interface {
	CreateCartItem(ctx context.Context, arg cartItemModel.AddCartItem) (int32, error)
	ListCartItemsByUserID(ctx context.Context, userId int32) ([]cartItemModel.FullCartItem, error)
}
