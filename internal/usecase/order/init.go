package order

import (
	"synapsis-challange/common/uuid"
	xenditCommon "synapsis-challange/common/xendit"
	cartItemRepo "synapsis-challange/internal/repository/cart-item"
	orderRepo "synapsis-challange/internal/repository/order"
	paymentRepo "synapsis-challange/internal/repository/payment"
	productRepo "synapsis-challange/internal/repository/product"
)

type orderUsecase struct {
	orderRepository    orderRepo.Repository
	cartItemRepository cartItemRepo.Repository
	productRepository  productRepo.Repository
	paymentRepository  paymentRepo.Repository

	xenditManager *xenditCommon.XenditManager
	uuidGenerator *uuid.UUIDGenerator
}

func NewOrderUsecase(
	orderRepository orderRepo.Repository,
	cartItemRepository cartItemRepo.Repository,
	productRepository productRepo.Repository,
	paymentRepository paymentRepo.Repository,
	xenditManager *xenditCommon.XenditManager,
	uuidGenerator *uuid.UUIDGenerator,
) *orderUsecase {
	return &orderUsecase{
		orderRepository:    orderRepository,
		cartItemRepository: cartItemRepository,
		productRepository:  productRepository,
		paymentRepository:  paymentRepository,
		xenditManager:      xenditManager,
		uuidGenerator:      uuidGenerator,
	}
}
