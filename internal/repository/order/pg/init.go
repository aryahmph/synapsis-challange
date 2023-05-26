package pg

import (
	"context"
	store "synapsis-challange/common/pg"
	"synapsis-challange/common/sqlc"
	cartItemModel "synapsis-challange/internal/model/cart-item"
	orderModel "synapsis-challange/internal/model/order"
	paymentModel "synapsis-challange/internal/model/payment"
)

type pgOrderRepository struct {
	querier *store.Store
}

func NewPGOrderRepository(querier *store.Store) *pgOrderRepository {
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
