package handler_health

import (
	"health-app/app/middlewares"
	"health-app/controllers/health/request"
	domain_health "health-app/domain/health"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HandlerHealth struct {
	BusinessHealth domain_health.Business
	
}

func NewHandlerHealth(bh domain_health.Business) HandlerHealth {
	return HandlerHealth{
		BusinessHealth: bh,
	}
}

func (hh *HandlerHealth) AddDataHealth(ctx echo.Context) error {
	req := request.JSONHealth{}
	ctx.Bind(&req)
	claim := middlewares.GetUser(ctx)
	err := hh.BusinessHealth.AddDataHealth(claim.UserID, request.ToDomain(req))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusCreated,
	})
}

func (hh *HandlerHealth) CalculateIdealWeight(ctx echo.Context) error {
	claim := middlewares.GetUser(ctx)
	data, err := hh.BusinessHealth.CalculateIdealWeight(claim.UserID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": err.Error(),
			"rescode": http.StatusBadRequest,
		})
	}
	return ctx.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"rescode": http.StatusOK,
		"result": data,
	})
}