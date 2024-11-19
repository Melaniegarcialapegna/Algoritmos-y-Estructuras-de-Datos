package lista

type Lista[T any] interface {

	//EstaVacia devuelve true si la lista no tiene elementos insertados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero agrega un nuevo elemento a la lista en la primera posicion.
	InsertarPrimero(T)

	// InsertarUltimo agrega un nuevo elemento a la lista en la ultima posicion.
	InsertarUltimo(T)

	// BorrarPrimero elimina el primer elemento de la lista. Si la lista tiene elementos se quita el primero, y devuelve ese valor.
	// Si la lista esta vacia, entra en panico con un mensaje 'La lista esta vacía'.
	BorrarPrimero() T

	// VerPrimero obtiene el valor del primer elemento de la lista. Si la lista tiene elementos, devuelve el valor del primero.
	//  Si la lista esta vacia, entra en panico con un mensaje 'La lista esta vacía'.
	VerPrimero() T

	// VerUltimo obtiene el valor del ultimo elemento de la lista. Si la lista tiene elementos, devuelve el valor del ultimo.
	//  Si la lista esta vacia, entra en panico con un mensaje 'La lista esta vacía'.
	VerUltimo() T

	//Largo devuelve la cantidad de elementos de la lista.
	Largo() int

	//Iterar permite iterar internamente la lista mientras la funcion pasada por parametro devuelve true.
	// Esta funcion se ejecuta una vez por cada elemento de la lista.
	Iterar(visitar func(T) bool)

	// Iterador permite iterar externamente la lista mediante sus metodos.
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {

	// VerActual obtiene el elemento actual de la lista que se esta iterando.
	VerActual() T

	// HaySiguiente devuelve true si existe un valor en la lista sobre el cual iterar.
	// En caso contrario devuelve false.
	HaySiguiente() bool

	// Siguiente mueve el puntero del iterador al siguiente elemento de la lista. En caso de que no haya siguiente
	// lanza un panic con un mensaje 'La lista esta vacía'.
	Siguiente()

	//Insertar permite insertar un nuevo elemento en la lista en la posicion anterior del actual del iterador.
	Insertar(T)

	//Borrar elimina el elemento actual de la lista y lo devuelve.
	Borrar() T
}
