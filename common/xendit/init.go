package xendit

import (
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/virtualaccount"
	"time"
)

type XenditManager struct {
	secretKey string
}

func NewXenditManager(secretKey string) *XenditManager {
	xendit.Opt.SecretKey = secretKey
	return &XenditManager{secretKey: secretKey}
}

func (*XenditManager) NewVirtualAccountPayment(id string, amount float64) (string, error) {
	trueVal := true
	expire := time.Now().Add(24 * time.Hour)
	data := virtualaccount.CreateFixedVAParams{
		ExternalID:     id,
		BankCode:       "BRI",
		Name:           "Synapsis.id Chllange",
		ExpectedAmount: amount,
		IsSingleUse:    &trueVal,
		IsClosed:       &trueVal,
		ExpirationDate: &expire,
	}

	resp, err := virtualaccount.CreateFixedVA(&data)
	if err != nil {
		return "", err
	}
	return resp.AccountNumber, nil
}
