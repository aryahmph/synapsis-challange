package http

import (
	"github.com/labstack/echo/v4"
	productUsecase "synapsis-challange/internal/usecase/product"
)

type HTTPProductDelivery struct {
	productUCase productUsecase.Usecase
}

func NewHTTPProductDelivery(g *echo.Group, productUCase productUsecase.Usecase) HTTPProductDelivery {
	h := HTTPProductDelivery{productUCase: productUCase}
	g.GET("/products", h.listProducts)

	return h
}
