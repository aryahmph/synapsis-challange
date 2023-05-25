package cart_item

import (
	"context"
	"errors"
)

var (
	ErrDeleteCartItem_UserNotAuthorized error = errors.New("DELETE_CART_ITEM.USER_NOT_AUTHORIZED")
)

func (u cartItemUsecase) DeleteCartItem(ctx context.Context, id, userId int32) error {
	item, err := u.cartItemRepository.GetCartItem(ctx, id)
	if err != nil {
		return err
	}

	if item.UserID != userId {
		return ErrDeleteCartItem_UserNotAuthorized
	}
	return u.cartItemRepository.DeleteCartItem(ctx, id)
}
