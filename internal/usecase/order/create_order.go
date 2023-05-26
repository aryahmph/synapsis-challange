package order

import (
	"context"
	orderModel "synapsis-challange/internal/model/order"
	paymentModel "synapsis-challange/internal/model/payment"
)

func (u orderUsecase) CreateOrder(ctx context.Context, userId int32) (int32, string, float64, error) {
	items, err := u.cartItemRepository.ListCartItemsByUserID(ctx, userId)
	if err != nil {
		return 0, "", 0, err
	}

	amount := float64(0)
	for _, item := range items {
		product, err := u.productRepository.GetProduct(ctx, item.ProductID)
		if err != nil {
			return 0, "", 0, err
		}
		amount += product.Price * float64(item.Quantity)
	}

	// Create payment
	pid, err := u.uuidGenerator.Generate()
	if err != nil {
		return 0, "", 0, err
	}
	vaNumber, err := u.xenditManager.NewVirtualAccountPayment(pid, amount)
	if err != nil {
		return 0, "", 0, err
	}

	order := orderModel.AddOrder{
		UserID:    userId,
		PaymentID: pid,
	}
	payment := paymentModel.AddPayment{
		ID:       pid,
		UserID:   userId,
		VaNumber: vaNumber,
		Amount:   amount,
	}

	rid, err := u.orderRepository.CreateOrder(ctx, order, payment, items)
	if err != nil {
		return 0, "", 0, err
	}
	return rid, vaNumber, amount, nil
}
