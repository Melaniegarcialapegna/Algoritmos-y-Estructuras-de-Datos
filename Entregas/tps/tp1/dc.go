package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	Calculadora "tp1/operaciones"
)

func main() {
	ejecutar(os.Stdin)
}

// ejecutar se encarga de manejar la entrada y procesar las operaciones
func ejecutar(scanner *os.File) {
	s := bufio.NewScanner(scanner)
	for s.Scan() {
		linea := s.Text()
		tokens := limpiarTokens(linea)
		resultado, err := Calculadora.CalcularOperacion(tokens)
		if err != nil {
			fmt.Println("ERROR")
		} else {
			fmt.Println(resultado)
		}
	}
}

// limpiarTokens se encarga de eliminar los espacios y devolver los tokens
func limpiarTokens(linea string) []string {
	tokensFiltrados := []string{}
	tokens := strings.Split(linea, " ") //Divide la linea en tokens utilizando un espacio como separador
	for _, token := range tokens {
		if token != "" { //Filtra los tokens que estan vacios
			tokensFiltrados = append(tokensFiltrados, token)
		}
	}
	return tokensFiltrados
}
