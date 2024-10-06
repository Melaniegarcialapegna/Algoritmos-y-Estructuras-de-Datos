package diccionario

import (
	"fmt"
)

const (
	VACIO               = 0
	OCUPADO             = 1
	BORRADO             = 2
	CANTIDAD_INICIAL    = 0
	TAM_INICIAL         = 10
	BORRADOS_INICIAL    = 0
	PRIMERA_POSICION    = 0
	FACTOR_REDIMENSION  = 3
	FACTOR_CARGA_MAXIMO = 0.7
	POSICION_INICIAL    = 0
)

type estado int

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	celdas := make([]celdaHash[K, V], TAM_INICIAL)
	return &hashCerrado[K, V]{tabla: celdas, cantidad: CANTIDAD_INICIAL, tam: TAM_INICIAL, borrados: BORRADOS_INICIAL}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	// Checkear si es necesario redimensionar la tabla
	if hash.factorCarga() > FACTOR_CARGA_MAXIMO {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}

	bytes := convertirABytes(clave)
	posicion := int(CityHash32(bytes)) % hash.tam

	//NOS FALTA VER SI LA CLAVE YA EXISTE -> REEMPLAZAR EL DATO
	for hash.tabla[posicion].estado == OCUPADO {
		if hash.tabla[posicion].clave == clave { //Si la clave ya existe -> reem dato
			hash.tabla[posicion].dato = dato
			return
		}
		if posicion == hash.tam-1 {
			posicion = PRIMERA_POSICION - 1
		}
		posicion++
	}
	hash.tabla[posicion].clave = clave
	hash.tabla[posicion].dato = dato
	hash.tabla[posicion].estado = OCUPADO
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	return hash.buscarElemento(clave) != -1
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	posicion := hash.buscarElemento(clave)
	if posicion == -1 {
		panic("La clave no pertenece al diccionario")
	}
	return hash.tabla[posicion].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	posicion := hash.buscarElemento(clave)
	if posicion == -1 {
		panic("La clave no pertenece al diccionario")
	}
	valor := hash.tabla[posicion].dato
	hash.tabla[posicion].estado = BORRADO
	hash.cantidad--
	hash.borrados++
	return valor
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, el := range hash.tabla {
		if el.estado == OCUPADO {
			if !f(el.clave, el.dato) {
				break
			}
		}
	}
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return &iteradorHashCerrado[K, V]{hash: hash, posicionActual: POSICION_INICIAL}
}

// tda iterador
type iteradorHashCerrado[K comparable, V any] struct {
	hash           *hashCerrado[K, V]
	posicionActual int
}

func (iterador *iteradorHashCerrado[K, V]) HaySiguiente() bool {
	for iterador.posicionActual < iterador.hash.tam {
		if iterador.hash.tabla[iterador.posicionActual].estado == OCUPADO {
			return true
		}
		iterador.posicionActual++
	}
	return false
}

func (iterador *iteradorHashCerrado[K, V]) VerActual() (K, V) {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.hash.tabla[iterador.posicionActual].clave, iterador.hash.tabla[iterador.posicionActual].dato
}

func (iterador *iteradorHashCerrado[K, V]) Siguiente() {
	iterador.posicionActual++
	if iterador.posicionActual > iterador.hash.tam {
		panic("El iterador termino de iterar")
	}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hashCerrado[K, V]) redimensionar(tam int) {
	tablaAnterior := hash.tabla
	hash.tabla = make([]celdaHash[K, V], tam)
	hash.tam = tam
	hash.cantidad = CANTIDAD_INICIAL
	hash.borrados = BORRADOS_INICIAL
	for _, elem := range tablaAnterior {
		if elem.estado == OCUPADO {
			hash.Guardar(elem.clave, elem.dato)
		}
	}
}

func (hash *hashCerrado[K, V]) factorCarga() float32 {
	return (float32(hash.cantidad) + float32(hash.borrados)) / float32(hash.tam)
}

func (hash *hashCerrado[K, V]) buscarElemento(clave K) int {
	bytes := convertirABytes(clave)
	posicion := int(CityHash32(bytes)) % hash.tam
	for hash.tabla[posicion].estado != VACIO {
		if hash.tabla[posicion].estado == OCUPADO && hash.tabla[posicion].clave == clave {
			return posicion
		}
		if posicion == hash.tam-1 {
			posicion = PRIMERA_POSICION - 1
		}
		posicion++
	}
	return -1
}
