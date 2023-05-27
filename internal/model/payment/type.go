package payment

import "time"

const (
	PaymentStatusPending   PaymentStatus = "Pending"
	PaymentStatusCompleted PaymentStatus = "Completed"
)

type (
	PaymentStatus string

	Payment struct {
		ID        string
		UserID    int32
		VaNumber  string
		Amount    float64
		Status    PaymentStatus
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	AddPayment struct {
		ID       string
		UserID   int32
		VaNumber string
		Amount   float64
	}

	UpdatePaymentStatus struct {
		ID     string
		Status PaymentStatus
	}
)
