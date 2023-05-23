package password

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHashManager struct {
}

func New() *PasswordHashManager {
	return &PasswordHashManager{}
}

func (p PasswordHashManager) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (p PasswordHashManager) CheckPasswordHash(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
