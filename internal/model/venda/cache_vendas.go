package venda

import "painel-vendas-realtime/internal/model"

type CacheVendas struct {
	vendas model.SafeValue[[]Venda]
}

func CriarCacheVendas() *CacheVendas {
	return &CacheVendas{}
}

func (cache *CacheVendas) Adicionar(venda Venda) {
	cache.vendas.Executar(func(vendas *[]Venda) {
		*vendas = append(*vendas, venda)
	})
}

func (cache *CacheVendas) RemoverPorNumero(numero string) {
	cache.vendas.Executar(func(lista *[]Venda) {
		vendas := *lista

		indice := -1

		for i, venda := range vendas {
			if venda.Numero == numero {
				indice = i
				break
			}
		}

		if indice == -1 {
			return
		}

		*lista = append(vendas[:indice], vendas[indice+1:]...)
	})
}

func (cache *CacheVendas) Listar() []Venda {
	vendas := cache.vendas.Obter()

	copia := make([]Venda, len(vendas))
	copy(copia, vendas)

	return copia
}

func (cache *CacheVendas) BuscarPorNumero(numero string) (Venda, bool) {
	vendas := cache.vendas.Obter()

	for _, venda := range vendas {
		if venda.Numero == numero {
			return venda, true
		}
	}

	var vazio Venda
	return vazio, false
}
