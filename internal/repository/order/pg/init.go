package pg

import (
	"context"
	store "synapsis-challange/common/pg"
	"synapsis-challange/common/sqlc"
	cartItemModel "synapsis-challange/internal/model/cart-item"
	orderModel "synapsis-challange/internal/model/order"
	paymentModel "synapsis-challange/internal/model/payment"
	orderRepo "synapsis-challange/internal/repository/order"
)

type pgOrderRepository struct {
	querier *store.Store
}

func NewPGOrderRepository(querier *store.Store) orderRepo.Repository {
	return &pgOrderRepository{querier: querier}
}

func (r pgOrderRepository) CreateOrder(
	ctx context.Context, order orderModel.AddOrder,
	payment paymentModel.AddPayment,
	items []cartItemModel.CartItem,
) (int32, error) {
	var oid int32
	err := r.querier.ExecTx(ctx, func(q sqlc.Querier) error {
		// Create payment
		err := q.CreatePayment(ctx, sqlc.CreatePaymentParams(payment))
		if err != nil {
			return err
		}

		// Create Order
		oid, err = q.CreateOrder(ctx, sqlc.CreateOrderParams(order))
		if err != nil {
			return err
		}

		// Move cart item to order item
		for _, item := range items {
			err = q.DeleteCartItem(ctx, item.ID)
			if err != nil {
				return err
			}

			_, err = q.CreateOrderItem(ctx, sqlc.CreateOrderItemParams{
				OrderID:   oid,
				ProductID: item.ProductID,
				Quantity:  item.Quantity,
			})
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return 0, err
	}

	return oid, nil
}

func (r pgOrderRepository) GetOrder(ctx context.Context, id int32) (orderModel.Order, error) {
	order, err := r.querier.Querier.GetOrder(ctx, id)
	if err != nil {
		return orderModel.Order{}, err
	}
	return orderModel.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		PaymentID: order.PaymentID,
		Status:    orderModel.OrderStatus(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}, nil
}

func (r pgOrderRepository) GetOrderByPaymentID(ctx context.Context, paymentId string) (orderModel.Order, error) {
	order, err := r.querier.Querier.GetOrderByPaymentID(ctx, paymentId)
	if err != nil {
		return orderModel.Order{}, err
	}
	return orderModel.Order{
		ID:        order.ID,
		UserID:    order.UserID,
		PaymentID: order.PaymentID,
		Status:    orderModel.OrderStatus(order.Status),
		CreatedAt: order.CreatedAt,
		UpdatedAt: order.UpdatedAt,
	}, nil
}

func (r pgOrderRepository) ListOrderItemsByOrderID(ctx context.Context, orderId int32) ([]orderModel.OrderItem, error) {
	list, err := r.querier.ListOrderItemsByOrderID(ctx, orderId)
	if err != nil {
		return nil, err
	}

	items := make([]orderModel.OrderItem, 0)
	for _, item := range list {
		items = append(items, orderModel.OrderItem(item))
	}
	return items, nil
}

func (r pgOrderRepository) UpdateOrderStatus(ctx context.Context, arg orderModel.UpdateOrderStatus) error {
	return r.querier.UpdateOrderStatus(ctx, sqlc.UpdateOrderStatusParams{
		ID:     arg.ID,
		Status: sqlc.OrderStatus(arg.Status),
	})
}
