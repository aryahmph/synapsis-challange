package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
)

func (d HTTPOrderDelivery) addOrder(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get(httpCommon.AUTH_USER).(int32)

	rid, va, amount, err := d.orderUCase.CreateOrder(ctx, userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, httpCommon.Response{
		Data: map[string]interface{}{
			"id":     rid,
			"va":     va,
			"amount": amount,
		},
	})
}
