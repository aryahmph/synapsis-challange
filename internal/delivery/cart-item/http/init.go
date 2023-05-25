package http

import (
	"github.com/labstack/echo/v4"
	httpCommon "synapsis-challange/common/http"
	jwtCommon "synapsis-challange/common/jwt"
	cartItemUsecase "synapsis-challange/internal/usecase/cart-item"
)

type HTTPCartItemDelivery struct {
	cartItemUCase cartItemUsecase.Usecase
}

func NewHTTPCartItemDelivery(g *echo.Group, cartItemUCase cartItemUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPCartItemDelivery {
	h := HTTPCartItemDelivery{cartItemUCase: cartItemUCase}
	g.POST("/cart-items", h.addCartItem, httpCommon.Auth(jwtManager))
	g.DELETE("/cart-items/:id", h.deleteCartItem, httpCommon.Auth(jwtManager))

	return h
}
