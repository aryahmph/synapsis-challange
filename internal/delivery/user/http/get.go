package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
	httpCartItemCommon "synapsis-challange/common/http/cart-item"
	httpProductCommon "synapsis-challange/common/http/product"
)

func (d HTTPUserDelivery) getCart(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get(httpCommon.AUTH_USER).(int32)

	items, err := d.cartItemUCase.ListCartItemsByUserID(ctx, userId)
	if err != nil {
		return err
	}

	data := []httpCartItemCommon.FullCartItem{}
	for _, item := range items {
		data = append(data, httpCartItemCommon.FullCartItem{
			ID:        item.ID,
			Product:   httpProductCommon.Product(item.Product),
			Quantity:  item.Quantity,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: data,
	})
}
