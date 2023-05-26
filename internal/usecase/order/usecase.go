package order

import (
	"context"
)

type Usecase interface {
	CreateOrder(ctx context.Context, userId int32) (int32, string, int32, error)
}
