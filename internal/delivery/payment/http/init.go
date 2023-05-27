package http

import (
	"github.com/labstack/echo/v4"
	paymentUsecase "synapsis-challange/internal/usecase/payment"
)

type HTTPPaymentDelivery struct {
	paymentUCase paymentUsecase.Usecase
}

func NewHTTPPaymentDelivery(g *echo.Group, paymentUCase paymentUsecase.Usecase) HTTPPaymentDelivery {
	h := HTTPPaymentDelivery{paymentUCase: paymentUCase}

	g.POST("/payments/paid", h.paymentPaidCallback)

	return h
}
