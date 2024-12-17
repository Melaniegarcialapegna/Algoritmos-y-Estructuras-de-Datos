package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestPilaVacia se encarga de verificar que se pueda crear una Pila vacía y que esta se comporte como tal.
func TestPilaVacia(t *testing.T) {
	//Creando una pila vacia para enteros.
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pilaEnteros.EstaVacia(), "Deberia devolver True, ya que la pila deberia estar vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaEnteros.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
	//Creando una pila vacia para cadenas.
	pilaCadena := TDAPila.CrearPilaDinamica[string]()
	require.True(t, pilaCadena.EstaVacia(), "Deberia devolver True, ya que la pila deberia estar vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCadena.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pilaCadena.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
}

// TestApilarElementos se encarga de verificar que se puedan apilar elementos y que al desapilarlos se mantenga el invariante de pila. Verifica que salgan el orden deseado.
func TestApilarElementos(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.Equal(t, 1, pila.VerTope(), "El tope de la pila deberia ser 1")
	require.Equal(t, 1, pila.Desapilar(), "Al desapilar deberia devolver el 1")
	require.True(t, pila.EstaVacia(), "Deberia devolver True, ya que la pila deberia estar vacia")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	pila.Apilar(5)
	pila.Apilar(6)
	require.Equal(t, 6, pila.VerTope(), "El tope de la pila deberia ser 6")
	require.Equal(t, 6, pila.Desapilar(), "Al desapilar deberia devolver el 6")
	pila.Apilar(7)
	pila.Apilar(8)
	require.Equal(t, 8, pila.VerTope(), "El tope de la pila deberia ser 8")
	require.Equal(t, 8, pila.Desapilar(), "Al desapilar deberia devolver el 8")
	require.Equal(t, 7, pila.Desapilar(), "Al desapilar deberia devolver el 7")
	require.Equal(t, 5, pila.Desapilar(), "Al desapilar deberia devolver el 5")
	require.Equal(t, 4, pila.VerTope(), "El tope de la pila deberia ser 4")
	require.False(t, pila.EstaVacia(), "Deberia devolver False, ya que la pila NO deberia estar vacia")
	require.Equal(t, 4, pila.Desapilar(), "Al desapilar deberia devolver el 4")
	require.Equal(t, 3, pila.Desapilar(), "Al desapilar deberia devolver el 3")
	require.Equal(t, 2, pila.Desapilar(), "Al desapilar deberia devolver el 2")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.True(t, pila.EstaVacia(), "Deberia devolver True, ya que la pila deberia estar vacia")
}

// TestVolumen se encarga de verificar que se puedan apilar y desapilar MUCHOS elementos hasta que esté vacía, comprobando que siempre cumpla el invariante.
func TestVolumen(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 1; i <= 10000; i++ {
		pila.Apilar(i)
		require.Equal(t, i, pila.VerTope(), "El tope de la pila deberia ser %d ", i)
	}
	for i := 10000; i != 1; i-- {
		require.Equal(t, i, pila.Desapilar(), "Al desapilar deberia devolver el %d", i)
		require.Equal(t, i-1, pila.VerTope(), "El tope de la pila deberia ser %d ", i-1)
	}
	pila.Desapilar()
	require.True(t, pila.EstaVacia(), "Deberia devolver True, ya que la pila deberia estar vacia")
}

// TestDesapilar se encarga de desapilar hasta que la pila este vacia para luego comprobar que se comporte como recien creada.
func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	//Apilamos y desapilamos hasta que la pila quede vacia
	for i := 0; i <= 4; i++ {
		pila.Apilar(i)
	}
	for i := 6; i != 1; i-- {
		pila.Desapilar()
	}
	//Probamos que se comporte como una pila recien creada
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
}

// TestFuncionesInvalidasPilaNueva se encarga de verificar que las acciones de desapilar y ver_tope en una pila recién creada sean inválidas.
func TestFuncionesInvalidasPilaNueva(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Error: No se deberia poder ver el tope de una pila recien creada")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Error: No se deberia poder desapilar una pila recien creada")
}

// TestEstaVacia verifica que la acción de esta_vacía en una pila recién creada sea verdadero.
func TestEstaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia(), "Una pila recien creada deberia estar vacia")
}

// TestFuncionesInvalidasPilaVacia se encarga de verificar que las acciones de desapilar y ver_tope en una pila a la que se le apiló y desapiló hasta estar vacía sean inválidas.
func TestFuncionesInvalidasPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	//Apilamos y desapilamos hasta que la pila quede vacia
	for i := 0; i <= 4; i++ {
		pila.Apilar(i)
	}
	for i := 6; i != 1; i-- {
		pila.Desapilar()
	}
	//Verificamos que desapilar y ver tope son acciones invalidas
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() }, "Error: deberia devolver ‘La pila esta vacia‘")
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() }, "Error: deberia devolver ‘La pila esta vacia‘")
}

// TestApilarDistintosDatos prueba apilar diferentes tipos de datos.
func TestApilarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaEnteros.Apilar(1)
	require.Equal(t, 1, pilaEnteros.VerTope(), "Deberia devolver 1")
	pilaEnteros.Apilar(2)
	pilaEnteros.Apilar(3)
	pilaEnteros.Apilar(4)
	require.Equal(t, 4, pilaEnteros.VerTope(), "Deberia devolver 4")

	//Con cadenas
	pilaCadena := TDAPila.CrearPilaDinamica[string]()
	pilaCadena.Apilar("A")
	require.Equal(t, "A", pilaCadena.VerTope(), "Deberia devolver A")
	pilaCadena.Apilar("B")
	pilaCadena.Apilar("C")
	pilaCadena.Apilar("D")
	require.Equal(t, "D", pilaCadena.VerTope(), "Deberia devolver D")

	//Con numeros con coma
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(4.1)
	require.Equal(t, 4.1, pilaFloat.VerTope(), "Deberia devolver 4.1")
	pilaFloat.Apilar(4.2)
	pilaFloat.Apilar(4.3)
	pilaFloat.Apilar(4.4)
	require.Equal(t, 4.4, pilaFloat.VerTope(), "Deberia devolver 4.4 ")

}

// TestDesapilarDistintosDatos prueba desapilar diferentes tipos de datos.
func TestDesapilarDistintosDatos(t *testing.T) {
	//Con numeros enteros
	pilaEnteros := TDAPila.CrearPilaDinamica[int]()
	pilaEnteros.Apilar(1)
	require.Equal(t, 1, pilaEnteros.Desapilar(), "Al desapilar deberia devolver el 1")
	pilaEnteros.Apilar(2)
	pilaEnteros.Apilar(3)
	pilaEnteros.Apilar(4)
	require.Equal(t, 4, pilaEnteros.Desapilar(), "Al desapilar deberia devolver el 4")
	require.Equal(t, 3, pilaEnteros.Desapilar(), "Al desapilar deberia devolver el 3")

	//Con cadenas
	pilaCadena := TDAPila.CrearPilaDinamica[string]()
	pilaCadena.Apilar("A")
	require.Equal(t, "A", pilaCadena.Desapilar(), "Al desapilar deberia devolver el A")
	pilaCadena.Apilar("B")
	pilaCadena.Apilar("C")
	pilaCadena.Apilar("D")
	require.Equal(t, "D", pilaCadena.Desapilar(), "Al desapilar deberia devolver el D")
	require.Equal(t, "C", pilaCadena.Desapilar(), "Al desapilar deberia devolver el C")

	//Con numeros con coma
	pilaFloat := TDAPila.CrearPilaDinamica[float64]()
	pilaFloat.Apilar(4.1)
	require.Equal(t, 4.1, pilaFloat.Desapilar(), "Al desapilar deberia devolver el 4.1")
	pilaFloat.Apilar(4.2)
	pilaFloat.Apilar(4.3)
	pilaFloat.Apilar(4.4)
	require.Equal(t, 4.4, pilaFloat.Desapilar(), "Al desapilar deberia devolver el 4.4")
	require.Equal(t, 4.3, pilaFloat.Desapilar(), "Al desapilar deberia devolver el 4.3")
}
