package order

import (
	"context"
	orderModel "synapsis-challange/internal/model/order"
)

type Usecase interface {
	CreateOrder(ctx context.Context, userId int32) (int32, string, float64, error)
	GetOrder(ctx context.Context, id, userId int32) (orderModel.FullOrder, error)
}
