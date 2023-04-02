package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/Stanxxy/stan-go-web/internal/context"
	"github.com/Stanxxy/stan-go-web/internal/core/errors"
	"github.com/Stanxxy/stan-go-web/internal/models"
)

type (
	User          struct{}
	UserViewModel struct {
		Name string
		ID   string
	}
)

func (ctrl User) AddUser(c echo.Context) error {

	cc := c.(*context.AppContext)

	var user models.User

	// user := models.User{}

	if err := c.BindJSON(&user); err != nil {
        return err
    }

	if err != nil {
		b := errors.NewBoom(errors.InvalidBindingModel, errors.ErrorText(errors.InvalidBindingModel), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	err := cc.UserStore.Create(&user)

	if err != nil {
		b := errors.NewBoom(errors.EntityCreationError, errors.ErrorText(errors.EntityCreationError), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	// userID := c.Param("id")

	// vm := UserViewModel{
	// 	Name: user.Name,
	// 	ID:   user.ID,
	// }

	// Do something with the user object
    return c.String(http.StatusOK, "User created")
	// return c.JSON(http.StatusOK, vm)

}

func (ctrl UserList) GetUsers(c echo.Context) error {
	cc := c.(*context.AppContext)

	users := []models.User{}

	err := cc.UserStore.Find(&users)

	if err != nil {
		b := errors.NewBoom(errors.UserNotFound, errors.ErrorText(errors.UserNotFound), err)
		c.Logger().Error(err)
		return c.JSON(http.StatusInternalServerError, b)
	}

	vm := UserListViewModel{
		Users: make([]UserViewModel, len(users)),
	}

	for index, user := range users {
		vm.Users[index] = UserViewModel{
			Name: user.Name,
			ID:   user.ID,
		}
	}

	return c.JSON(http.StatusOK, vm)
}

func (ctrl User) GetUser(c echo.Context) error {
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

	return c.JSON(http.StatusOK, vm)
}
