package payment

import "context"

type Usecase interface {
	UpdatePaymentPaidStatus(ctx context.Context, id string) error
}
