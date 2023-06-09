package payment

import (
	"context"
	paymentModel "synapsis-challange/internal/model/payment"
)

type Repository interface {
	GetPayment(ctx context.Context, id string) (paymentModel.Payment, error)
	UpdatePaymentStatus(ctx context.Context, arg paymentModel.UpdatePaymentStatus) error
}
