package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
)

type (
	UserList          struct{}
	UserListViewModel struct {
		Users []UserViewModel
	}
)

func (ctrl UserList) GetUsersAndRender(c echo.Context) error {
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

func (ctrl User) GetUserAndRender(c echo.Context) error {
	cc := c.(*context.AppContext)
	userID := c.Param("id")

	user := models.User{ID: userID}

	err := cc.UserStore.First(&user)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusNotFound, b)
	}

	vm := UserViewModel{
		Name: user.Name,
		ID:   user.ID,
	}

	return c.Render(http.StatusOK, "user.html", vm)

}
