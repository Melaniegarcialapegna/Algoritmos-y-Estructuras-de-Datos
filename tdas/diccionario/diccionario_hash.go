package diccionario

// import(
// 	"fmt"
// )

const (
	VACIO            = 0
	OCUPADO          = 1
	BORRADO          = 2
	CANTIDAD_INICIAL = 0
	TAM_INICIAL      = 10
	BORRADOS_INICIAL = 0
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
	celdas := []celdaHash[K, V]{}
	return &hashCerrado[K, V]{tabla: celdas, cantidad: CANTIDAD_INICIAL, tam: TAM_INICIAL, borrados: BORRADOS_INICIAL}
}

func (hash *hashCerrado[K, V]) Guardar(clave K, dato V) {
	return 2
}

func (hash *hashCerrado[K, V]) Pertenece(clave K) {
	return
}

func (hash *hashCerrado[K, V]) Obtener(clave K) {
	return
}

func (hash *hashCerrado[K, V]) Borrar(clave K) {
	return
}

func (hash *hashCerrado[K, V]) Cantidad(clave K) {
	return
}

func (hash *hashCerrado[K, V]) Iterar(clave K) {
	return
}

func (hash *hashCerrado[K, V]) Iterador(clave K) {
	return
}
