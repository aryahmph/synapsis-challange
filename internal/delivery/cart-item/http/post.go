package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
	httpCartItemCommon "synapsis-challange/common/http/cart-item"
	cartItemModel "synapsis-challange/internal/model/cart-item"
)

func (d HTTPCartItemDelivery) addCartItem(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get(httpCommon.AUTH_USER).(int32)

	cartItem := &httpCartItemCommon.AddCartItem{}
	if err := c.Bind(cartItem); err != nil {
		return err
	}

	if err := c.Validate(cartItem); err != nil {
		return err
	}

	rid, err := d.cartItemUCase.CreateCartItem(ctx, cartItemModel.AddCartItem{
		ProductID: cartItem.ProductID,
		UserID:    userId,
		Quantity:  cartItem.Quantity,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, httpCommon.Response{
		Data: rid,
	})
}
