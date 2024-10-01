package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

/*
[] 1-Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
[] 2-Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
[] 3-Insertar un elemento en el medio se hace en la posición correcta.
[] 4-Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
[] 5-Remover el último elemento con el iterador cambia el último de la lista.
[] 6-Verificar que al remover un elemento del medio, este no está.
[] 7-Otros casos borde que pueden encontrarse al utilizar el iterador externo.
[] 8-Casos del iterador interno, incluyendo casos con corte (la función visitar devuelve false eventualmente).
*/

// TestListaVacia se encarga de verificar que se pueda crear una Lista vacia y que se comporte como tal
func TestListaVacia(t *testing.T) {
	//Creando una lista vacia para enteros
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	require.True(t, listaEnteros.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, listaEnteros.Largo(), "ERROR: deberia devolver '0'")
	//Creando una lista vacia para cadenas
	listaCadena := TDALista.CrearListaEnlazada[string]()
	require.True(t, listaCadena.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaCadena.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaCadena.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, listaCadena.Largo(), "ERROR: deberia devolver '0'")
}

// TestInsertarEliminar se encarga de verificar que se puedan insertar elementos y que al eliminarlos se mantenga el invariante de lista.
func TestInsertarEliminar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
}

// TestVolumen se encarga de verificar que se puedan insertar y eliminar MUCHOS elementos hasta que este vacia,comprobando que siempre cumpla el invariante
func TestVolumen(t *testing.T) {
	// Agregar varios elementos y borrar del ultimo
	for i := 0; i >= 0; i-- {

	}
	// Agregar varios elementos y borrar del primero
	// Agregar varios elementos y borrar de la mitad
	// Agregar varios elementos y borrar intercaladamente
}
