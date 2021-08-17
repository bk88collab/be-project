package routes

import (
	"movie/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController users.UserController
}

func (controller *ControllerList) RouteRegister(e *echo.Echo) {
	users := e.Group("users")
	users.POST("/register", controller.UserController.Register)
	users.PUT("/update/:userId", controller.UserController.Update)
	users.DELETE("/delete/:userId", controller.UserController.Delete)
	users.GET("/profile/:userName", controller.UserController.Profile)
}
