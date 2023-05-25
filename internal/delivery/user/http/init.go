package http

import (
	"github.com/labstack/echo/v4"
	httpCommon "synapsis-challange/common/http"
	jwtCommon "synapsis-challange/common/jwt"
	cartItemUsecase "synapsis-challange/internal/usecase/cart-item"
	userUsecase "synapsis-challange/internal/usecase/user"
)

type HTTPUserDelivery struct {
	userUCase     userUsecase.Usecase
	cartItemUCase cartItemUsecase.Usecase
}

func NewHTTPUserDelivery(g *echo.Group, userUCase userUsecase.Usecase,
	cartItemUCase cartItemUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPUserDelivery {
	h := HTTPUserDelivery{userUCase: userUCase, cartItemUCase: cartItemUCase}
	g.POST("/auth", h.loginUser)

	g.POST("/users", h.addUser)
	g.GET("/users/carts", h.getCart, httpCommon.Auth(jwtManager))

	return h
}
