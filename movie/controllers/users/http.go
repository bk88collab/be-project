package users

import (
	"movie/businesses/users"
	"net/http"

	"movie/controllers/users/request"

	"github.com/labstack/echo/v4"

	controller "movie/controllers"
)

type UserController struct {
	userUseCase users.UseCase
}

func NewUserController(uc users.UseCase) *UserController {
	return &UserController{
		userUseCase: uc,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	request := request.Users{}
	if err := c.Bind(&request); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.userUseCase.Register(ctx, request.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Success Submit Data")
}

func (ctrl *UserController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	request := request.Users{}
	if err := c.Bind(&request); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.userUseCase.Update(ctx, request.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Success Update Data")
}
