package cart_item

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

type Repository interface {
	CreateCartItem(ctx context.Context, arg cartItemModel.AddCartItem) (int32, error)
	GetCartItem(ctx context.Context, id int32) (cartItemModel.CartItem, error)
	VerifyAvailableCartItem(ctx context.Context, productId, userId int32) (int32, error)
	UpdateCartItemQuantity(ctx context.Context, id, quantity int32) error
	ListCartItemsByUserID(ctx context.Context, cartId int32) ([]cartItemModel.CartItem, error)
	DeleteCartItem(ctx context.Context, id int32) error
}
