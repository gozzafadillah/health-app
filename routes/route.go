package routes

import (
	handler_users "health-app/controllers/users"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserHandler		   handler_users.UsersHandler
}

// const server = "https://36e2-2001-448a-1102-1a0f-350a-677f-f95c-668a.ap.ngrok.io/"

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// log
	// middlewares.LogMiddleware(e)

	// access public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	e.POST("/login", cl.UserHandler.Login)
	e.POST("/register", cl.UserHandler.Register)
}
