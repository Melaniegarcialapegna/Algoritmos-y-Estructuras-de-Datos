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
	nodoPadre, encontrado := abb.buscarNodo(clave)
	if !encontrado{
		if nodoPadre.izquierdo == nil{
			nodoPadre.izquierdo = nodoNuevo
		} else{
			nodoPadre.derecho = nodoNuevo
		}
		abb.cantidad++
	} else if nodoPadre.izquierdo.clave == clave{
		nodoPadre.izquierdo = nodoNuevo
	} else{
		nodoPadre.derecho = nodoNuevo
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	if abb.raiz == nil {
		return false
	}
	_, encontrado := abb.buscarNodo(clave)
	return encontrado
}

func (abb *abb[K, V]) Obtener(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}

	nodoPadre , encontrado := abb.buscarNodo(clave)
	if encontrado{
		if nodoPadre.izquierdo.clave == clave{
			return nodoPadre.izquierdo.dato
		} else{
			return nodoPadre.derecho.dato
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}
	if !abb.Pertenece(clave) {
		panic("La clave no pertenece al diccionario")
	}
	
	nodoActual := abb.raiz
	valor, borrado := abb.borrarNodo(abb.raiz, clave)
	if borrado{
		return valor
	}
	
	for nodoActual != nil{
		valor, borrado := abb.borrarNodo(nodoActual.izquierdo, clave)
		if borrado{
			return valor
		}
		if !borrado{
			abb.borrarNodo(nodoActual.derecho, clave)
		} else {//Si no es izquierdo ni el derecho
			condicion := abb.cmp(clave, nodoActual.clave)
			//Si el nuevo es mas chico que el actual
			if condicion < 0 {
				nodoActual = nodoActual.izquierdo 
			} else { //Si el nuevo es mas grande que el actual
				nodoActual = nodoActual.derecho
			}
		}
	}
	panic("La clave no pertenece al diccionario")
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

func (abb *abb[K, V]) buscarNodo(clave K) (*nodoAbb[K, V], bool){
	var nodoPadre *nodoAbb[K, V] = nil
	nodoActual := abb.raiz
	for nodoActual != nil {
		condicion := abb.cmp(clave, nodoActual.clave)
		//-1 clave < nodoActual.clave
		//1 clave > nodoActual.clave
		//0 clave = nodoActual.clave

		//Si el nuevo es mas chico que el actual
		if condicion < 0 {
			nodoPadre = nodoActual
			nodoActual = nodoActual.izquierdo
		} else if condicion > 0 { //Si el nuevo es mas grande que el actual
			nodoPadre = nodoActual
			nodoActual = nodoActual.derecho
		} else { //Si el nuevo es igual que el actual
			return nodoPadre, true
		}
	}
	return nodoPadre, false
}

func (abb *abb[K, V]) borrarNodo(nodoActual *nodoAbb[K, V], clave K) (V, bool){
	if nodoActual.clave == clave{
		//Si no tiene hijos
		if nodoActual.izquierdo == nil && nodoActual.derecho == nil {
			nodoActual = nil
		}
		
		//Tiene solo un hijo
		if nodoActual.izquierdo == nil && nodoActual.derecho != nil{
			nodoActual = nodoActual.derecho
		}
		if nodoActual.izquierdo != nil && nodoActual.derecho == nil{
			nodoActual = nodoActual.izquierdo

		}

		//Tiene dos hijos
		if nodoActual.izquierdo != nil && nodoActual.derecho != nil{
			nodoReemplazo := nodoPadre.derecho.izquierdo // el mas chiquito de la derecha
			for nodoReemplazo.izquierdo != nil{
				nodoReemplazo = nodoReemplazo.izquierdo
			}
			nodoActual.clave = nodoReemplazo.clave
			nodoActual.dato = nodoReemplazo.dato
		}
		abb.cantidad--
		return nodoActual.dato, true
	}
	return nodoActual.dato, false
}