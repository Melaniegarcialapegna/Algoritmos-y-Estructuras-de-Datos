package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const CANTIDAD_VOLUMEN = 2418

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
	//Con un elemento
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 1, lista.Largo(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, lista.VerPrimero(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, lista.VerUltimo(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, lista.BorrarPrimero(), "ERROR: deberia devolver '1'")
	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, lista.Largo(), "ERROR: deberia devolver '0'")

	//Con varios elementos
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	require.False(t, lista.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 4, lista.Largo(), "ERROR: deberia devolver '4'")
	require.Equal(t, 2, lista.VerPrimero(), "ERROR: deberia devolver '2'")
	require.Equal(t, 5, lista.VerUltimo(), "ERROR: deberia devolver '5'")
	require.Equal(t, 2, lista.BorrarPrimero(), "ERROR: deberia devolver '2'")
	require.Equal(t, 3, lista.BorrarPrimero(), "ERROR: deberia devolver '3'")
	require.Equal(t, 4, lista.BorrarPrimero(), "ERROR: deberia devolver '4'")
	require.False(t, lista.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 5, lista.VerPrimero(), "ERROR: deberia devolver '5'")
	require.Equal(t, 5, lista.VerUltimo(), "ERROR: deberia devolver '5'")
	require.Equal(t, 5, lista.BorrarPrimero(), "ERROR: deberia devolver '5'")
	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, lista.Largo(), "ERROR: deberia devolver '0'")

}

// TestVolumen se encarga de verificar que se puedan insertar y eliminar MUCHOS elementos hasta que este vacia,comprobando que siempre cumpla el invariante
func TestVolumen(t *testing.T) {
	// Agregar varios elementos al principio y borrar del primero
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.Largo(), "Al agregar un nuevo elemento al principio debería cambiar el largo")
	}
	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.Largo(), "El largo de la lista debe estar actualizado en todo momento")
		require.Equal(t, i, lista.BorrarPrimero(), "BorrarPrimero devolver el elemento en la primera posicion de la lista")
		require.Equal(t, i-1, lista.Largo(), "El largo de la lista debe actualizarse al borrar un elemento")
	}
	require.True(t, lista.EstaVacia(), "Al eliminar todos los elementos, ")

	// Agregar varios elementos al final y borrar del primero
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.Largo(), "Al agregar un nuevo elemento debería cambiar el largo")
	}
	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.Largo(), "El largo de la lista debe mantenerse actualizado luego de insertar en el ultimo")
		require.Equal(t, CANTIDAD_VOLUMEN+1-i, lista.BorrarPrimero(), "BorrarPrimero devuelve el elemento correcto")
		require.Equal(t, i-1, lista.Largo(), "El largo de la lista se actualiza luego de borrar el ultimo")
	}
	require.True(t, lista.EstaVacia(), "Luego de agregar al final y borrar del primero la misma cantidad, la lista se debe vaciar")

	// Agregar varios elementos al principio y final y borrar del principio
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		lista.InsertarPrimero(i)
		require.Equal(t, i*2, lista.Largo(), "Al agregar dos nuevos elemento, el largo deberia ser %d", i*2)
	}
	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.BorrarPrimero(), "BorrarPrimero borra el primero luego de insertar intercaladamente primero y ultimo")
	}
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		require.Equal(t, i, lista.BorrarPrimero(), "BorrarPrimero borra el primero luego de insertar intercaladamente primero y ultimo")
	}
	require.True(t, lista.EstaVacia(), "Luego de agregar al final y borrar del primero la misma cantidad, la lista se debe vaciar")
}

// TestFuncionesInvalidasListaNueva se encarga de verificar que las acciones de BorrarPrimero, VerPrimero y VerUltimo en una lista recien creadas sean invalidas.
func TestFuncionesInvalidasListaNueva(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
}

// TestEstaVacia verifica que la accion de EstaVacia en una lista recien creada sea verdadero
func TestEstaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
}

// TestFuncionesInvalidasListaVacia se encarga de verificar que las acciones de BorrarPrimero, VerPrimero y VerUltimo en una lista en la cual se agregaron elementos que posteriormente fueron eliminados hasta estar vacia sean invalidos
func TestFuncionesInvalidasListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	// Agregar elementos y eliminarlos hasta quela lista quede vacia
	for i := 1; i <= 4; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.Largo(), "Al agregar un nuevo elemento, el largo deberia ser %d", i)
	}
	for i := 2418; i >= 1; i-- {
		require.Equal(t, i, lista.Largo(), "El largo de la lista debe mantenerse actualizado luego de insertar en el ultimo")
		require.Equal(t, 2419-i, lista.BorrarPrimero(), "BorrarPrimero devuelve el elemento correcto")
		require.Equal(t, i-1, lista.Largo(), "El largo de la lista se actualiza luego de borrar el ultimo")
	}
	require.True(t, lista.EstaVacia(), "Luego de agregar al final y borrar del primero la misma cantidad, la lista se debe vaciar")
}

/*
















asdasd



asdasd


asd





asd






*/
