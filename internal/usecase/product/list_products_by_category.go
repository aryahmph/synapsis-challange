package product

import (
	"context"
	productModel "synapsis-challange/internal/model/product"
)

func (u productUsecase) ListProducts(ctx context.Context) ([]productModel.Product, error) {
	return u.productRepository.ListProducts(ctx)
}

func (u productUsecase) ListProductsByCategory(ctx context.Context, category string) ([]productModel.Product, error) {
	return u.productRepository.ListProductByCategory(ctx, category)
}
