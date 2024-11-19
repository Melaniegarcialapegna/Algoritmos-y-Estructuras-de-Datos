package pila

type pilaDinamica[T any] struct {
	elementos []T
	cantidad  int
}

const TAMAÑO_INICIAL int = 4
const FACTOR_REDIMENSION int = 2

func CrearPilaDinamica[T any]() Pila[T] {
	slice := make([]T, TAMAÑO_INICIAL)
	pila := &pilaDinamica[T]{slice, 0}
	return pila
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	return pila.elementos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) redimensionar(nuevoTamaño int) {
	nuevosElementos := make([]T, nuevoTamaño)
	copy(nuevosElementos, pila.elementos)
	pila.elementos = nuevosElementos
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if pila.cantidad == len(pila.elementos) {
		pila.redimensionar(len(pila.elementos) * FACTOR_REDIMENSION)
	}

	pila.elementos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}

	elemento := pila.elementos[pila.cantidad-1]
	pila.cantidad--

	if pila.cantidad <= len(pila.elementos)/(2*FACTOR_REDIMENSION) && len(pila.elementos) > TAMAÑO_INICIAL {
		pila.redimensionar(len(pila.elementos) / FACTOR_REDIMENSION)
	}

	return elemento
}
