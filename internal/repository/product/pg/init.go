package pg

import (
	"context"
	"github.com/jackc/pgx/v4"
	errorCommon "synapsis-challange/common/error"
	"synapsis-challange/common/sqlc"
	productModel "synapsis-challange/internal/model/product"
	productRepo "synapsis-challange/internal/repository/product"
)

type pgProductRepository struct {
	querier sqlc.Querier
}

func NewPGProductRepository(querier sqlc.Querier) productRepo.Repository {
	return &pgProductRepository{querier: querier}
}

func (r pgProductRepository) ListProducts(ctx context.Context) ([]productModel.Product, error) {
	list, err := r.querier.ListProducts(ctx)
	if err != nil {
		return nil, err
	}

	products := make([]productModel.Product, 0)
	for _, product := range list {
		products = append(products, productModel.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Category:    product.Category,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	return products, nil
}

func (r pgProductRepository) ListProductByCategory(ctx context.Context, category string) ([]productModel.Product, error) {
	list, err := r.querier.ListProductsByCategory(ctx, category)
	if err != nil {
		return nil, err
	}

	products := make([]productModel.Product, 0)
	for _, product := range list {
		products = append(products, productModel.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Category:    product.Category,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}
	return products, nil
}

func (r pgProductRepository) GetProduct(ctx context.Context, id int32) (productModel.Product, error) {
	product, err := r.querier.GetProduct(ctx, id)
	if err == pgx.ErrNoRows {
		return productModel.Product{}, errorCommon.NewNotFoundError("Product not found")
	}
	return productModel.Product{
		ID:          product.ID,
		Name:        product.Name,
		Price:       product.Price,
		Description: product.Description,
		Category:    product.Category,
		CreatedAt:   product.CreatedAt,
		UpdatedAt:   product.UpdatedAt,
	}, err
}
