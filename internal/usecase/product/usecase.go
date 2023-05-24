package product

import (
	"context"
	productModel "synapsis-challange/internal/model/product"
)

type Usecase interface {
	ListProducts(ctx context.Context) ([]productModel.Product, error)
	ListProductsByCategory(ctx context.Context, category string) ([]productModel.Product, error)
}
