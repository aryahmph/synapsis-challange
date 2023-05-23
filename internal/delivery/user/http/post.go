package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
	httpCommon "synapsis-challange/common/http"
	userModel "synapsis-challange/internal/model/user"
)

func (d HTTPUserDelivery) addUser(c echo.Context) error {
	ctx := c.Request().Context()

	user := &httpCommon.AddUser{}
	if err := c.Bind(user); err != nil {
		return err
	}

	if err := c.Validate(user); err != nil {
		return err
	}

	nid, err := d.userUCase.CreateUser(ctx, userModel.AddUser{
		Email:        user.Email,
		PasswordHash: user.Password,
		Name:         user.Name,
		Address:      user.Address,
	})
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, httpCommon.Response{
		Data: nid,
	})
}
