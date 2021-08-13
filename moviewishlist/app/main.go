package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Alamat string `json:"alamat"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type BaseResponse struct {
	Code    int         `json:"code"`
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserOneResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    User   `json:"data"`
}

func main() {
	//fmt.Print("Hello World!")
	e := echo.New()
	e.GET("/users", GetUserController)
	e.POST("/users/login", LoginUserController)
	e.GET("/users/:userID", DetailUserController)
	e.Start(":8000")
}

func LoginUserController(c echo.Context) error {
	//email := c.FormValue("email")
	//password := c.FormValue("password")
	var userLogin UserLogin
	c.Bind(&userLogin)

	var response = BaseResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Data:    userLogin,
	}
	return c.JSON(http.StatusOK, response)
}

func DetailUserController(c echo.Context) error {
	userID, _ := strconv.Atoi(c.Param("userID"))

	var user = User{userID, "Budi", "budi@email.com", "Jakarta"}
	var response = UserOneResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}

func GetUserController(c echo.Context) error {

	alamat := c.QueryParam("alamat")

	var user = []User{
		User{1, "Budi", "budi@email.com", alamat},
		User{2, "Budi2", "budi2@email.com", "Jakarta2"},
		User{3, "Budi3", "budi2@email.com", "Jakarta3"},
	}
	var response = BaseResponse{
		Code:    http.StatusOK,
		Status:  true,
		Message: "Success",
		Data:    user,
	}
	return c.JSON(http.StatusOK, response)
}
