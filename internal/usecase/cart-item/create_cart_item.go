package cart_item

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

func (u cartItemUsecase) CreateCartItem(ctx context.Context, arg cartItemModel.AddCartItem) (int32, error) {
	rid, err := u.cartItemRepository.VerifyAvailableCartItem(ctx, arg.ProductID, arg.UserID)
	if err != nil {
		return 0, err
	}

	if rid != 0 {
		err = u.cartItemRepository.UpdateCartItemQuantity(ctx, rid, arg.Quantity)
		return rid, err
	}
	return u.cartItemRepository.CreateCartItem(ctx, arg)
}
