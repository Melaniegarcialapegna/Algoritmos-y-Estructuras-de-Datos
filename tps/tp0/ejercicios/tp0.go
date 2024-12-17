package ejercicios

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	mayor := vector[0]
	posicion := 0
	for i := 1; i < len(vector); i++ {
		if vector[i] > mayor {
			mayor = vector[i]
			posicion = i
		}
	}
	return posicion
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	lenv1 := len(vector1)
	lenv2 := len(vector2)
	chico := menorLongitud(lenv1, lenv2)
	for i := 0; chico > i; i++ {
		if vector1[i] == vector2[i] {
			continue
		} else if vector1[i] > vector2[i] {
			return 1
		} else if vector2[i] > vector1[i] {
			return -1
		}
	}
	if lenv1 == lenv2 {
		return 0
	} else if lenv1 > lenv2 {
		return 1
	} else {
		return -1
	}
}

// Devuelve la cantidad de elementos del vector con menor longitud
func menorLongitud(lenv1 int, lenv2 int) int {
	if lenv1 > lenv2 {
		return lenv2
	}
	return lenv1
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	lenv := len(vector)
	for i := lenv - 1; i > 0; i-- {
		noOrdenado := vector[:i+1]
		posicionMaximo := Maximo(noOrdenado)
		if posicionMaximo != -1 {
			Swap(&vector[posicionMaximo], &vector[i])
		}
	}

}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	if len(vector) == 0 {
		return 0
	}
	return vector[0] + Suma(vector[1:])
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) <= 1 {
		return true
	} else if cadena[0] != cadena[len(cadena)-1] {
		return false
	}
	return EsCadenaCapicua(cadena[1 : len(cadena)-1])
}
