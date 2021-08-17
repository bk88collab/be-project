package main

import (
	"movie/app/routes"
	_UserUseCase "movie/businesses/users"
	_UserController "movie/controllers/users"
	"movie/drivers/mysql"
	_UserRepository "movie/drivers/repository/users"

	echo "github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func dbMigrate(db *gorm.DB) {
	db.AutoMigrate(
		&_UserRepository.Users{},
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

	routeInit := routes.ControllerList{
		UserController: *userCtrl,
	}

	routeInit.RouteRegister(e)
	e.Start(
		":8000",
	)

}
