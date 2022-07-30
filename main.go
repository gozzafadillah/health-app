package main

import (
	"health-app/app/config"
	middlewares "health-app/app/middlewares"
	service_users "health-app/business/users"
	handler_users "health-app/controllers/users"
	migrate "health-app/migrator"
	mysql_users "health-app/repository/users/mysql"
	routes "health-app/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	db := config.InitDB()
	migrate.AutoMigrate(db)

	configJWT := middlewares.ConfigJwt{
		SecretJWT: config.Conf.JWTSecret,
	}

	e := echo.New()

	//Factory
	// Users
	userRepo := mysql_users.NewUsersRepo(db)
	userServ := service_users.NewUsersBusiness(userRepo, &configJWT)
	UserHandler := handler_users.NewUsersHandler(userServ)
	
	// Route
	routeInit := routes.ControllerList{
		JWTMiddleware:      configJWT.Init(),
		UserHandler:        UserHandler,
	}

	routeInit.RouteRegister(e)
	// start server
	e.Logger.Fatal(e.Start(":8080"))

}
