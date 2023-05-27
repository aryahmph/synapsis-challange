package payment

type (
	PaymentCallback struct {
		ExternalID string `json:"external_id" validate:"required"`
	}
)
