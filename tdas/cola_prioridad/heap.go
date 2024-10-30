package cola_prioridad

// Normal Rockwell

const (
	TAM_INICIAL                    = 6                              //Tamaño inicial del heap
	FACTOR_REDIMENSION_AUMENTO     = 2                              //Factor de aumento para redimensionar heap
	FACTOR_REDIMENSION_DISMINUCION = 2 * FACTOR_REDIMENSION_AUMENTO //Factor de disminucion para redimensionar heap
	MULTIPLICADOR_REDUCIR_CANTIDAD = 4                              //Factor de multiplicacion para redimensionar heap
)

// heap representa un heap implementado con un arreglo
type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

// CrearHeap crea y devuelve un heap vacio con una funcion de comparacion
func CrearHeap[T any](funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arr := make([]T, TAM_INICIAL)
	return &heap[T]{datos: arr, cantidad: 0, cmp: funcion_cmp}
}

// CrearHeapArr crea y devuelve un heap a partir de un arreglo y una funcion de comparacion
func CrearHeapArr[T any](arreglo []T, funcion_cmp func(T, T) int) ColaPrioridad[T] {
	arregloNuevo := make([]T, len(arreglo))
	copy(arregloNuevo, arreglo)
	if len(arregloNuevo) > 0 {
		arregloNuevo = heapify(arregloNuevo, len(arregloNuevo), funcion_cmp)
	}
	return &heap[T]{datos: arregloNuevo, cantidad: len(arreglo), cmp: funcion_cmp}
}

func (heap *heap[T]) EstaVacia() bool {
	return heap.cantidad == 0
}

func (heap *heap[T]) Encolar(valor T) {
	if heap.cantidad == len(heap.datos) {
		heap.redimensionar(heap.cantidad * FACTOR_REDIMENSION_AUMENTO)
	}
	heap.datos[heap.cantidad] = valor
	heap.datos = upHeap(heap.datos, heap.cantidad, heap.cmp)
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
	heap.datos = downHeap(heap.datos, 0, heap.cantidad, heap.cmp)

	if (!heap.EstaVacia()) && (heap.cantidad <= (len(heap.datos) / (FACTOR_REDIMENSION_DISMINUCION))) {
		heap.redimensionar(len(heap.datos) / MULTIPLICADOR_REDUCIR_CANTIDAD)
	}

	return borrado
}

func (heap *heap[T]) Cantidad() int {
	return heap.cantidad
}

// unHeap reordena el heap de abajo hacia arriba a partir de un indice
func upHeap[T any](arreglo []T, indice int, funcion_cmp func(T, T) int) []T {
	for indice > 0 {
		padre := (indice - 1) / 2
		if funcion_cmp(arreglo[padre], arreglo[indice]) > 0 {
			break
		}
		//Swap
		arreglo[padre], arreglo[indice] = arreglo[indice], arreglo[padre]
		indice = padre
	}
	return arreglo
}

// downHeap reordena el heap de arriba hacia abajo a partir de un indice
func downHeap[T any](arreglo []T, indice int, cantidad int, funcion_cmp func(T, T) int) []T {
	for indice < cantidad {
		hijoIzq := (2 * indice) + 1
		hijoDer := (2 * indice) + 2

		if hijoIzq >= cantidad {
			break //No tiene hijos
		}

		indiceHijo := hijoIzq
		if (hijoDer < cantidad) && (funcion_cmp(arreglo[hijoDer], arreglo[indiceHijo]) > 0) {
			indiceHijo = hijoDer
		}

		if funcion_cmp(arreglo[indiceHijo], arreglo[indice]) < 0 {
			break
		}
		//Swap
		arreglo[indiceHijo], arreglo[indice] = arreglo[indice], arreglo[indiceHijo]
		indice = indiceHijo
	}
	return arreglo
}

// redimensionar cambia el tamaño del array que contiene los elementos del heap
func (heap *heap[T]) redimensionar(tam int) {
	nuevoArr := make([]T, tam)
	copy(nuevoArr, heap.datos)
	heap.datos = nuevoArr
}

// HeapSort ordena los elementos de menor a mayor a partir de un arrayS con forma de heap
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	//1: Le aplicamos heapify a los elementos
	heapify(elementos, len(elementos), funcion_cmp)
	for i := len(elementos) - 1; i > 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downHeap(elementos, 0, i, funcion_cmp)
	}
}

// heapify convierte un array en un heap
func heapify[T any](elementos []T, cantidad int, funcion_cmp func(T, T) int) []T {
	indice := cantidad
	for indice > 0 {
		elementos = downHeap(elementos, indice-1, cantidad, funcion_cmp)
		indice--
	}
	return elementos
}
