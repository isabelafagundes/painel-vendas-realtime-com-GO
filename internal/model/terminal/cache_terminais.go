package terminal

import "painel-vendas-realtime/internal/model"

type CacheTerminais struct {
	terminais model.SafeValue[[]Terminal]
}

func CriarCacheTerminais() *CacheTerminais {
	return &CacheTerminais{}
}

func (cache *CacheTerminais) Adicionar(venda Terminal) {
	cache.terminais.Executar(func(terminais *[]Terminal) {
		*terminais = append(*terminais, venda)
	})
}

func (cache *CacheTerminais) RemoverPorNumero(numero string) {
	cache.terminais.Executar(func(lista *[]Terminal) {
		terminais := *lista

		indice := -1

		for i, terminal := range terminais {
			if terminal.Numero == numero {
				indice = i
				break
			}
		}

		if indice == -1 {
			return
		}

		*lista = append(terminais[:indice], terminais[indice+1:]...)
	})
}

func (cache *CacheTerminais) Listar() []Terminal {
	terminais := cache.terminais.Obter()

	copia := make([]Terminal, len(terminais))
	copy(copia, terminais)

	return copia
}

func (cache *CacheTerminais) BuscarPorNumero(numero string) (Terminal, bool) {
	terminais := cache.terminais.Obter()

	for _, terminal := range terminais {
		if terminal.Numero == numero {
			return terminal, true
		}
	}

	var vazio Terminal
	return vazio, false
}
