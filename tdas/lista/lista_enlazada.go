package lista

// nodoLista representa un nodo de la lista enlazada.
type nodoLista[T any] struct {
	dato T
	sig  *nodoLista[T]
}

// listaEnlazada es la implementacion de la interfaz lista utilizando nodos.
type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

// crearNodo crea y devuelve un puntero a un nuevo nodo.
func crearNodo[T any](valor T, siguiente *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: valor, sig: siguiente}
}

// CrearListaEnlazada crea y devuelve una lista enlazada vacia.
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevoNodo := crearNodo(valor, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevoNodo := crearNodo(valor, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.sig = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	valor := lista.primero.dato
	lista.primero = lista.primero.sig
	lista.largo--
	if lista.primero == nil {
		lista.ultimo = nil
	}
	return valor
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	continuar := true
	for actual != nil && continuar {
		continuar = visitar(actual.dato)
		actual = actual.sig
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista: lista, actual: lista.primero, anterior: nil}
}

// iteradorListaEnlazada representa un iterador para una lista enlazada.
type iteradorListaEnlazada[T any] struct {
	lista    *listaEnlazada[T]
	actual   *nodoLista[T]
	anterior *nodoLista[T]
}

func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.sig
}

func (iterador *iteradorListaEnlazada[T]) Insertar(valor T) {
	nuevoNodo := crearNodo(valor, iterador.actual)

	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.sig = nuevoNodo
	}
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	valorBorrado := iterador.actual.dato

	//Cuando estamos borrando el primero
	if iterador.anterior == nil {
		iterador.lista.primero = iterador.actual.sig
		if iterador.lista.primero == nil {
			iterador.lista.ultimo = nil
		}
	} else { //Si estamos borrando uno del medio o el ultimo
		iterador.anterior.sig = iterador.actual.sig
		if iterador.actual.sig == nil { //Si estamos borrando el ultimo
			iterador.lista.ultimo = iterador.anterior
		}
	}
	//Caso general
	iterador.actual = iterador.actual.sig
	iterador.lista.largo--
	return valorBorrado
}
