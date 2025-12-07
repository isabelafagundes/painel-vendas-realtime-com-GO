package main

import (
	"painel-vendas-realtime/internal/controller"
	"painel-vendas-realtime/internal/service"

	"github.com/labstack/echo/v4"
)

var (
	//controllers
	healthController *controller.HealthController

	//service
	simuladorVendasService *service.SimuladorVendasService
)

func InicializarDependencias(contexto *echo.Group) {
	instanciarServices()
	inicializarServices()

	instanciarControllers()
	inicializarControllers(contexto)
}

func instanciarServices() {
	simuladorVendasService = &service.SimuladorVendasService{}
}

func inicializarServices() {
	simuladorVendasService.InicializarSimuladorVendasService()
}

func instanciarControllers() {
	healthController = &controller.HealthController{}
}

func inicializarControllers(contexto *echo.Group) {
	healthController.InicializarHealthController(contexto)
}
