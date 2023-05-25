package cart_item

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

func (u cartItemUsecase) ListCartItemsByUserID(ctx context.Context, userId int32) ([]cartItemModel.FullCartItem, error) {
	fullCartItems := []cartItemModel.FullCartItem{}
	cartItems, err := u.cartItemRepository.ListCartItemsByUserID(ctx, userId)
	if err != nil {
		return nil, err
	}

	for _, item := range cartItems {
		product, err := u.productRepository.GetProduct(ctx, item.ProductID)
		if err != nil {
			return nil, err
		}
		fullCartItems = append(fullCartItems, cartItemModel.FullCartItem{
			ID:        item.ID,
			Product:   product,
			UserID:    item.UserID,
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return fullCartItems, nil
}
