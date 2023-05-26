package order

import (
	"context"
	cartItemModel "synapsis-challange/internal/model/cart-item"
	orderModel "synapsis-challange/internal/model/order"
	paymentModel "synapsis-challange/internal/model/payment"
)

type Repository interface {
	CreateOrder(ctx context.Context, order orderModel.AddOrder, payment paymentModel.AddPayment, items []cartItemModel.CartItem) (int32, error)
}
