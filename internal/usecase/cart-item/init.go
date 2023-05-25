package cart_item

import (
	cartItemRepo "synapsis-challange/internal/repository/cart-item"
	productRepo "synapsis-challange/internal/repository/product"
)

type cartItemUsecase struct {
	cartItemRepository cartItemRepo.Repository
	productRepository  productRepo.Repository
}

func NewCartItemUsecase(
	cartItemRepository cartItemRepo.Repository,
	productRepository productRepo.Repository,
) *cartItemUsecase {
	return &cartItemUsecase{
		cartItemRepository: cartItemRepository,
		productRepository:  productRepository,
	}
}
