package product

import (
	"context"
	productModel "synapsis-challange/internal/model/product"
)

func (u productUsecase) ListProducts(ctx context.Context) ([]productModel.Product, error) {
	return u.productRepository.ListProducts(ctx)
}
