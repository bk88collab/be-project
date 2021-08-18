package routes

import (
	"movie/controllers/linktrailers"
	"movie/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController    users.UserController
	LinkUrlController linktrailers.LinkTrailerController
}

func (controller *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", controller.UserController.Register)
	users.PUT("/update/:userId", controller.UserController.Update)
	users.DELETE("/delete/:userId", controller.UserController.Delete)
	users.GET("/profile/:userName", controller.UserController.Profile)

	linkurl := e.Group("url")
	linkurl.POST("/create", controller.LinkUrlController.CreateLink)
	linkurl.PUT("/update/:id_trailer", controller.LinkUrlController.UpdateLink)
	linkurl.DELETE("/delete/:id_trailer", controller.LinkUrlController.DeleteLink)
	linkurl.GET("/", controller.LinkUrlController.GetLink)
}
