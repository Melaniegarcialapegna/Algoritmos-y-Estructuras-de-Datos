package diccionario

import (
	"fmt"
)

const (
	VACIO               = 0
	OCUPADO             = 1
	BORRADO             = 2
	BORRADOS_INICIAL    = 0
	CANTIDAD_INICIAL    = 0
	FACTOR_CARGA_MAXIMO = 0.7
	FACTOR_CARGA_MINIMO = 0.3
	FACTOR_REDIMENSION  = 3
	POSICION_INICIAL    = 0
	PRIMERA_POSICION    = 0
	TAM_INICIAL         = 10
)

type estado int

// celdaHash representa una celda de una tabla de hash.
type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

// hashCerrado es la implementacion de una tabla de hash.
type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	cantidad int
	tam      int
	borrados int
}

// CrearHash crea y devuelve un diccionario implementado con una tabla de Hash cerrada.
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	celdas := make([]celdaHash[K, V], TAM_INICIAL)
	return &hashCerrado[K, V]{tabla: celdas, cantidad: CANTIDAD_INICIAL, tam: TAM_INICIAL, borrados: BORRADOS_INICIAL}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {

	// En caso de ser necesario, redimesiona la tabla de hash
	if hash.factorCarga() > FACTOR_CARGA_MAXIMO {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}

	bytes := convertirABytes(clave)
	posicion := int(CityHash32(bytes)) % hash.tam

	for hash.tabla[posicion].estado == OCUPADO {
		//Si la clave ya existe reemplaza el dato
		if hash.tabla[posicion].clave == clave {
			hash.tabla[posicion].dato = dato
			return
		}
		//Si llegamos a la ultima celda seguimos buscando desde la primera
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

	if hash.factorCarga() < FACTOR_CARGA_MINIMO {
		hash.redimensionar(hash.tam / FACTOR_REDIMENSION)
	}
	// En caso de ser necesario, redimesiona la tabla de hash

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

// iteradorHashCerrado representa un iterador para un hash cerrado.
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

// convertirABytes transforma un tipo de dato genérico a un array de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// redimensionar toma la tabla actual del hash y transfiere sus elementos a una nueva tabla de mayor tamaño
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

// factorCarga
func (hash *hashCerrado[K, V]) factorCarga() float32 {
	return (float32(hash.cantidad) + float32(hash.borrados)) / float32(hash.tam)
}

// buscarElemento
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
