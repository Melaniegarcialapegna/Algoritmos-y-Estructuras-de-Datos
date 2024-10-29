package pila

const (
	TAM_INICIAL                    = 6
	FACTOR_REDIMENSION_AUMENTO     = 2
	FACTOR_REDIMENSION_DISMINUCION = 2
	MULTIPLICADOR_REDUCIR_CANTIDAD = 4
)

// Definición del struct pila proporcionado por la cátedra.
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int // cantidad de elementos almacenados
}

// CrearPilaDinamica crea y devuelve una pila dinamica vacia con capacidad para un elemento.
func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, TAM_INICIAL)
	pila.cantidad = 0
	return pila
}

// EstaVacia devuelve VERDADERO si la pila no tiene elementos apilados, FALSO en caso contrario.
func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0

}

// VerTope obtiene el valor del tope de la pila. Si la pila tiene elementos se devuelve el valor del tope. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".
func (pila *pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[(pila.cantidad)-1]

}

// Apilar agrega un nuevo elemento a la pila.En caso de que la cantidad de elementos sea igual a la capacidad del slice, se llama a la funcion redimensionar para asi duplicar la capacidad de este.
func (pila *pilaDinamica[T]) Apilar(valor T) {
	if pila.cantidad == len(pila.datos) {
		pila.redimensionar(len(pila.datos) * FACTOR_REDIMENSION_AUMENTO)
	}
	pila.datos[pila.cantidad] = valor
	pila.cantidad++
}

// Desapilar saca el elemento tope de la pila. Si la pila tiene elementos, se quita el tope de la pila, y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La pila esta vacia".En caso de que la capacidad maxima del slice sea mayor o igual a 4 veces la cantidad de elementos, se llama a la funcion redimensionar para asi reducir su capacidad a la mitad.
func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	valor := pila.datos[pila.cantidad-1]
	pila.cantidad--
	if (!pila.EstaVacia()) && ((pila.cantidad * FACTOR_REDIMENSION_DISMINUCION) <= len(pila.datos)) {
		pila.redimensionar(len(pila.datos) / MULTIPLICADOR_REDUCIR_CANTIDAD)
	}
	return valor
}

// redimensionar se encarga de crear un nuevo slice y copiar los elementos de la pila en este nuevo.
func (pila *pilaDinamica[T]) redimensionar(tam int) {
	nuevosDatos := make([]T, tam)
	copy(nuevosDatos, pila.datos)
	pila.datos = nuevosDatos
}
