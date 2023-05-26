package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	httpCommon "synapsis-challange/common/http"
	httpOrderCommon "synapsis-challange/common/http/order"
)

func (d HTTPOrderDelivery) getOrder(c echo.Context) error {
	ctx := c.Request().Context()
	userId := c.Get(httpCommon.AUTH_USER).(int32)

	id := c.Param("id")
	idn, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return err
	}

	order, err := d.orderUCase.GetOrder(ctx, int32(idn), userId)
	if err != nil {
		return err
	}

	orderItems := make([]httpOrderCommon.OrderItem, 0)
	for _, item := range order.OrderItems {
		orderItems = append(orderItems, httpOrderCommon.OrderItem{
			ID:        item.ID,
			ProductID: item.Product.ID,
			Quantity:  item.Quantity,
			Name:      item.Product.Name,
			Price:     item.Product.Price,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	response := httpOrderCommon.Order{
		ID: order.ID,
		Payment: httpOrderCommon.Payment{
			ID:        order.Payment.ID,
			VaNumber:  order.Payment.VaNumber,
			Amount:    order.Payment.Amount,
			Status:    string(order.Payment.Status),
			CreatedAt: order.CreatedAt,
			UpdatedAt: order.UpdatedAt,
		},
		OrderItems: orderItems,
		Status:     string(order.Status),
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}

	return c.JSON(http.StatusOK, httpCommon.Response{
		Data: response,
	})
}
