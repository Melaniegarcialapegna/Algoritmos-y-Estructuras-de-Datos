package diccionario

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
	condicion := abb.cmp(nodoPadre.clave, clave)
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

	_ ,nodo, encontrado := abb.buscarElemento(clave)
	if !encontrado {
		panic("La clave no pertenece al diccionario")
	}

	return nodo.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	
	nodoPadre, nodoActual, encontrado := abb.buscarNodo(clave)

	dato := nodoActual.dato
	if nodoActual.clave == abb.raiz.clave{
		if nodoActual.esHoja(){
			abb.raiz = nil
			abb.cantidad--
			return dato
		}else if nodoActual.unHijo(){
			if nodoActual.izquierdo != nil{
				abb.raiz = nodoActual.izquierdo
			} else{
				abb.raiz = nodoActual.derecho
			}
			abb.cantidad--
			return dato
		}
	}

	//Si es una hoja
	if nodoActual.esHoja(){
		if nodoPadre.izquierdo.clave == nodoActual.clave{
			nodoPadre.izquierdo = nil
		}else{
			nodoPadre.derecho = nil
		}
	} else if nodoActual.unHijo(){
		if nodoPadre.izquierdo.clave == nodoActual.clave{
			if nodoActual.izquierdo != nil{
				nodoPadre.izquierdo = nodo.izquierdo
			}else{
				nodoPadre.izquierdo = nodoActual.derecho
			}
		} else{
			if nodoActual.izquierdo != nil{
				nodoPadre.derecho = nodoActual.izquierdo
			}else{
				nodoPadre.derecho = nodoActual.derecho
			}
		}
	} else {//Si tengo dos hijitos
		nodoReemplazo := buscarReemplazante(nodoActual)
		if nodoActual.clave == abb.raiz.clave{
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

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	return
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return
}

func (abb *abb[K, V]) IterarRango(*K, *K, func(K, V) bool) {
	return
}

func (abb *abb[K, V]) IteradorRango(*K, *K) IterDiccionario[K, V] {
	return
}







//Nuestras funcion
func (abb *abb[K, V]) buscarElemento(clave K) (*nodoAbb[K, V], bool) {
	return abb.buscarNodo(nil,abb.raiz, clave)
}

func (abb *abb[K, V]) buscarNodo(nodoPadre *nodoAbb[K, V], nodoActual *nodoAbb[K, V], clave K) (*nodoAbb[K, V], bool) {
	//Caso Base
	if nodoActual == nil {
		return nil, false
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
	return nil, nil,false
}

func (nodo *nodoAbb[K,V]) esHoja() bool{
	return nodo.izquierdo == nil && nodo.derecho == nil
}

func (nodo *nodoAbb[K,V]) unHijo() bool{
	return (nodo.izquierdo != nil && nodo.derecho == nil) || (nodo.derecho == nil && nodo.derecho != nil)
}

func buscarReemplazante(nodo *nodoAbb[K, V]) *nodoAbb[K, V]{
	nodoPadre := nil
	nodoReemplazo = nodo.derecho.izquierdo // el mas chiquito de la derecha
	for nodoReemplazo.izquierdo != nil{
		nodoPadre = nodoReemplazo
		nodoReemplazo = nodoReemplazo.izquierdo
	}
	if nodoPadre != nil{
		nodoPadre.izquierdo = nil
	} else{
		nodo.derecho = nil
	}
	return nodoReemplazo
}