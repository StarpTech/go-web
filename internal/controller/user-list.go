package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/starptech/go-web/internal/context"
	"github.com/starptech/go-web/internal/core/errors"
	"github.com/starptech/go-web/internal/models"
)

type (
	UserList          struct{}
	UserListViewModel struct {
		Users []UserViewModel
	}
)

func (ctrl UserList) GetUsers(c echo.Context) error {
	cc := c.(*context.AppContext)

	users := []models.User{}

	err := cc.UserStore.Find(&users)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	viewModel := UserListViewModel{
		Users: make([]UserViewModel, len(users)),
	}

	for index, user := range users {
		viewModel.Users[index] = UserViewModel{
			Name: user.Name,
			ID:   user.ID,
		}
	}

	return c.Render(http.StatusOK, "user-list.html", viewModel)

}
