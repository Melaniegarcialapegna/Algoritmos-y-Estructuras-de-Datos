package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func crearNodoAbb[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

// InOrder -> prim izq desp yo desp der
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	nodoNuevo := crearNodoAbb[K, V](clave, dato)
	if abb.raiz == nil {
		abb.raiz = nodoNuevo
		abb.cantidad++
		return
	}
	_, nodo, encontrado := abb.buscarElemento(clave)
	condicion := abb.cmp(nodo.clave, clave)
	if !encontrado {
		if condicion > 0 {
			nodo.izquierdo = nodoNuevo
		} else {
			nodo.derecho = nodoNuevo
		}
		abb.cantidad++
	}
	nodo.dato = nodoNuevo.dato
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.raiz == nil {
		return false
	}
	_, _, encontrado := abb.buscarElemento(clave)
	return encontrado
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}

	_, nodo, encontrado := abb.buscarElemento(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}

	nodoPadre, nodoActual, _ := abb.buscarElemento(clave)

	dato := nodoActual.dato
	if nodoActual.clave == abb.raiz.clave {
		if nodoActual.esHoja() {
			abb.raiz = nil
			abb.cantidad--
			return dato
		} else if nodoActual.unHijo() {
			if nodoActual.izquierdo != nil {
				abb.raiz = nodoActual.izquierdo
			} else {
				abb.raiz = nodoActual.derecho
			}
			abb.cantidad--
			return dato
		}
	}

	//Si es una hoja
	if nodoActual.esHoja() {
		if nodoPadre.izquierdo.clave == nodoActual.clave {
			nodoPadre.izquierdo = nil
		} else {
			nodoPadre.derecho = nil
		}
	} else if nodoActual.unHijo() {
		if nodoPadre.izquierdo.clave == nodoActual.clave {
			if nodoActual.izquierdo != nil {
				nodoPadre.izquierdo = nodoActual.izquierdo
			} else {
				nodoPadre.izquierdo = nodoActual.derecho
			}
		} else {
			if nodoActual.izquierdo != nil {
				nodoPadre.derecho = nodoActual.izquierdo
			} else {
				nodoPadre.derecho = nodoActual.derecho
			}
		}
	} else { //Si tengo dos hijitos
		nodoReemplazo := buscarReemplazante(nodoActual)
		if nodoActual.clave == abb.raiz.clave {
			abb.raiz = nodoReemplazo
		}
		dato = nodoActual.dato
		nodoActual.clave = nodoReemplazo.clave
		nodoActual.dato = nodoReemplazo.dato
	}

	abb.cantidad--
	return dato
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(funcion func(clave K, dato V) bool) {
	//iteramos in order
	//izq
	//yo
	//der
	abb.iterar(abb.raiz, funcion)
}

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

type iteradorABB[K comparable, V any] struct {
	arbol      *abb[K, V]
	pila       TDAPila.Pila[*nodoAbb[K, V]]
	funcionCmp func(K, K) int
	desde      *K
	hasta      *K
}

// Iterador Rango

// interno rangos
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, funcion func(clave K, dato V) bool) {
	abb.iterarRango(abb.raiz, desde, hasta, funcion)
}

func (abb *abb[K, V]) iterarRango(nodo *nodoAbb[K, V], desde *K, hasta *K, funcion func(clave K, dato V) bool) bool {
	if nodo == nil {
		return true
	}
	// Vemos si el nodo actual es mayor a DESDE(inicio)

	if abb.cmp(nodo.clave, *desde) > 0 {
		seguirIzq := abb.iterarRango(nodo.izquierdo, desde, hasta, funcion)
		if !seguirIzq {
			return false
		}
	}

	if abb.cmp(nodo.clave, *desde) > 0 && abb.cmp(nodo.clave, *hasta) < 0 {
		if !funcion(nodo.clave, nodo.dato) { //:) Le aplico la fun al nodo y si dev FALSO termmino de iterar
			return false
		}
	}

	if abb.cmp(nodo.clave, *hasta) > 0 {
		seguirIzq := abb.iterarRango(nodo.izquierdo, desde, hasta, funcion)
		if !seguirIzq {
			return false
		}
	}
	return true
}

// externo rangos
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	nodoActual := abb.raiz
	for nodoActual != nil {
		pila.Apilar(nodoActual)
		nodoActual = nodoActual.izquierdo
	}
	return &iteradorABB[K, V]{arbol: abb, pila: pila, funcionCmp: abb.cmp, desde: desde, hasta: hasta}
}

// Iterador externo
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

	nodo := iter.pila.Desapilar()
	nodoActual := nodo.derecho
	for nodoActual != nil {
		iter.pila.Apilar(nodoActual)
		nodoActual = nodoActual.izquierdo
	}
}

// Nuestras funcion
func (abb *abb[K, V]) buscarElemento(clave K) (*nodoAbb[K, V], *nodoAbb[K, V], bool) {
	return abb.buscarNodo(nil, abb.raiz, clave)
}

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
	} else if condicion == 0 { //clave = nodo
		return nodoPadre, nodoActual, true
	} else if nodoActual.izquierdo == nil && nodoActual.derecho == nil { //Si estamos en una hoja
		return nodoPadre, nodoActual, false
	}
	return nil, nil, false
}

func (nodo *nodoAbb[K, V]) esHoja() bool {
	return nodo.izquierdo == nil && nodo.derecho == nil
}

func (nodo *nodoAbb[K, V]) unHijo() bool {
	return (nodo.izquierdo != nil && nodo.derecho == nil) || (nodo.derecho == nil && nodo.derecho != nil)
}

func buscarReemplazante[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	var nodoPadre *nodoAbb[K, V] = nil
	nodoReemplazo := nodo.derecho.izquierdo // el mas chiquito de la derecha
	for nodoReemplazo.izquierdo != nil {
		nodoPadre = nodoReemplazo
		nodoReemplazo = nodoReemplazo.izquierdo
	}
	if nodoPadre != nil {
		nodoPadre.izquierdo = nil
	} else {
		nodo.derecho = nil
	}
	return nodoReemplazo
}
