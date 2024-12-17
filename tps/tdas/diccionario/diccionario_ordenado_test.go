package diccionario_test

import (
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"math/rand"

	"github.com/stretchr/testify/require"
)

const (
	MAX_VALOR_RANDOM = 2
	N                = 8000
)

var DESDE_RANGOS int = 18
var HASTA_RANGOS int = 82
var desdeNuevo int = 7
var hastaNuevo int = 30
var arbolEspecifico []int = []int{24, 18, 22, 20, 10, 26, 28}

func comparacionEnteros(a, b int) int {
	return a - b
}

func TestDiccionarioVacioAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("A") })
}

func TestDiccionarioClaveDefaultAbb(t *testing.T) {
	t.Log("Prueba sobre un Abb vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestUnElementAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar("A", 10)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece("A"))
	require.False(t, abb.Pertenece("B"))
	require.EqualValues(t, 10, abb.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("B") })
	require.EqualValues(t, 1, abb.Cantidad(), "La cantidad de elementos es incorrecta")
}

func TestDiccionarioGuardarAbb(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))

	require.False(t, abb.Pertenece(claves[1]))
	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[1], valores[1])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))

	require.False(t, abb.Pertenece(claves[2]))
	abb.Guardar(claves[2], valores[2])
	require.True(t, abb.Pertenece(claves[0]))
	require.True(t, abb.Pertenece(claves[1]))
	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, 3, abb.Cantidad())
	require.EqualValues(t, valores[0], abb.Obtener(claves[0]))
	require.EqualValues(t, valores[1], abb.Obtener(claves[1]))
	require.EqualValues(t, valores[2], abb.Obtener(claves[2]))
}

func TestReemplazoDatoAbb(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(clave, "miau")
	abb.Guardar(clave2, "guau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, "miau", abb.Obtener(clave))
	require.EqualValues(t, "guau", abb.Obtener(clave2))
	require.EqualValues(t, 2, abb.Cantidad())

	abb.Guardar(clave, "miu")
	abb.Guardar(clave2, "baubau")
	require.True(t, abb.Pertenece(clave))
	require.True(t, abb.Pertenece(clave2))
	require.EqualValues(t, 2, abb.Cantidad())
	require.EqualValues(t, "miu", abb.Obtener(clave))
	require.EqualValues(t, "baubau", abb.Obtener(clave2))
}

func TestReemplazoDatoHopscotchAbb(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	for i := 0; i < 500; i++ {
		abb.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		abb.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = abb.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioBorrarAbb(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)

	require.False(t, abb.Pertenece(claves[0]))
	require.False(t, abb.Pertenece(claves[0]))
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])

	require.True(t, abb.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], abb.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[2]) })
	require.EqualValues(t, 2, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[2]))

	require.True(t, abb.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], abb.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[0]) })
	require.EqualValues(t, 1, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[0]) })

	require.True(t, abb.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], abb.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(claves[1]) })
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(claves[1]) })
}

func TestConClavesNumericasAbb(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	abb := TDADiccionario.CrearABB[int, string](comparacionEnteros)
	clave := 10
	valor := "Gatito"

	abb.Guardar(clave, valor)
	require.EqualValues(t, 1, abb.Cantidad())
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, valor, abb.Obtener(clave))
	require.EqualValues(t, valor, abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestClaveVaciaAbb(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	clave := ""
	abb.Guardar(clave, clave)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, clave, abb.Obtener(clave))
}

func TestValorNuloAbb(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	clave := "Pez"
	abb.Guardar(clave, nil)
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, (*int)(nil), abb.Obtener(clave))
	require.EqualValues(t, (*int)(nil), abb.Borrar(clave))
	require.False(t, abb.Pertenece(clave))
}

func TestGuardarYBorrarRepetidasVecesAbb(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces")

	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	for i := 0; i < 1000; i++ {
		abb.Guardar(i, i)
		require.True(t, abb.Pertenece(i))
		require.EqualValues(t, 1, abb.Cantidad())
		abb.Borrar(i)
		require.False(t, abb.Pertenece(i))
		require.EqualValues(t, 0, abb.Cantidad())
	}

}

// TEST DE ITERADORES
func TestIteradorInternoClavesAbb(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	abb := TDADiccionario.CrearABB[string, *int](strings.Compare)
	abb.Guardar(claves[0], nil)
	abb.Guardar(claves[1], nil)
	abb.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	abb.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarAbb(cs[0], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[1], claves))
	require.NotEqualValues(t, -1, buscarAbb(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func buscarAbb(clave string, claves []string) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}

func TestIteradorInternoValoresAbb(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestIteradorInternoValoresConBorradosAbb(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave0 := "Elefante" //
	clave1 := "Gato"     //V
	clave2 := "Perro"    //
	clave3 := "Vaca"     //
	clave4 := "Burrito"  //
	clave5 := "Hamster"  //

	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	abb.Guardar(clave0, 7)
	abb.Guardar(clave1, 6)
	abb.Guardar(clave2, 2)
	abb.Guardar(clave3, 3)
	abb.Guardar(clave4, 4)
	abb.Guardar(clave5, 5)

	abb.Borrar(clave0)

	factorial := 1
	ptrFactorial := &factorial
	abb.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestEjecutarPruebaVolumenAbb(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	claves := make([]int, N)
	valores := make([]int, N)

	/* Inserta 'n' parejas en el abb */
	for i := 0; i < N; i++ {
		valores[i] = i
		claves[i] = i
		abb.Guardar(claves[i], valores[i])
	}

	require.EqualValues(t, N, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < N; i++ {
		ok = abb.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abb.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(t, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(t, N, abb.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < N; i++ {
		ok = abb.Borrar(claves[i]) == valores[i]
		require.Equal(t, N-1-i, abb.Cantidad(), "La cantidad de elementos es incorrecta")
		if !ok {
			break
		}
		ok = !abb.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(t, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(t, 0, abb.Cantidad())
}

func TestIterarDiccionarioVacioAbb(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	abb := TDADiccionario.CrearABB[string, int](strings.Compare)
	iter := abb.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestDiccionarioIterarAbb(t *testing.T) {
	t.Log("Guardamos 3 valores en un Diccionario, e iteramos validando que las claves sean todas diferentes " +
		"pero pertenecientes al diccionario. Además los valores de VerActual y Siguiente van siendo correctos entre sí")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	abb.Guardar(claves[0], valores[0])
	abb.Guardar(claves[1], valores[1])
	abb.Guardar(claves[2], valores[2])
	iter := abb.Iterador()

	require.True(t, iter.HaySiguiente())
	primero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(primero, claves))

	iter.Siguiente()
	segundo, segundo_valor := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(segundo, claves))
	require.True(t, segundo > primero, "El iterador no itero en el orden correcto")
	require.EqualValues(t, valores[buscarAbb(segundo, claves)], segundo_valor)
	require.NotEqualValues(t, primero, segundo)
	require.True(t, iter.HaySiguiente())

	iter.Siguiente()
	require.True(t, iter.HaySiguiente())
	tercero, _ := iter.VerActual()
	require.NotEqualValues(t, -1, buscarAbb(tercero, claves))
	require.True(t, tercero > segundo, "El iterador no itero en el orden correcto")
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, segundo, tercero)
	iter.Siguiente()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorNoLlegaAlFinalAbb(t *testing.T) {
	t.Log("Crea un iterador y no lo avanza. Luego crea otro iterador y lo avanza.")
	abb := TDADiccionario.CrearABB[string, string](strings.Compare)
	claves := []string{"A", "B", "C"}
	abb.Guardar(claves[0], "")
	abb.Guardar(claves[1], "")
	abb.Guardar(claves[2], "")

	abb.Iterador()
	iter2 := abb.Iterador()
	iter2.Siguiente()
	iter3 := abb.Iterador()
	primero, _ := iter3.VerActual()
	iter3.Siguiente()
	segundo, _ := iter3.VerActual()
	iter3.Siguiente()
	tercero, _ := iter3.VerActual()
	iter3.Siguiente()
	require.False(t, iter3.HaySiguiente())
	require.NotEqualValues(t, primero, segundo)
	require.NotEqualValues(t, tercero, segundo)
	require.NotEqualValues(t, primero, tercero)
	require.NotEqualValues(t, -1, buscarAbb(primero, claves))
	require.NotEqualValues(t, -1, buscarAbb(segundo, claves))
	require.NotEqualValues(t, -1, buscarAbb(tercero, claves))
}

func TestEjecutarPruebasVolumenIteradorAbb(t *testing.T) {
	t.Log("Se inserta una gran cantidad de elementos desordenados en el ABB. " +
		"Luego se comprueba que los elementos se recorren de forma ordenada.")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	claves := make([]int, N)
	valores := make([]int, N)

	for i := 0; i < N; i++ {
		numeroRandom := int(rand.Float64() * MAX_VALOR_RANDOM)
		claves[i] = numeroRandom
		valores[i] = i
		abb.Guardar(claves[i], valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := abb.Iterador()
	require.True(t, iter.HaySiguiente())

	ok := true
	var i int
	var anteriorNumero int

	for i = 0; i < N; i++ {
		if !iter.HaySiguiente() {
			break
		}
		c1, _ := iter.VerActual()
		if i != 0 && anteriorNumero > c1 {
			ok = false
			break
		}
		anteriorNumero = c1
		iter.Siguiente()
	}
	require.True(t, ok, "Iteracion en volumen no funciona correctamente")
	require.False(t, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")
}

func TestVolumenIteradorCorteAbb(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	/* Inserta 'n' parejas en el abb */
	for i := 0; i < N; i++ {
		abb.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	abb.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func TestIteradorCorteAbb(t *testing.T) {
	t.Log("Verifica que no se hagan iteraciones de mas en distintos puntos de corte.")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 0; i < len(arbolEspecifico); i++ {
		abb.Guardar(arbolEspecifico[i], i+1)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	//cortar en Raiz
	abb.Iterar(func(clave int, dato int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
		}
		if clave == 24 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	//cortar en Hoja
	seguirEjecutando = true
	abb.Iterar(func(clave int, dato int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
		}
		if clave == 22 {
			seguirEjecutando = false
			return false
		}
		return true
	})
	//cortar en un hijo
	seguirEjecutando = true
	abb.Iterar(func(clave int, dato int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
		}
		if clave == 26 {
			seguirEjecutando = false
			return false
		}
		return true
	})
	//cortar en dos hijos
	seguirEjecutando = true
	abb.Iterar(func(clave int, dato int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
		}
		if clave == 18 {
			seguirEjecutando = false
			return false
		}
		return true
	})
	require.False(t, siguioEjecutandoCuandoNoDebia, "El iterador itero de mas")
}

// Suma
// con rangos
func TestIteradorInternoSuma(t *testing.T) {
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 0; i < len(arbolEspecifico); i++ {
		abb.Guardar(arbolEspecifico[i], i+1)
	}

	suma := 0
	abb.Iterar(func(clave int, dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 28, suma, "ERROR: Deberia devolver %d", 28)
}

func TestIteradorInternoRangosSuma(t *testing.T) {
	t.Log("Corrobora que el iterador interno funciona bien al hacer operaciones con los valores iterados")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 0; i < len(arbolEspecifico); i++ {
		abb.Guardar(arbolEspecifico[i], i+1)
	}

	suma := 0
	iteroFueraDeRango := false
	abb.IterarRango(&DESDE_RANGOS, &HASTA_RANGOS, func(clave int, dato int) bool {
		if clave < DESDE_RANGOS || clave > HASTA_RANGOS {
			iteroFueraDeRango = true
			return false
		}
		suma += dato
		return true
	})
	require.False(t, iteroFueraDeRango, "ERROR: Deberia devolver %d", 5050)
	require.Equal(t, 23, suma, "ERROR: Deberia devolver %d", 23)
}

func TestVolumenIteradorInternoRangos(t *testing.T) {
	t.Log("Corrobora que el iterador interno solo itere los numeros dentro del rango requerido.")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 1; i <= N; i++ {
		abb.Guardar(i, i)
	}
	iteroFueraDeRango := false
	abb.IterarRango(&DESDE_RANGOS, &HASTA_RANGOS, func(clave int, dato int) bool {
		if clave < DESDE_RANGOS || clave > HASTA_RANGOS {
			iteroFueraDeRango = true
			return false
		}
		return true
	})
	require.False(t, iteroFueraDeRango, "ERROR")
}

func TestVolumenIteradorExternoRangos(t *testing.T) {
	t.Log("Corrobora que el iterador externo solo itere los numeros dentro del rango requerido.")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 1; i <= N; i++ {
		abb.Guardar(i, i)
	}

	iteroFueraDeRango := false
	for iter := abb.IteradorRango(&DESDE_RANGOS, &HASTA_RANGOS); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		if clave < DESDE_RANGOS || clave > HASTA_RANGOS {
			iteroFueraDeRango = true
		}
	}
	require.False(t, iteroFueraDeRango, "ERROR")
}

func TestRamaIzquierdaSinHijosDerechos(t *testing.T) {
	t.Log("Corrobora que al recorrer por rangos un arbol que solo tiene una rama izquierda sin ningun hijo derecho" +
		"no haya ningun salto a un nodo izquierdo de un nodo derecho que no existe")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	abb.Guardar(20, 20)
	abb.Guardar(10, 20)
	abb.Guardar(5, 20)
	require.EqualValues(t, 3, abb.Cantidad())

	iteroFueraDeRango := false
	for iter := abb.IteradorRango(&desdeNuevo, &hastaNuevo); iter.HaySiguiente(); iter.Siguiente() {
		clave, _ := iter.VerActual()
		if clave < desdeNuevo || clave > hastaNuevo {
			iteroFueraDeRango = true
		}
	}
	require.False(t, iteroFueraDeRango, "ERROR")
}

func TestAgregarClaveBorrada(t *testing.T) {
	t.Log("Corrobora que al recorrer por rangos un arbol que solo tiene una rama izquierda sin ningun hijo derecho" +
		"no haya ningun salto a un nodo izquierdo de un nodo derecho que no existe")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	arbolRamaIzquierdaSinDerechos := []int{10, 5, 15, 12, 13, 17}

	for i, el := range arbolRamaIzquierdaSinDerechos {
		abb.Guardar(el, i)
	}

	require.True(t, abb.Pertenece(10))
	abb.Borrar(10)
	require.False(t, abb.Pertenece(10))
	abb.Guardar(10, 3)
	require.True(t, abb.Pertenece(10))
	require.EqualValues(t, 3, abb.Obtener(10))

}

func TestBorrarDosHijos(t *testing.T) {
	t.Log("Borra un nodo con dos hijos y corrobora que no pertenece. Luego lo vuelve a agregar y corrobora que si pertenece")
	abb := TDADiccionario.CrearABB[int, int](comparacionEnteros)

	for i := 0; i < 10000; i++ {
		abb.Guardar(i, i)
	}

	require.True(t, abb.Pertenece(20))
	require.Equal(t, 20, abb.Obtener(20))
	abb.Borrar(20)
	require.False(t, abb.Pertenece(20))
	abb.Guardar(20, 5)
	require.True(t, abb.Pertenece(20))
	require.Equal(t, 5, abb.Obtener(20))

	abb2 := TDADiccionario.CrearABB[int, int](comparacionEnteros)
	for i, el := range arbolEspecifico {
		abb2.Guardar(el, i)
	}

	require.True(t, abb2.Pertenece(18))
	require.Equal(t, 1, abb2.Obtener(18))
	abb2.Borrar(18)
	require.False(t, abb2.Pertenece(18))
	abb2.Guardar(18, 5)
	require.True(t, abb2.Pertenece(18))
	require.Equal(t, 5, abb2.Obtener(18))
}
