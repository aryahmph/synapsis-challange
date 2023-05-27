package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
	httpPaymentCommon "synapsis-challange/common/http/payment"
)

func (d HTTPPaymentDelivery) paymentPaidCallback(c echo.Context) error {
	ctx := c.Request().Context()

	payload := &httpPaymentCommon.PaymentCallback{}
	if err := c.Bind(payload); err != nil {
		return err
	}

	if err := c.Validate(payload); err != nil {
		return err
	}

	err := d.paymentUCase.UpdatePaymentPaidStatus(ctx, payload.ExternalID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: payload.ExternalID,
	})
}
