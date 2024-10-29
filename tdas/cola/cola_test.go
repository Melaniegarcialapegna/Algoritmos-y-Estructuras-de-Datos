package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestColaVacia se encarga de verificar que se pueda crear una cola vacia y que esta se comporte como tal.
func TestColaVacia(t *testing.T) {
	//Creando una cola vacia para enteros.
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	require.True(t, colaEnteros.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaEnteros.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	//Creando una cola vacia para cadenas.
	colaCadena := TDACola.CrearColaEnlazada[string]()
	require.True(t, colaCadena.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadena.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { colaCadena.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestEncolarElementos se encarga de verificar que se puedan encolar elementos y que al desencolarlos se mantenga el invariante de cola. Verifica que salgan el orden deseado.
func TestEncolarElementos(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.Equal(t, 1, cola.VerPrimero(), "El primer elemento de la cola deberia ser 1")
	require.Equal(t, 1, cola.Desencolar(), "Al desencolar deberia devolver el 1")
	require.True(t, cola.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	cola.Encolar(5)
	cola.Encolar(6)
	require.Equal(t, 2, cola.VerPrimero(), "El primer elemento de la cola deberia ser 2")
	require.Equal(t, 2, cola.Desencolar(), "Al desencolar deberia devolver el 2")
	cola.Encolar(7)
	cola.Encolar(8)
	require.Equal(t, 3, cola.VerPrimero(), "El primer elemento de la cola deberia ser 3")
	require.Equal(t, 3, cola.Desencolar(), "Al desencolar deberia devolver el 3")
	require.Equal(t, 4, cola.Desencolar(), "Al desencolar deberia devolver el 4")
	require.Equal(t, 5, cola.Desencolar(), "Al desencolar deberia devolver el 5")
	require.Equal(t, 6, cola.VerPrimero(), "El primer elemento de la cola deberia ser 6")
	require.False(t, cola.EstaVacia(), "Deberia devolver False, ya que la cola NO deberia estar vacia")
	require.Equal(t, 6, cola.Desencolar(), "Al desencolar deberia devolver el 6")
	require.Equal(t, 7, cola.Desencolar(), "Al desencolar deberia devolver el 7")
	require.Equal(t, 8, cola.Desencolar(), "Al desencolar deberia devolver el 8")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.True(t, cola.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestVolumen se encarga de verificar que se puedan encolar y desencolar MUCHOS elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
func TestVolumen(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 1; i <= 10000; i++ {
		cola.Encolar(i)
	}
	require.Equal(t, 1, cola.VerPrimero(), "El primer elemento de la cola deberia ser 1")
	for i := 1; i <= 9999; i++ {
		require.Equal(t, i, cola.Desencolar(), "Al desedncolar deberia devolver el %d", i)
		require.Equal(t, i+1, cola.VerPrimero(), "El primer elemento de la cola deberia ser %d ", i+1)
	}
	require.False(t, cola.EstaVacia(), "Deberia devolver False, ya que la cola NO deberia estar vacia")
	cola.Desencolar()
	require.True(t, cola.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestDesencolar se encarga de desencoalar hasta que la cola este vacia para luego comprobar que se comporte como recien creada.
func TestDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	//Encolamos y desencolamos hasta que la cola quede vacia
	for i := 0; i <= 4; i++ {
		cola.Encolar(i)
	}
	for i := 0; i <= 4; i++ {
		cola.Desencolar()
	}
	//Probamos que se comporte como una cola recien creada
	require.True(t, cola.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestFuncionesInvalidasColaNueva se encarga de verificar que las acciones de desencolar y ver_primero en una cola recién creada sean inválidas.
func TestFuncionesInvalidasColaNueva(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestEstaVacia verifica que la acción de esta_vacía en una cola recién creada sea verdadero.
func TestEstaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia(), "Deberia devolver True, ya que la cola deberia estar vacia")
}

// TestFuncionesInvalidasColaVacia se encarga de verificar que las acciones de desencolar y primero en una cola a la que se le encolo y desencolo hasta estar vacía sean inválidas.
func TestFuncionesInvalidasColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	//Encolamos y desencolamos hasta que la cola quede vacia
	for i := 0; i <= 4; i++ {
		cola.Encolar(i)
	}
	for i := 0; i <= 4; i++ {
		cola.Desencolar()
	}
	//Verificamos que desencolar y ver tope son acciones invalidas
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() }, "Error: deberia devolver ‘La cola esta vacia‘")
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() }, "Error: deberia devolver ‘La cola esta vacia‘")
}

// TestEncolarDistintosDatos prueba Encolar diferentes tipos de datos.
func TestEncolarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaEnteros.Encolar(1)
	require.Equal(t, 1, colaEnteros.VerPrimero(), "El primer elemento de la cola deberia ser 1")
	colaEnteros.Encolar(2)
	colaEnteros.Encolar(3)
	colaEnteros.Encolar(4)
	require.Equal(t, 1, colaEnteros.VerPrimero(), "El primer elemento de la cola deberia ser 1")

	//Con cadenas
	colaCadena := TDACola.CrearColaEnlazada[string]()
	colaCadena.Encolar("A")
	require.Equal(t, "A", colaCadena.VerPrimero(), "El primer elemento de la cola deberia ser A")
	colaCadena.Encolar("B")
	colaCadena.Encolar("C")
	colaCadena.Encolar("D")
	require.Equal(t, "A", colaCadena.VerPrimero(), "El primer elemento de la cola deberia ser A")

	//Con numeros con coma
	colaFloat := TDACola.CrearColaEnlazada[float64]()
	colaFloat.Encolar(4.1)
	require.Equal(t, 4.1, colaFloat.VerPrimero(), "El primer elemento de la cola deberia ser 4.1")
	colaFloat.Encolar(4.2)
	colaFloat.Encolar(4.3)
	colaFloat.Encolar(4.4)
	require.Equal(t, 4.1, colaFloat.VerPrimero(), "El primer elemento de la cola deberia ser 4.1 ")

}

// TestDesencolarDistintosDatos prueba desencolar diferentes tipos de datos.
func TestDesencolarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	colaEnteros := TDACola.CrearColaEnlazada[int]()
	colaEnteros.Encolar(1)
	require.Equal(t, 1, colaEnteros.Desencolar(), "Al desencolar deberia devolver el 1")
	colaEnteros.Encolar(2)
	colaEnteros.Encolar(3)
	colaEnteros.Encolar(4)
	require.Equal(t, 2, colaEnteros.Desencolar(), "Al desencolar deberia devolver el 2")
	require.Equal(t, 3, colaEnteros.Desencolar(), "Al desencolar deberia devolver el 3")
	require.Equal(t, 4, colaEnteros.Desencolar(), "Al desencolar deberia devolver el 4")

	//Con cadenas
	colaCadena := TDACola.CrearColaEnlazada[string]()
	colaCadena.Encolar("A")
	require.Equal(t, "A", colaCadena.Desencolar(), "Al desencolar deberia devolver el A")
	colaCadena.Encolar("B")
	colaCadena.Encolar("C")
	colaCadena.Encolar("D")
	require.Equal(t, "B", colaCadena.Desencolar(), "Al desencolar deberia devolver el B")
	require.Equal(t, "C", colaCadena.Desencolar(), "Al desencolar deberia devolver el C")
	require.Equal(t, "D", colaCadena.Desencolar(), "Al desencolar deberia devolver el D")

	//Con numeros con coma
	colaFloat := TDACola.CrearColaEnlazada[float64]()
	colaFloat.Encolar(4.1)
	require.Equal(t, 4.1, colaFloat.Desencolar(), "Al desencolar deberia devolver el 4.1")
	colaFloat.Encolar(4.2)
	colaFloat.Encolar(4.3)
	colaFloat.Encolar(4.4)
	require.Equal(t, 4.2, colaFloat.Desencolar(), "Al desencolar deberia devolver el 4.2")
	require.Equal(t, 4.3, colaFloat.Desencolar(), "Al desencolar deberia devolver el 4.3")
	require.Equal(t, 4.4, colaFloat.Desencolar(), "Al desencolar deberia devolver el 4.4")
}
