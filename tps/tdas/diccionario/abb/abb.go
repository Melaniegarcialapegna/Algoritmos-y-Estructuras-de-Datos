package diccionario

import (
	TDAPila "tdas/pila"
)

// nodoAbb representa un nodo del arbol binario de busqueda.
type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

// abb es la implementacion de un arbol binario de busqueda.
type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

// crearNodoAbb crea y devuelve un nuevo nodoAbb.
func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

// CrearABB crea y devuelve un diccionario implementado con un abb
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodoNuevo := crearNodoAbb[K, V](clave, dato)

	nodoPadre, nodo, encontrado := abb.buscarElemento(clave)
	if !encontrado {
		if nodoPadre == nil {
			abb.raiz = nodoNuevo
		} else if abb.cmp(nodoPadre.clave, clave) > 0 {
			nodoPadre.izquierdo = nodoNuevo
		} else {
			nodoPadre.derecho = nodoNuevo
		}
		abb.cantidad++
	} else {
		nodo.dato = dato //Si la clave ya existe actualiza el dato
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	_, _, encontrado := abb.buscarElemento(clave)
	return encontrado
}

func (abb *abb[K, V]) Obtener(clave K) V {
	_, nodo, encontrado := abb.buscarElemento(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	nodoPadre, nodoActual, encontrado := abb.buscarElemento(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}
	dato := nodoActual.dato

	if nodoActual.dosHijos() { //Si tiene dos hijos
		nodoReemplazo := abb.buscarReemplazante(nodoActual) //Buscamos un reemplazante para el nodo que queremos eliminar
		claveReemplazo := nodoReemplazo.clave
		datoReemplazo := nodoReemplazo.dato
		abb.Borrar(nodoReemplazo.clave)         //Borramos el reemplazante
		if nodoActual.clave == abb.raiz.clave { //Si es raiz y tiene dos hijos
			abb.raiz.clave = claveReemplazo
			abb.raiz.dato = datoReemplazo
		} else { //Si tiene dos hijos y NO es raiz
			dato = nodoActual.dato
			nodoActual.clave = claveReemplazo
			nodoActual.dato = datoReemplazo
		}
		abb.cantidad++

	} else if nodoActual.clave == abb.raiz.clave { //Si es raiz y tiene uno o ningun hijo
		if nodoActual.izquierdo != nil {
			abb.raiz = nodoActual.izquierdo
		} else {
			abb.raiz = nodoActual.derecho
		}
	} else { //Si NO es raiz y tiene uno o dos hijos
		if nodoPadre.izquierdo != nil && nodoPadre.izquierdo == nodoActual {
			abb.borrarUnHijo(nodoPadre, nodoActual, intercambiarIzquierdo)
		} else {
			abb.borrarUnHijo(nodoPadre, nodoActual, intercambiarDerecho)
		}
	}

	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(funcion func(clave K, dato V) bool) {
	abb.iterar(abb.raiz, funcion)
}

// iterar es una funcion interna de Iterar
func (abb *abb[K, V]) iterar(nodo *nodoAbb[K, V], funcion func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}

	if !abb.iterar(nodo.izquierdo, funcion) {
		return false
	}
	if !funcion(nodo.clave, nodo.dato) {
		return false
	}
	return abb.iterar(nodo.derecho, funcion)
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

// ITERADORES POR RANGO
// Iterador interno por rangos
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, funcion func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, funcion)
}

// iterarRango es un wrapper de IterarRango
func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, funcion func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	if desde == nil || abb.cmp(nodo.clave, *desde) > 0 { // Si el nodo es mayor a DESDE
		if !(abb.iterarRango(nodo.izquierdo, desde, hasta, funcion)) {
			return false
		}
	}

	if (desde == nil || (abb.cmp(nodo.clave, *desde) >= 0)) && (hasta == nil || (abb.cmp(nodo.clave, *hasta) <= 0)) { //Si el nodo esta dentro del rango que queremos iterar
		if !funcion(nodo.clave, nodo.dato) { //Le aplicamos la funcion al nodo
			return false
		}
	}

	if hasta == nil || abb.cmp(nodo.clave, *hasta) < 0 { //Si el nodo es manor a HASTA
		if !(abb.iterarRango(nodo.derecho, desde, hasta, funcion)) {
			return false
		}
	}
	return true
}

// Iterador externo por rangos
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter := &iteradorABB[K, V]{arbol: abb, pila: pila, funcionCmp: abb.cmp, desde: desde, hasta: hasta}
	if abb.cantidad > 0 {
		iter.apilarIzquierdos(abb.raiz)
	}
	return iter
}

// iteradorABB representa un iterador para el arbol binario de busqueda.
type iteradorABB[K comparable, V any] struct {
	arbol      *abb[K, V]
	pila       TDAPila.Pila[*nodoAbb[K, V]]
	funcionCmp func(K, K) int
	desde      *K
	hasta      *K
}

func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	actual := iter.pila.Desapilar()
	iter.apilarIzquierdos(actual.derecho)
}

// Funciones Auxiliares

// intercambiarIzquierdo intercambia el hijo izquiero del padre por el nuevo hijo
func intercambiarIzquierdo[K comparable, V any](padre, hijo *nodoAbb[K, V]) {
	padre.izquierdo = hijo
}

// intercambiarDerecho intercambia el hijo derecho del padre por el nuevo hijo
func intercambiarDerecho[K comparable, V any](padre, hijo *nodoAbb[K, V]) {
	padre.derecho = hijo
}

// borrarUnHijo se encarga de borrar hojas o nodos con un solo hijo
func (abb *abb[K, V]) borrarUnHijo(nodoPadre, nodoActual *nodoAbb[K, V], criterio func(*nodoAbb[K, V], *nodoAbb[K, V])) {
	if nodoActual.izquierdo != nil {
		criterio(nodoPadre, nodoActual.izquierdo)
	} else {
		criterio(nodoPadre, nodoActual.derecho)
	}
}

// buscarElemento busca un nodo por su clave y devuelve su padre, el nodo y un bool que indica si encontro la clave.
func (abb *abb[K, V]) buscarElemento(clave K) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	return abb.buscarNodo(nil, abb.raiz, clave)
}

// buscarNodo es una funcion interna de buscarElemento, busca la clave de forma recursiva.
func (abb *abb[K, V]) buscarNodo(nodoPadre *nodoAbb[K, V], nodoActual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	//Caso Base
	if nodoActual == nil {
		return nodoPadre, nil, false
	}

	condicion := abb.cmp(nodoActual.clave, clave)
	if condicion > 0 { //nodo > clave
		return abb.buscarNodo(nodoActual, nodoActual.izquierdo, clave)
	} else if condicion < 0 { //nodo < clave
		return abb.buscarNodo(nodoActual, nodoActual.derecho, clave)
	}
	//clave = nodo
	return nodoPadre, nodoActual, true
}

// unHijo devuelve TRUE si el nodo tiene DOS hijos.
func (nodo *nodoAbb[K, V]) dosHijos() bool {
	return nodo.izquierdo != nil && nodo.derecho != nil
}

// buscarReemplazante busca el nodo reemplazante para el nodo que queremos eliminar.
func (abb *abb[K, V]) buscarReemplazante(nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	nodoReemplazo := nodo.derecho // el mas chico a la derecha del nodo
	for nodoReemplazo.izquierdo != nil {
		nodoReemplazo = nodoReemplazo.izquierdo
	}
	return nodoReemplazo
}

// apilarIzquierdos apila todos los nodos que se encuentren a la izquierda del nodo pasado por parametro.
func (iter *iteradorABB[K, V]) apilarIzquierdos(nodo *nodoAbb[K, V]) {
	for nodo != nil { //Mientras haya un nodo a la izquierda
		if (iter.desde == nil || iter.funcionCmp(*iter.desde, nodo.clave) <= 0) && (iter.hasta == nil || iter.funcionCmp(*iter.hasta, nodo.clave) >= 0) { //Si esta dentro del rango
			iter.pila.Apilar(nodo)
		}
		if (iter.desde != nil) && (iter.funcionCmp(*iter.desde, nodo.clave) > 0) {
			nodo = nodo.derecho
			continue
		}
		nodo = nodo.izquierdo
	}
}
