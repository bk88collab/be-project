package linktrailers

import (
	linktrailers "movie/businesses/linktrailers"
	controller "movie/controllers"
	"movie/controllers/linktrailers/request"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type LinkTrailerController struct {
	linkTrailerUseCase linktrailers.UseCase
}

func NewLinkTrailerController(luc linktrailers.UseCase) *LinkTrailerController {
	return &LinkTrailerController{
		linkTrailerUseCase: luc,
	}
}

func (ctrl *LinkTrailerController) CreateLink(c echo.Context) error {
	ctx := c.Request().Context()
	request := request.LinkTrailer{}
	if err := c.Bind(&request); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.linkTrailerUseCase.Create(ctx, request.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, request.ToDomain())
}

func (ctrl *LinkTrailerController) UpdateLink(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.LinkTrailer{}
	paramId := c.Param("id_trailer")
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	err = ctrl.linkTrailerUseCase.Update(ctx, id, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, req.ToDomain())
}

func (ctrl *LinkTrailerController) DeleteLink(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.LinkTrailer{}
	paramId := c.Param("id_trailer")
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	id, err := strconv.Atoi(paramId)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	err = ctrl.linkTrailerUseCase.Delete(ctx, id, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, req.ToDomain())
}

func (ctrl *LinkTrailerController) GetLink(c echo.Context) error {
	ctx := c.Request().Context()
	request := request.LinkTrailer{}
	if err := c.Bind(&request); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err := ctrl.linkTrailerUseCase.GetUrl(ctx, request.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, request.ToDomain())
}
