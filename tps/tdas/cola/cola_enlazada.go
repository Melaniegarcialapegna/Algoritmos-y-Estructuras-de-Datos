package cola

// nodoCola representa a un nodo de la cola enlazada.
type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

// colaEnlazada es la implementacion de la cola utilizando nodos.
type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

// CrearColaEnlazada crea y devuelve una cola enlazada vacia.
func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}

// crearNuevoNodo crea y devuelve un nuevo nodo.
func (cola *colaEnlazada[T]) crearNuevoNodo(valor T) *nodoCola[T] {
	nuevoNodo := new(nodoCola[T])
	nuevoNodo.dato = valor
	nuevoNodo.prox = nil
	return nuevoNodo
}

// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, falso en caso contrario.
func (cola *colaEnlazada[T]) EstaVacia() bool {
	return cola.primero == nil

}

// VerPrimero obtiene el valor del primero de la cola. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) VerPrimero() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.primero.dato

}

// Encolar agrega un nuevo elemento a la cola, al final de la misma.
func (cola *colaEnlazada[T]) Encolar(valor T) {
	nuevoNodo := cola.crearNuevoNodo(valor)

	if cola.EstaVacia() {
		cola.primero = nuevoNodo
	} else {
		cola.ultimo.prox = nuevoNodo
	}
	cola.ultimo = nuevoNodo
}

// Desencolar saca el primer elemento de la cola. Si la cola tiene elementos, se quita el primero de la misma, y se devuelve ese valor. Si está vacía, entra en pánico con un mensaje "La cola esta vacia".
func (cola *colaEnlazada[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	valor := cola.primero.dato
	cola.primero = cola.primero.prox
	if cola.primero == nil {
		cola.ultimo = nil
	}
	return valor

}
