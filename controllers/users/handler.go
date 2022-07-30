package handler_users

import (
	"health-app/controllers/users/request"
	domain_users "health-app/domain/users"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UsersHandler struct {
	UsersBusiness domain_users.Business
}
func NewUsersHandler(uc domain_users.Business) UsersHandler {
	return UsersHandler{
		UsersBusiness: uc,
	}
}

func (uh *UsersHandler) Login(ctx echo.Context) error {
	req := request.JSONUsers{}
	ctx.Bind(&req)
	token, err := uh.UsersBusiness.Login(req.Email, req.Password)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"result": token,
	})
}

func (uh *UsersHandler) Register(ctx echo.Context) error {
	req := request.JSONUsers{}
	ctx.Bind(&req)
	err := uh.UsersBusiness.Register(request.ToDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
	}) 
}