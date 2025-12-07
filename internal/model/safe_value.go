package model

import "sync"

type SafeValue[T any] struct {
	mutex sync.RWMutex
	valor T
}

func (safeValue *SafeValue[T]) Executar(funcao func(*T)) {
	safeValue.mutex.Lock()
	defer safeValue.mutex.Unlock()
	funcao(&safeValue.valor)
}

func (safeValue *SafeValue[T]) Obter() T {
	safeValue.mutex.RLock()
	defer safeValue.mutex.RUnlock()
	return safeValue.valor
}

func (safeValue *SafeValue[T]) ObterCom(funcao func(T)) {
	safeValue.mutex.RLock()
	defer safeValue.mutex.RUnlock()
	funcao(safeValue.valor)
}

func (safeValue *SafeValue[T]) Definir(novoValor T) {
	safeValue.mutex.Lock()
	defer safeValue.mutex.Unlock()
	safeValue.valor = novoValor
}
