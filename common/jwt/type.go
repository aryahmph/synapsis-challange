package jwt

import "github.com/golang-jwt/jwt/v5"

type (
	CustomClaims struct {
		ID int32 `json:"id"`
		jwt.RegisteredClaims
	}
)
