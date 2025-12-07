package loja

import (
	"painel-vendas-realtime/internal/model"
)

type CacheLojas struct {
	lojas model.SafeValue[[]Loja]
}

func CriarCacheLojas() *CacheLojas {
	return &CacheLojas{}
}

func (cache *CacheLojas) Adicionar(loja Loja) {
	cache.lojas.Executar(func(lojas *[]Loja) {
		*lojas = append(*lojas, loja)
	})
}

func (cache *CacheLojas) RemoverPorNumero(numero string) {
	cache.lojas.Executar(func(lista *[]Loja) {
		lojas := *lista

		indice := -1

		for i, loja := range lojas {
			if loja.Numero == numero {
				indice = i
				break
			}
		}

		if indice == -1 {
			return
		}

		*lista = append(lojas[:indice], lojas[indice+1:]...)
	})
}

func (cache *CacheLojas) Listar() []Loja {
	lojas := cache.lojas.Obter()

	copia := make([]Loja, len(lojas))
	copy(copia, lojas)

	return copia
}

func (cache *CacheLojas) BuscarPorNumero(numero string) (Loja, bool) {
	lojas := cache.lojas.Obter()

	for _, loja := range lojas {
		if loja.Numero == numero {
			return loja, true
		}
	}

	var vazio Loja
	return vazio, false
}
