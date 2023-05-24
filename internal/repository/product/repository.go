package product

import (
	"context"
	productModel "synapsis-challange/internal/model/product"
)

type Repository interface {
	ListProducts(ctx context.Context) ([]productModel.Product, error)
	ListProductByCategory(ctx context.Context, category string) ([]productModel.Product, error)
	GetProduct(ctx context.Context, id int32) (productModel.Product, error)
}
