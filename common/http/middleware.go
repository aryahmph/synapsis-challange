package http

import (
	"github.com/labstack/echo/v4"
	errorCommon "synapsis-challange/common/error"
	jwtCommon "synapsis-challange/common/jwt"
)

func Auth(jwtManager *jwtCommon.JWTManager) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			a := c.Request().Header.Get("Authorization")
			if len(a) <= BEARER {
				return errorCommon.NewInvariantError("Authorization header not valid")
			}
			idToken := a[BEARER:]

			// verify token is real from user
			uid, err := jwtManager.Verify(idToken)
			if err != nil {
				return errorCommon.NewUnauthorizedError("Token not valid")
			}

			c.Set(AUTH_USER, uid)
			return next(c)
		}
	}
}
