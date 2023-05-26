package order

import (
	"context"
	"errors"
	orderModel "synapsis-challange/internal/model/order"
)

var (
	ErrGetOrder_UserNotAuthorized error = errors.New("GET_ORDER.USER_NOT_AUTHORIZED")
)

func (u orderUsecase) GetOrder(ctx context.Context, id, userId int32) (orderModel.FullOrder, error) {
	order, err := u.orderRepository.GetOrder(ctx, id)
	if err != nil {
		return orderModel.FullOrder{}, err
	}

	if order.ID != userId {
		return orderModel.FullOrder{}, ErrGetOrder_UserNotAuthorized
	}

	payment, err := u.paymentRepository.GetPayment(ctx, order.PaymentID)
	if err != nil {
		return orderModel.FullOrder{}, err
	}

	items, err := u.orderRepository.ListOrderItemsByOrderID(ctx, order.ID)
	if err != nil {
		return orderModel.FullOrder{}, err
	}

	orderItems := make([]orderModel.FullOrderItem, 0)
	for _, item := range items {
		product, err := u.productRepository.GetProduct(ctx, item.ProductID)
		if err != nil {
			return orderModel.FullOrder{}, err
		}
		orderItems = append(orderItems, orderModel.FullOrderItem{
			ID:        item.ID,
			Quantity:  item.Quantity,
			Product:   product,
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		})
	}

	return orderModel.FullOrder{
		ID:         id,
		UserID:     userId,
		Payment:    payment,
		OrderItems: orderItems,
		Status:     order.Status,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}, nil
}
