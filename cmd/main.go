package main

import (
	"context"
	"os"
	"painel-vendas-realtime/internal/config"
	"time"

	"github.com/labstack/echo/v4"
)

const (
	PORTA         = "PORTA"
	TIPO_AMBIENTE = "TIPO_AMBIENTE"
	INTERVALO     = 2 * time.Second
)

func main() {
	echo := echo.New()
	grupo := echo.Group("/api")

	InicializarDependencias(grupo)
	configuracao := carregarConfiguracao()

	contexto, cancel := context.WithCancel(context.Background())
	defer cancel()

	go simuladorVendasService.Iniciar(contexto, INTERVALO)

	echo.Start(":" + configuracao.Porta)

}

func carregarConfiguracao() *config.Configuracao {
	porta := os.Getenv(PORTA)
	tipoAmbiente := os.Getenv(TIPO_AMBIENTE)

	configuracao := config.CriarConfiguracao()

	if porta != "" {
		configuracao.Porta = porta
	}

	if tipoAmbiente != "" {
		tipo, existe := config.StringParaTipoAmbiente(tipoAmbiente)
		if existe {
			configuracao.TipoAmbiente = tipo
		}
	}

	return configuracao
}
