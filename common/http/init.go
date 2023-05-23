package http

import (
	"errors"
	"net/http"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type HTTPServer struct {
	E *echo.Echo
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.validator.Struct(i); err != nil {
		var ves validator.ValidationErrors
		errors.As(err, &ves)
		keys := make(map[string]string)
		for _, ve := range ves {
			keys[ve.Field()] = ve.Tag()
		}

		return &echo.HTTPError{
			Code: http.StatusBadRequest,
			Message: map[string]interface{}{
				"message": ves.Error(),
				"errors":  keys,
			},
		}
	}
	return nil
}

func NewHTTPServer() HTTPServer {
	v := validator.New()
	v.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return fld.Name
		}
		return name
	})

	e := echo.New()
	e.Validator = &CustomValidator{validator: v}
	h := HTTPServer{E: e}

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.Recover())

	e.HTTPErrorHandler = h.errorHandler

	return h
}
