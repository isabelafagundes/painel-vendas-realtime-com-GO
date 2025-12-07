package model

type ResumoLoja struct {
	NumeroLoja  string  `json:"numeroLoja"`
	NomeLoja    string  `json:"nomeLoja"`
	TotalVendas int64   `json:"totalVendas"`
	ValorTotal  float64 `json:"valorTotal"`
}

type EstadoPainelVendas struct {
	QuantidadeTotalVendas int64        `json:"quantidadeTotalVendas"`
	ValorTotalVenda       float64      `json:"valorTotalVendas"`
	ResumoPorLoja         []ResumoLoja `json:"resumoPorLoja"`
	UltimasVendas         []Venda      `json:"ultimasVendas"`
}
