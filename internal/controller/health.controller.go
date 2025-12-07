package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct {
}

func InstanciarHealthController() *HealthController {
	controller := &HealthController{}
	return controller
}

func (controller *HealthController) InicializarHealthController(contexto *echo.Group) {
	grupo := contexto.Group("/health")
	grupo.GET("", controller.verificarServidor)
}

func (controller *HealthController) verificarServidor(contexto echo.Context) error {
	mapaOk := make(map[string]string)
	mapaOk["status"] = "ok"
	return contexto.JSON(http.StatusOK, mapaOk)
}
