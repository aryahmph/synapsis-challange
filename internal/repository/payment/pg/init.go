package pg

import (
	"context"
	"synapsis-challange/common/sqlc"
	paymentModel "synapsis-challange/internal/model/payment"
	paymentRepo "synapsis-challange/internal/repository/payment"
)

type pgPaymentRepository struct {
	querier sqlc.Querier
}

func NewPGPaymentRepository(querier sqlc.Querier) paymentRepo.Repository {
	return &pgPaymentRepository{querier: querier}
}

func (r pgPaymentRepository) GetPayment(ctx context.Context, id string) (paymentModel.Payment, error) {
	payment, err := r.querier.GetPayment(ctx, id)
	if err != nil {
		return paymentModel.Payment{}, err
	}
	return paymentModel.Payment{
		ID:        payment.ID,
		UserID:    payment.UserID,
		VaNumber:  payment.VaNumber,
		Amount:    payment.Amount,
		Status:    paymentModel.PaymentStatus(payment.Status),
		CreatedAt: payment.CreatedAt,
		UpdatedAt: payment.UpdatedAt,
	}, nil
}
