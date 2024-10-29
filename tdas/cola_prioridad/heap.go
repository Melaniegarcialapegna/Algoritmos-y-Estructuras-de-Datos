package cola_prioridad

// Normal Rockwell

const (
	TAM_INICIAL                    = 2
	FACTOR_REDIMENSION_AUMENTO     = 2
	FACTOR_REDIMENSION_DISMINUCION = 2
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearColaPrioridad[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, TAM_INICIAL)
	return &heap[T]{datos: arr, cantidad: 0, cmp: funcion_cmp}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(valor T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(heap.cantidad * FACTOR_REDIMENSION_AUMENTO)
	}
	heap.datos[heap.cantidad] = valor
	heap.upHeap(heap.cantidad)
	heap.cantidad++
}

func (heap *heap[T]) VerMax() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	return heap.datos[0]
}

func (heap *heap[T]) Desencolar() T {
	if heap.EstaVacia() {
		panic("La cola esta vacia")
	}
	borrado := heap.datos[0]
	heap.datos[0], heap.datos[heap.cantidad-1] = heap.datos[heap.cantidad-1], heap.datos[0]
	heap.cantidad--
	heap.downHeap(0)

	return borrado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

func (heap *heap[T]) upHeap(indice int) {
	for indice > 0 {
		padre := (indice - 1) / 2
		if heap.cmp(heap.datos[padre], heap.datos[indice]) > 0 {
			break
		}
		//Swap
		heap.datos[padre], heap.datos[indice] = heap.datos[indice], heap.datos[padre]
		indice = padre
	}
}

func (heap *heap[T]) downHeap(indice int) {
	for indice < heap.cantidad {
		hijoIzq := (2 * indice) + 1
		hijoDer := (2 * indice) + 2

		if hijoIzq >= heap.cantidad {
			break
		}

		indiceHijo := hijoIzq
		if (hijoDer < heap.cantidad) && (heap.cmp(heap.datos[hijoDer], heap.datos[indiceHijo]) > 0) {
			indiceHijo = hijoDer
		}

		if heap.cmp(heap.datos[indiceHijo], heap.datos[indice]) < 0 {
			break
		}
		//Swap
		heap.datos[indiceHijo], heap.datos[indice] = heap.datos[indice], heap.datos[indiceHijo]
		indice = indiceHijo
	}
}

func (heap *heap[T]) redimensionar(tam int) {
	nuevoArr := make([]T, tam)
	copy(nuevoArr, heap.datos)
	heap.datos = nuevoArr
}

//Hacer HeapSort
