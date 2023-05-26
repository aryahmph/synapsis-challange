package http

import (
	"github.com/labstack/echo/v4"
	httpCommon "synapsis-challange/common/http"
	jwtCommon "synapsis-challange/common/jwt"
	orderUsecase "synapsis-challange/internal/usecase/order"
)

type HTTPOrderDelivery struct {
	orderUCase orderUsecase.Usecase
}

func NewHTTPOrderDelivery(g *echo.Group, orderUCase orderUsecase.Usecase, jwtManager *jwtCommon.JWTManager) HTTPOrderDelivery {
	h := HTTPOrderDelivery{orderUCase: orderUCase}
	g.POST("/orders", h.addOrder, httpCommon.Auth(jwtManager))

	return h
}
