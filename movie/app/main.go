package main

import (
	"movie/app/routes"
	_UserUseCase "movie/businesses/users"
	_UserController "movie/controllers/users"
	_UserRepository "movie/drivers/repository/users"

	_LinkUrlUseCase "movie/businesses/linktrailers"
	_LinkUrlController "movie/controllers/linktrailers"
	_LinkUrlRepository "movie/drivers/repository/linktrailers"

	"movie/drivers/mysql"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_UserRepository.Users{},
		&_LinkUrlRepository.LinkUrl{},
	)
}

func main() {
	configDB := mysql.ConfigDB{
		DB_Username: "root",
		DB_Password: "",
		DB_Host:     "localhost",
		DB_Port:     "3306",
		DB_Database: "movie_db",
	}
	db := configDB.InitialDB()
	dbMigrate(db)

	e := echo.New()
	userRepo := _UserRepository.NewUserRepository(db)
	userUsecase := _UserUseCase.NewUserUseCase(userRepo)
	userCtrl := _UserController.NewUserController(userUsecase)

	linkUrlRepo := _LinkUrlRepository.NewLinkRepository(db)
	linkUrlUsecase := _LinkUrlUseCase.NewLinkUseCase(linkUrlRepo)
	linkUrlCtrl := _LinkUrlController.NewLinkTrailerController(linkUrlUsecase)

	routeInit := routes.ControllerList{
		UserController:    *userCtrl,
		LinkUrlController: *linkUrlCtrl,
	}

	routeInit.RouteRegister(e)
	e.Start(
		":8000",
	)

}
