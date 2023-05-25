package product

import (
	"context"
	productModel "synapsis-challange/internal/model/product"
)

func (u productUsecase) ListProductsByCategory(ctx context.Context, category string) ([]productModel.Product, error) {
	return u.productRepository.ListProductByCategory(ctx, category)
}
