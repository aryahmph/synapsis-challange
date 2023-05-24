package product

import (
	productRepo "synapsis-challange/internal/repository/product"
)

type productUsecase struct {
	productRepository productRepo.Repository
}

func NewProductUsecase(productRepository productRepo.Repository) *productUsecase {
	return &productUsecase{
		productRepository: productRepository,
	}
}
