package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManager struct {
	AccessTokenKey []byte
}

func New(accessTokenKey string) *JWTManager {
	return &JWTManager{AccessTokenKey: []byte(accessTokenKey)}
}

func (j JWTManager) Verify(tokenString string) (id int32, err error) {
	claims := &CustomClaims{}
	if _, err = jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		return j.AccessTokenKey, nil
	}); err != nil {
		return 0, err
	}
	return claims.ID, err
}

func (j JWTManager) Generate(id int32) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
		},
	})
	return token.SignedString(j.AccessTokenKey)
}
