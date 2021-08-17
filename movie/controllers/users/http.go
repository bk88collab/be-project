package users

import (
	"movie/businesses/users"
	"movie/controllers/users/request"
	"net/http"
	"strconv"

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

	return controller.NewSuccessResponse(c, request.ToDomain())
}

func (ctrl *UserController) Update(c echo.Context) error {

	ctx := c.Request().Context()
	req := request.Users{}
	paramId := c.Param("userId")
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	err = ctrl.userUseCase.Update(ctx, id, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, req.ToDomain())
}

func (ctrl *UserController) Delete(c echo.Context) error {

	ctx := c.Request().Context()
	req := request.Users{}
	paramId := c.Param("userId")
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	id, err := strconv.Atoi(paramId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	err = ctrl.userUseCase.Delete(ctx, id, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, req.ToDomain())
}

func (ctrl *UserController) Profile(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Users{}
	paramId := c.Param("userName")

	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	userName := paramId
	response, errRes := ctrl.userUseCase.GetProfile(ctx, userName)
	if errRes != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, response)
}
