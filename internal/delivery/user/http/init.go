package http

import (
	"github.com/labstack/echo/v4"
	userUsecase "synapsis-challange/internal/usecase/user"
)

type HTTPUserDelivery struct {
	userUCase userUsecase.Usecase
}

func NewHTTPUserDelivery(g *echo.Group, userUCase userUsecase.Usecase) HTTPUserDelivery {
	h := HTTPUserDelivery{userUCase: userUCase}
	g.POST("/users", h.addUser)

	return h
}
