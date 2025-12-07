package model

type Venda struct {
	Numero      string   `json:"numero"`
	NumeroLoja  string   `json:"numeroLoja"`
	Valor       float64  `json:"valor"`
	DataCriacao DataHora `json:"dataCriacao"`
}
