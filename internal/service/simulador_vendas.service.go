package service

import (
	"context"
	"math/rand"
	"painel-vendas-realtime/internal/model"
	"painel-vendas-realtime/internal/model/loja"
	"painel-vendas-realtime/internal/model/terminal"
	"painel-vendas-realtime/internal/model/venda"
	"sync"
	"time"
)

type SimuladorVendasService struct {
	vendas        venda.CacheVendas
	lojas         loja.CacheLojas
	terminais     terminal.CacheTerminais
	totalVendas   int64
	valorVendas   float64
	mutex         sync.Mutex
	EventosVendas chan venda.Venda
}

func InstanciarSimuladorVendasService() *SimuladorVendasService {
	return &SimuladorVendasService{}
}

func (service *SimuladorVendasService) InicializarSimuladorVendasService() {
	rand.Seed(time.Now().UnixNano())
	service.vendas = *venda.CriarCacheVendas()
	service.lojas = *loja.CriarCacheLojas()
	service.terminais = *terminal.CriarCacheTerminais()

	service.EventosVendas = make(chan venda.Venda, 100)

	service.lojas.Adicionar(loja.Loja{
		Numero: "1",
		Nome:   "Mercadinho 1",
	})
	service.lojas.Adicionar(loja.Loja{
		Numero: "2",
		Nome:   "Mercadinho 2",
	})

	service.terminais.Adicionar(terminal.Terminal{
		Numero:     "1",
		Nome:       "Terminal 1",
		NumeroLoja: "1",
	})
	service.terminais.Adicionar(terminal.Terminal{
		Numero:     "2",
		Nome:       "Terminal 2",
		NumeroLoja: "2",
	})
}

func (service *SimuladorVendasService) Iniciar(
	contexto context.Context,
	intervalo time.Duration,
) {
	ticker := time.NewTicker(intervalo)
	defer ticker.Stop()

	for {
		select {
		case <-contexto.Done():
			close(service.EventosVendas)
			return
		case <-ticker.C:
			novaVenda, ok := service.CarregarNovaVenda()
			if !ok {
				continue
			}
			if service.EventosVendas != nil {
				service.EventosVendas <- novaVenda
			}
		}
	}

}

func (service *SimuladorVendasService) CarregarNovaVenda() (venda.Venda, bool) {
	novaVenda, ok := service.gerarVendaAleatoria()
	if ok {
		service.mutex.Lock()
		service.totalVendas++
		service.valorVendas += novaVenda.Valor
		service.mutex.Unlock()

		service.vendas.Adicionar(novaVenda)

	}
	return novaVenda, ok
}

func (service *SimuladorVendasService) gerarVendaAleatoria() (venda.Venda, bool) {
	lojas := service.lojas.Listar()
	terminais := service.terminais.Listar()

	if len(lojas) == 0 || len(terminais) == 0 {
		return venda.Venda{}, false
	}

	lojaEscolhida := lojas[rand.Intn(len(lojas))]
	terminalEscolhido := terminais[rand.Intn(len(terminais))]

	valor := gerarValorAleatorio()

	novaVenda := venda.Venda{
		NumeroLoja:     lojaEscolhida.Numero,
		NumeroTerminal: terminalEscolhido.Numero,
		Valor:          valor,
		DataCriacao:    *model.DataAgora(),
	}

	return novaVenda, true
}

func gerarValorAleatorio() float64 {
	min := 10.0
	max := 100.0
	return min + rand.Float64()*(max-min)
}

func (service *SimuladorVendasService) ObterTotais() (int64, float64) {
	service.mutex.Lock()
	defer service.mutex.Unlock()

	return service.totalVendas, service.valorVendas
}
