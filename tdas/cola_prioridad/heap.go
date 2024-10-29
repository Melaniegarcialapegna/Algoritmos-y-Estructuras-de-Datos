package cola_prioridad

// Normal Rockwell

const (
	TAM_INICIAL                    = 2 //Tamaño inicial del heap
	FACTOR_REDIMENSION_AUMENTO     = 2 //Factor de aumento para redimensionar heap
	FACTOR_REDIMENSION_DISMINUCION = 2 //Factor de disminucion para redimensionar heap
)

// heap representa un elemento del heap???????
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
	heap := CrearHeap(funcion_cmp)
	//Aplicarle heapify al arreglo -> asi es O(N) SINO O(N LOG N)
	return &heap[T]{datos: arreglo, cantidad: 0, cmp: funcion_cmp}
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

// unHeap reordena el heap de abajo hacia arriba a partir de un indice
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

// downHeap reordena el heap de arriba hacia abajo a partir de un indice
func (heap *heap[T]) downHeap(indice int) {
	for indice < heap.cantidad {
		hijoIzq := (2 * indice) + 1
		hijoDer := (2 * indice) + 2

		if hijoIzq >= heap.cantidad {
			break //No tiene hijos
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

// redimensionar cambia el tamaño del array que contiene los elementos del heap
func (heap *heap[T]) redimensionar(tam int) {
	nuevoArr := make([]T, tam)
	copy(nuevoArr, heap.datos)
	heap.datos = nuevoArr
}

// HeapSort ordena los elementos de menor a mayor a partir de un arrayS con forma de heap
func HeapSort[T any](elementos []T, funcion_cmp func(T, T) int) {
	//1: Le aplicamos heapify a los elementos
	//Vemos el max y lo swap con el ultimo
	//Tenemos un coso log menos en el arr
	//Asi hasta llegar a 1 elem logico
}

// heapify convierte un array en un heap
func heapify[T any](elementos []T, funcion_cmp func(T, T) int) {
	//cantElem := len(elementos)
	//Hacer downheap desde elementos[cantElem -1] hasta el elem[0]
	//Este downheap es de ABAJO HACIA ARRIBA !!!!!!!
}

//Pruebas con el de CrearHeapArr
//Pruebas con heapsort
//En pruebas falta probar el prim CANTIDAD -> LO DE ABAJO!!
//Falta redim cuando tenemos muy poca cantidad y muchas celdas(??) -> en desencolar
