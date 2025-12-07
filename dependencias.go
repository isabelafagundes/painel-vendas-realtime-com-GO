package main

import (
	"painel-vendas-realtime/internal/controller"

	"github.com/labstack/echo/v4"
)

var (
	healthController *controller.HealthController
)

func InicializarDependencias(contexto *echo.Group) {
	instanciarControllers()
	inicializarControllers(contexto)
}

func instanciarControllers() {
	healthController = &controller.HealthController{}
}

func inicializarControllers(contexto *echo.Group) {
	healthController.InicializarHealthController(contexto)
}
