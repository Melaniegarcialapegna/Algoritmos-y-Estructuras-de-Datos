package cola_prioridad_test

import (
	"strings"
	TDAColaPrioridad "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	TAM_VOLUMEN = 10000
)

func cmp(a, b int) int {
	return a - b
}

// TestColaPrioridadVacia se encarga de verificar que se pueda crear una cola vacia y que esta se comporte como tal.
func TestColaPrioridadVacia(t *testing.T) {
	//Creando una cola vacia para enteros.
	colaPrioridadEnteros := TDAColaPrioridad.CrearHeap[int](cmp)
	require.True(t, colaPrioridadEnteros.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.Equal(t, 0, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 0")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridadEnteros.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	//Creando una cola vacia para cadenas.
	colaCadena := TDAColaPrioridad.CrearHeap[string](strings.Compare)
	require.True(t, colaCadena.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadena.VerMax() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.Equal(t, 0, colaCadena.Cantidad(), "La cantidad de la cola deberia ser 0")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadena.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestEncolarPrioridadElementos se encarga de verificar que se puedan encolar elementos y que al desencolarlos se mantenga el invariante de cola. Verifica que salgan el orden deseado.
func TestEncolarPrioridadElementos(t *testing.T) {
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	colaPrioridad.Encolar(1)
	require.Equal(t, 1, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 1")
	require.Equal(t, 1, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 1")
	require.True(t, colaPrioridad.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.Equal(t, 0, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 0")
	colaPrioridad.Encolar(2)
	colaPrioridad.Encolar(3)
	colaPrioridad.Encolar(4)
	colaPrioridad.Encolar(5)
	colaPrioridad.Encolar(6)
	require.Equal(t, 5, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 5")
	require.Equal(t, 6, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 6")
	require.Equal(t, 6, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 6")
	require.Equal(t, 4, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 4")
	colaPrioridad.Encolar(7)
	colaPrioridad.Encolar(8)
	require.Equal(t, 6, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 6")
	require.Equal(t, 8, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 8")
	require.Equal(t, 8, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 8")
	require.Equal(t, 7, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 7")
	require.Equal(t, 5, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 5")
	require.Equal(t, 3, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 3")
	require.Equal(t, 4, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 4")
	require.False(t, colaPrioridad.EstaVacia(), "Deberia devolver False, ya que la cola NO deberia estar vacia")
	require.Equal(t, 4, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 4")
	require.Equal(t, 3, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 3")
	require.Equal(t, 2, colaPrioridad.Desencolar(), "Al desencolar deberia devolver el 2")
	require.Equal(t, 0, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 0")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.True(t, colaPrioridad.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestDesencolar se encarga de desencoalar hasta que la cola este vacia para luego comprobar que se comporte como recien creada.
func TestDesencolar(t *testing.T) {
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	//Encolamos y desencolamos hasta que la cola quede vacia
	for i := 0; i <= 4; i++ {
		colaPrioridad.Encolar(i)
	}
	for i := 0; i <= 4; i++ {
		colaPrioridad.Desencolar()
	}
	//Probamos que se comporte como una cola recien creada
	require.True(t, colaPrioridad.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.Equal(t, 0, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 0")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestVolumen se encarga de verificar que se puedan encolar y desencolar MUCHOS elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
func TestVolumen(t *testing.T) {
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	for i := 1; i <= TAM_VOLUMEN; i++ {
		colaPrioridad.Encolar(i)
	}
	require.Equal(t, TAM_VOLUMEN, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser %d", TAM_VOLUMEN)
	require.Equal(t, TAM_VOLUMEN, colaPrioridad.Desencolar(), "El primer elemento de la cola deberia ser %d", TAM_VOLUMEN)
	for i := TAM_VOLUMEN - 1; i > 1; i-- {
		require.Equal(t, i, colaPrioridad.Desencolar(), "Al desedncolar deberia devolver el %d", i)
		require.Equal(t, i-1, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser %d ", i-1)
	}
	require.False(t, colaPrioridad.EstaVacia(), "Deberia devolver False, ya que la cola NO deberia estar vacia")
	colaPrioridad.Desencolar()
	require.True(t, colaPrioridad.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestFuncionesInvalidasColaPrioridadNueva se encarga de verificar que las acciones de desencolar y ver_primero en una cola recién creada sean inválidas.
func TestFuncionesInvalidasColaPrioridadNueva(t *testing.T) {
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	require.Equal(t, 0, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 0")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.VerMax() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaPrioridad.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestEstaVacia verifica que la acción de esta_vacía en una cola recién creada sea verdadero.
func TestEstaVacia(t *testing.T) {
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	require.True(t, colaPrioridad.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestEncolarDistintosDatos prueba Encolar diferentes tipos de datos.
func TestEncolarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	colaPrioridad := TDAColaPrioridad.CrearHeap[int](cmp)
	require.Equal(t, 0, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 0")
	colaPrioridad.Encolar(1)
	require.Equal(t, 1, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 1")
	require.Equal(t, 1, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 1")
	colaPrioridad.Encolar(2)
	colaPrioridad.Encolar(3)
	colaPrioridad.Encolar(4)
	require.Equal(t, 4, colaPrioridad.Cantidad(), "La cantidad de la cola deberia ser 4")
	require.Equal(t, 4, colaPrioridad.VerMax(), "El primer elemento de la cola deberia ser 4")

	//Con cadenas
	colaPrioridadCadena := TDAColaPrioridad.CrearHeap[string](strings.Compare)
	colaPrioridadCadena.Encolar("A")
	require.Equal(t, 1, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 1")
	require.Equal(t, "A", colaPrioridadCadena.VerMax(), "El primer elemento de la cola deberia ser A")
	colaPrioridadCadena.Encolar("B")
	colaPrioridadCadena.Encolar("C")
	colaPrioridadCadena.Encolar("D")
	require.Equal(t, 4, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 4")
	require.Equal(t, "D", colaPrioridadCadena.VerMax(), "El primer elemento de la cola deberia ser D")

}

// TestDesencolarDistintosDatos prueba desencolar diferentes tipos de datos.
func TestDesencolarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	colaPrioridadEnteros := TDAColaPrioridad.CrearHeap[int](cmp)
	colaPrioridadEnteros.Encolar(1)
	require.Equal(t, 1, colaPrioridadEnteros.Desencolar(), "Al desencolar deberia devolver el 1")
	require.Equal(t, 0, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 0")
	colaPrioridadEnteros.Encolar(2)
	colaPrioridadEnteros.Encolar(3)
	colaPrioridadEnteros.Encolar(4)
	require.Equal(t, 3, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 3")
	require.Equal(t, 4, colaPrioridadEnteros.Desencolar(), "Al desencolar deberia devolver el 4")
	require.Equal(t, 2, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 2")
	require.Equal(t, 3, colaPrioridadEnteros.Desencolar(), "Al desencolar deberia devolver el 3")
	require.Equal(t, 1, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 1")
	require.Equal(t, 2, colaPrioridadEnteros.Desencolar(), "Al desencolar deberia devolver el 2")
	require.Equal(t, 0, colaPrioridadEnteros.Cantidad(), "La cantidad de la cola deberia ser 0")

	//Con cadenas
	colaPrioridadCadena := TDAColaPrioridad.CrearHeap[string](strings.Compare)
	colaPrioridadCadena.Encolar("A")
	require.Equal(t, "A", colaPrioridadCadena.Desencolar(), "Al desencolar deberia devolver el A")
	require.Equal(t, 0, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 0")
	colaPrioridadCadena.Encolar("B")
	colaPrioridadCadena.Encolar("C")
	colaPrioridadCadena.Encolar("D")
	require.Equal(t, 3, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 3")
	require.Equal(t, "D", colaPrioridadCadena.Desencolar(), "Al desencolar deberia devolver el D")
	require.Equal(t, 2, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 2")
	require.Equal(t, "C", colaPrioridadCadena.Desencolar(), "Al desencolar deberia devolver el C")
	require.Equal(t, 1, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 1")
	require.Equal(t, "B", colaPrioridadCadena.Desencolar(), "Al desencolar deberia devolver el B")
	require.Equal(t, 0, colaPrioridadCadena.Cantidad(), "La cantidad de la cola deberia ser 0")

}

// TestHeapArr
func TestHeapArr(t *testing.T) {
	arr := []int{2, 26, 22, 12, 24, 4, 12, 8}
	arrOrdenado := []int{26, 24, 22, 12, 12, 8, 4, 2}
	heap := TDAColaPrioridad.CrearHeapArr[int](arr, cmp)

	for i := 0; i < len(arr); i++ {
		require.Equal(t, arrOrdenado[i], heap.Desencolar())
	}

	require.True(t, heap.EstaVacia())
}

// // TestHeapSort
// func TestHeapSort(t *testing.T) {
// 	arr := []int{2, 26, 22, 12, 24, 4, 12, 8}
// 	arrOrdenado := []int{2, 4, 8, 12, 12, 22, 24, 26}
// 	arregloHeapSort := []

// 	for i := 0; i < len(arr); i++ {

// 	}
// }
