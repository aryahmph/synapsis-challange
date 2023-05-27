package payment

import (
	orderRepo "synapsis-challange/internal/repository/order"
	paymentRepo "synapsis-challange/internal/repository/payment"
)

type paymentUsecase struct {
	orderRepository   orderRepo.Repository
	paymentRepository paymentRepo.Repository
}

func NewPaymentUsecase(
	orderRepository orderRepo.Repository,
	paymentRepository paymentRepo.Repository,
) *paymentUsecase {
	return &paymentUsecase{
		orderRepository:   orderRepository,
		paymentRepository: paymentRepository,
	}
}
