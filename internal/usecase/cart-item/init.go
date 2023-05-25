package cart_item

import (
	cartItemRepo "synapsis-challange/internal/repository/cart-item"
)

type cartItemUsecase struct {
	cartItemRepository cartItemRepo.Repository
}

func NewCartItemUsecase(
	cartItemRepository cartItemRepo.Repository,
) *cartItemUsecase {
	return &cartItemUsecase{
		cartItemRepository: cartItemRepository,
	}
}
