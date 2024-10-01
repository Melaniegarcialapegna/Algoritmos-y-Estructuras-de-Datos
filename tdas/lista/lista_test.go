package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

const CANTIDAD_VOLUMEN = 2418

/*------------------ TEST PARA LISTA -------------------------*/

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
		require.Equal(t, i, lista.Largo(), "ERROR: Al agregar un nuevo elemento el largo deberia ser %d", i)
	}

	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.Largo(), "ERROR: El largo de la lista debe ser %d", i)
		require.Equal(t, i, lista.BorrarPrimero(), "ERROR: BorrarPrimero deberia devolver el numero %d", i)
		require.Equal(t, i-1, lista.Largo(), "ERROR: El largo de la lista deberia ser %d", i-1)
	}

	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")

	// Agregar varios elementos al final y borrar del primero
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		require.Equal(t, i, lista.Largo(), "ERROR: El largo de la lista deberia ser %d", i)
	}
	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.Largo(), "ERROR: El largo de la lista debe ser %d", i)
		require.Equal(t, CANTIDAD_VOLUMEN+1-i, lista.BorrarPrimero(), "ERROR: Al borrar el primer elemento deberia devolver el %d", CANTIDAD_VOLUMEN+1-i)
		require.Equal(t, i-1, lista.Largo(), "ERROR: El largo de la lista deberia ser %d", i-1)
	}
	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")

	// Agregar varios elementos al principio y final y borrar del principio
	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		lista.InsertarUltimo(i)
		lista.InsertarPrimero(i)
		require.Equal(t, i*2, lista.Largo(), "ERROR: El largo deberia ser %d", i*2)
	}

	for i := CANTIDAD_VOLUMEN; i >= 1; i-- {
		require.Equal(t, i, lista.BorrarPrimero(), "ERROR: Deberia devolver %d", i)
	}

	for i := 1; i <= CANTIDAD_VOLUMEN; i++ {
		require.Equal(t, i, lista.BorrarPrimero(), "ERROR: Deberia devolver %d", i)
	}

	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
}

// TestFuncionesInvalidasListaNueva se encarga de verificar que las acciones de BorrarPrimero, VerPrimero y VerUltimo en una lista recien creadas sean invalidas.
func TestFuncionesInvalidasListaNueva(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
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
	// Agregar elementos y eliminarlos hasta que la lista quede vacia
	for i := 1; i <= 4; i++ {
		lista.InsertarPrimero(i)
		require.Equal(t, i, lista.Largo(), "ERROR: Al agregar un nuevo elemento, el largo deberia ser %d", i)
	}
	for i := 4; i >= 1; i-- {
		require.Equal(t, i, lista.BorrarPrimero(), "ERROR : Deberia borrar y devolver %d", i)
	}

	require.True(t, lista.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.Equal(t, 0, lista.Largo(), "ERROR: El largo de una lista vaciada deberia ser 0")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia' despues de vaciar la lista")
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia' despues de vaciar la lista")
}

// TestDistintosTiposDatos prueba insertar y eliminar elementos en listas con distintos tipos de datos
func TestDistintosTiposDatos(t *testing.T) {
	//Con enteros
	listaEnteros := TDALista.CrearListaEnlazada[int]()
	listaEnteros.InsertarPrimero(1)
	require.False(t, listaEnteros.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 1, listaEnteros.Largo(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, listaEnteros.VerPrimero(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, listaEnteros.VerUltimo(), "ERROR: deberia devolver '1'")
	require.Equal(t, 1, listaEnteros.BorrarPrimero(), "ERROR: deberia devolver '1'")
	require.True(t, listaEnteros.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaEnteros.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, listaEnteros.Largo(), "ERROR: deberia devolver '0'")

	//Con cadenas
	listaCadena := TDALista.CrearListaEnlazada[string]()
	listaCadena.InsertarPrimero("m")
	require.False(t, listaCadena.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 1, listaCadena.Largo(), "ERROR: deberia devolver '1'")
	require.Equal(t, "m", listaCadena.VerPrimero(), "ERROR: deberia devolver 'm'")
	require.Equal(t, "m", listaCadena.VerUltimo(), "ERROR: deberia devolver 'm'")
	require.Equal(t, "m", listaCadena.BorrarPrimero(), "ERROR: deberia devolver 'm'")
	require.True(t, listaCadena.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaCadena.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaCadena.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, listaCadena.Largo(), "ERROR: deberia devolver '0'")

	//Con floats
	listaFloats := TDALista.CrearListaEnlazada[float64]()
	listaFloats.InsertarPrimero(float64(4.1))
	require.False(t, listaFloats.EstaVacia(), "Deberia devolver False, ya que la lista NO deberia estar vacia")
	require.Equal(t, 1, listaFloats.Largo(), "ERROR: deberia devolver '1'")
	require.Equal(t, float64(4.1), listaFloats.VerPrimero(), "ERROR: deberia devolver '4.1'")
	require.Equal(t, float64(4.1), listaFloats.VerUltimo(), "ERROR: deberia devolver '4.1'")
	require.Equal(t, float64(4.1), listaFloats.BorrarPrimero(), "ERROR: deberia devolver '4.1'")
	require.True(t, listaFloats.EstaVacia(), "Deberia devolver True, ya que la lista deberia estar vacia")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloats.VerPrimero() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.PanicsWithValue(t, "La lista esta vacia", func() { listaFloats.VerUltimo() }, "ERROR: deberia devolver 'La lista esta vacia'")
	require.Equal(t, 0, listaFloats.Largo(), "ERROR: deberia devolver '0'")
}

/*------------------ TEST PARA ITERADOR EXTERNO LISTA -------------------------*/
/*
[X] 1-Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
[X] 2-Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
[X] 3-Insertar un elemento en el medio se hace en la posición correcta.
[X] 4-Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
[X] 5-Remover el último elemento con el iterador cambia el último de la lista.
[X] 6-Verificar que al remover un elemento del medio, este no está.
[ ] 7-Otros casos borde que pueden encontrarse al utilizar el iterador externo.
[ ] 8-Casos del iterador interno, incluyendo casos con corte (la función visitar devuelve false eventualmente).
*/

// TestIteradorExternoInsertarPrincipio corrobora que al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
func TestIteradorExternoInsertarPrincipio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.False(t, iterador.HaySiguiente())
	elemento := 24
	iterador.Insertar(elemento)
	require.Equal(t, elemento, lista.VerPrimero(), "ERROR : deberia devolver %d", 24)
}

// TestIteradorExternoInsertarAlFinal corrobora que insertar en un iterador que esta exactamente al final es lo mismo que agregar al final de la lista
func TestIteradorExternoInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()

	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(4) //🤓☝
	require.Equal(t, lista.VerUltimo(), 4, "ERROR: Deberia devolver %d", 4)

	//Vemos que se comporte igual que al insertar al final
	ultimoConIterador := lista.VerUltimo()
	lista2 := TDALista.CrearListaEnlazada[int]()
	lista2.InsertarUltimo(4)
	require.Equal(t, ultimoConIterador, lista2.VerUltimo(), "ERROR: El ultimo de la lista deberia ser %d", ultimoConIterador)
}

// TestIteradorExternoInsertarMedio verifica al insertar un elemento en el medio se hace en la posición correcta
func TestIteradorExternoInsertarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(6)
	iterador := lista.Iterador()
	for i := 0; i < 2; i++ {
		iterador.Siguiente()
	}
	elemento := 24
	iterador.Insertar(elemento)
	require.Equal(t, 24, iterador.VerActual(), "ERROR: Deberia devolver %d", 24)
}

// TestIteradorExternoBorrarPrimero corrobora que al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista
func TestIteradorExternoBorrarPrimero(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(4)
	iterador := lista.Iterador()
	elementoBorrado := iterador.Borrar()
	require.Equal(t, elementoBorrado, lista.VerPrimero(), "ERROR: Deberia devolver %d", elementoBorrado)
}

func TestIteradorExternoBorrarUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	iterador := lista.Iterador()

	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}

	iterador.Insertar(5)
	require.Equal(t, 5, lista.VerUltimo(), "ERROR: Deberia devolver %d", 5)
}

// TestIteradorExternoBorrarMedio verifica que al remover un elemento del medio, este no está
func TestIteradorExternoBorrarMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	lista.InsertarPrimero(2)
	lista.InsertarPrimero(3)
	lista.InsertarPrimero(4)
	lista.InsertarPrimero(5)
	iterador := lista.Iterador()
	for i := 0; i < 3; i++ {
		iterador.Siguiente()
	}

	require.Equal(t, 5, lista.Largo(), "ERROR: Deberia devolver %d", 5)
	require.Equal(t, 2, iterador.VerActual(), "ERROR: Deberia devolver %d", 2)
	//Se elimina el elemento del medio y se verifica que ya no este
	require.Equal(t, 2, iterador.Borrar(), "ERROR: Deberia devolver %d", 2)
	require.Equal(t, 4, lista.Largo(), "ERROR: Deberia devolver %d", 4)
	require.Equal(t, 1, iterador.VerActual(), "ERROR: Deberia devolver %d", 1)
}

func TestIteradorListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador ya termino de iterar", func() { iterador.Borrar() }, "ERROR: Deberia lanzar panic'")
	require.PanicsWithValue(t, "El iterador ya termino de iterar", func() { iterador.Siguiente() }, "ERROR: Deberia lanzar panic'")
	require.PanicsWithValue(t, "El iterador ya termino de iterar", func() { iterador.VerActual() }, "ERROR: Deberia lanzar panic'")
}
