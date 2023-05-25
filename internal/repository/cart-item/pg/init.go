package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	errorCommon "synapsis-challange/common/error"
	"synapsis-challange/common/sqlc"
	cartItemModel "synapsis-challange/internal/model/cart-item"
	cartItemRepo "synapsis-challange/internal/repository/cart-item"
)

type pgCartItemRepository struct {
	querier sqlc.Querier
}

func NewPGPCartItemRepository(querier sqlc.Querier) cartItemRepo.Repository {
	return &pgCartItemRepository{querier: querier}
}

func (r pgCartItemRepository) CreateCartItem(ctx context.Context, arg cartItemModel.AddCartItem) (int32, error) {
	return r.querier.CreateCartItem(ctx, sqlc.CreateCartItemParams(arg))
}

func (r pgCartItemRepository) GetCartItem(ctx context.Context, id int32) (cartItemModel.CartItem, error) {
	cartItem, err := r.querier.GetCartItem(ctx, id)
	if err == pgx.ErrNoRows {
		return cartItemModel.CartItem{}, errorCommon.NewNotFoundError("CartItem not found")
	}
	return cartItemModel.CartItem(cartItem), err
}

func (r pgCartItemRepository) VerifyAvailableCartItem(ctx context.Context, productId, userId int32) (int32, error) {
	rid, err := r.querier.VerifyAvailableCartItem(ctx, sqlc.VerifyAvailableCartItemParams{
		ProductID: productId,
		UserID:    userId,
	})
	if err == pgx.ErrNoRows {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	return rid, nil
}

func (r pgCartItemRepository) UpdateCartItemQuantity(ctx context.Context, id, quantity int32) error {
	return r.querier.UpdateCartItemQuantity(ctx, sqlc.UpdateCartItemQuantityParams{
		ID:       id,
		Quantity: quantity,
	})
}

func (r pgCartItemRepository) ListCartItemsByUserID(ctx context.Context, cartId int32) ([]cartItemModel.CartItem, error) {
	list, err := r.querier.ListCartItemsByUserID(ctx, cartId)
	if err != nil {
		return nil, err
	}

	cartItems := make([]cartItemModel.CartItem, 0)
	for _, item := range list {
		cartItems = append(cartItems, cartItemModel.CartItem(item))
	}
	return cartItems, nil
}

func (r pgCartItemRepository) DeleteCartItem(ctx context.Context, id int32) error {
	return r.querier.DeleteCartItem(ctx, id)
}
