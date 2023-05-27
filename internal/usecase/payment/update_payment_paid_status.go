package payment

import (
	"context"
	orderModel "synapsis-challange/internal/model/order"
	paymentModel "synapsis-challange/internal/model/payment"
)

func (u paymentUsecase) UpdatePaymentPaidStatus(ctx context.Context, id string) error {
	err := u.paymentRepository.UpdatePaymentStatus(ctx, paymentModel.UpdatePaymentStatus{
		ID:     id,
		Status: paymentModel.PaymentStatusCompleted,
	})
	if err != nil {
		return err
	}

	order, err := u.orderRepository.GetOrderByPaymentID(ctx, id)
	if err != nil {
		return err
	}

	return u.orderRepository.UpdateOrderStatus(ctx, orderModel.UpdateOrderStatus{
		ID:     order.ID,
		Status: orderModel.OrderStatusProcessing,
	})
}
