package venda

import "painel-vendas-realtime/internal/model"

type Venda struct {
	Numero         string         `json:"numero"`
	NumeroLoja     string         `json:"numeroLoja"`
	NumeroTerminal string         `json:"numeroTerminal"`
	Valor          float64        `json:"valor"`
	DataCriacao    model.DataHora `json:"dataCriacao"`
}
