package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
	httpProductCommon "synapsis-challange/common/http/product"
	productModel "synapsis-challange/internal/model/product"
)

func (d HTTPProductDelivery) listProducts(c echo.Context) error {
	ctx := c.Request().Context()

	category := c.QueryParam("category")
	var products []productModel.Product
	var err error
	if category != "" {
		products, err = d.productUCase.ListProductsByCategory(ctx, category)
		if err != nil {
			return err
		}
	} else {
		products, err = d.productUCase.ListProducts(ctx)
		if err != nil {
			return err
		}
	}

	data := []httpProductCommon.Product{}
	for _, product := range products {
		data = append(data, httpProductCommon.Product{
			ID:          product.ID,
			Name:        product.Name,
			Price:       product.Price,
			Description: product.Description,
			Category:    product.Category,
			CreatedAt:   product.CreatedAt,
			UpdatedAt:   product.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: data,
	})
}
