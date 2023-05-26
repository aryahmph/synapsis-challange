package uuid

import (
	"github.com/google/uuid"
)

type UUIDGenerator struct {
}

func New() *UUIDGenerator {
	return &UUIDGenerator{}
}

func (*UUIDGenerator) Generate() (string, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return uid.String(), nil
}
