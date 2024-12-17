package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"tp0/ejercicios"
)

const ruta1 = "archivo1.in"
const ruta2 = "archivo2.in"

func main() {
	arreglo1 := cargarArreglos(ruta1)
	arreglo2 := cargarArreglos(ruta2)

	mayor := ejercicios.Comparar(arreglo1, arreglo2)
	if mayor == 1 {
		imprimirArreglo(arreglo1)
	} else {
		imprimirArreglo(arreglo2)

	}

}

// Abre, lee el archivo y carga en memoria los arreglos correspondientes
func cargarArreglos(ruta string) []int {
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Printf("Error %v al abrir el archivo %s\n", ruta, err)
	}
	defer archivo.Close()

	var slice []int

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		entero, err := strconv.Atoi(s.Text())
		if err != nil {
			fmt.Printf("No se pudo convertir: %v\n", err)
		}
		slice = append(slice, entero)
	}
	return slice
}

// Ordena el arreglo y lo imprime
func imprimirArreglo(arreglo []int) {
	ejercicios.Seleccion(arreglo)
	for _, valor := range arreglo {
		fmt.Println(valor)
	}
}
