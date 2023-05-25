package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	httpCommon "synapsis-challange/common/http"
)

func (d HTTPCartItemDelivery) deleteCartItem(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get(httpCommon.AUTH_USER).(int32)

	id := c.Param("id")
	idn, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	err = d.cartItemUCase.DeleteCartItem(ctx, int32(idn), userId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: idn,
	})
}
