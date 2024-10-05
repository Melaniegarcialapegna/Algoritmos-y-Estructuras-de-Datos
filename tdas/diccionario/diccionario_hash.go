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
	bytes := convertirABytes(clave)
	posicion := int(Hash32(bytes))

	// Checkear si es necesario redimensionar la tabla
	if hash.factorCarga() > FACTOR_CARGA_MAXIMO {
		hash.redimensionar(hash.tam * FACTOR_REDIMENSION)
	}

	for hash.tabla[posicion].estado == OCUPADO {
		if posicion == hash.tam-1 {
			posicion = PRIMERA_POSICION
		}
		posicion++
	}
	hash.tabla[posicion].clave = clave
	hash.tabla[posicion].dato = dato
	hash.tabla[posicion].estado = OCUPADO
	hash.cantidad++
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) bool {
	bytes := convertirABytes(clave)
	posicion := int(Hash32(bytes)) % hash.tam
	for 
	return hash.tabla[posicion].estado == OCUPADO
}

func (hash *hashCerrado[K, V]) Obtener(clave K) V {
	return hash.tabla[0].dato
}

func (hash *hashCerrado[K, V]) Borrar(clave K) V {
	return
}

func (hash *hashCerrado[K, V]) Cantidad() int {
	return hash.cantidad
}

func (hash *hashCerrado[K, V]) Iterar(func(clave K, dato V) bool) {
	return
}

func (hash *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	return
}

// tda iterador
type iteradorHashCerrado[K comparable, V any] struct {
	Laura int
	Va    int
}

func (iterador *iteradorHashCerrado) HaySiguiente() bool {
	return
}
func (iterador *iteradorHashCerrado) VerActual() (K, V) {
	return
}
func (iterador *iteradorHashCerrado) Siguiente() {
	return
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func (hash *hashCerrado[K, V]) redimensionar(tam int) {
	nuevosDatos := make([]celdaHash[K, V], tam)
	//for i := 0; i < hash.
	hash.tabla = nuevosDatos
	hash.tam = tam
}

func (hash *hashCerrado[K, V]) factorCarga() float32 {
	return (1 + float32(hash.cantidad) + float32(hash.borrados)) / float32(hash.tam)
}

func guardarElemento(celda celdaHash[K,V], tabla []celdaHash[K, V]){
	for tabla.[posicion].estado == OCUPADO {
		if posicion == hash.tam-1 {
			posicion = PRIMERA_POSICION
		}
		posicion++
	}
	hash.tabla[posicion].clave = clave
	hash.tabla[posicion].dato = dato
	hash.tabla[posicion].estado = OCUPADO
	hash.cantidad++
}