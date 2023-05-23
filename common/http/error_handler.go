package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	errorCommon "synapsis-challange/common/error"
)

func (h HTTPServer) errorHandler(err error, c echo.Context) {
	if he, ok := err.(*errorCommon.ClientError); ok {
		err = &echo.HTTPError{
			Code: he.StatusCode,
			Message: map[string]interface{}{
				"message": he.Message,
			},
		}
	} else if he, ok = errorCommon.DomainErrorTranslatorDirectories[err]; ok {
		err = &echo.HTTPError{
			Code: he.StatusCode,
			Message: map[string]interface{}{
				"message": he.Message,
			},
		}
	} else if _, ok := err.(*echo.HTTPError); !ok {
		err = &echo.HTTPError{
			Code:    http.StatusInternalServerError,
			Message: http.StatusInternalServerError,
		}
	}

	h.E.DefaultHTTPErrorHandler(err, c)
}
