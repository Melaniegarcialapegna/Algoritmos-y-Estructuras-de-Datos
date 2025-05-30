package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	Funciones "tp2/funciones"
)

const (
	MSJ_ERROR = "Error en comando %s"
)

// cmpIps compara dos direcciones IP
func cmpIps(ip1, ip2 Funciones.IP) int {
	ip1Separada := strings.Split(string(ip1), ".")
	ip2Separada := strings.Split(string(ip2), ".")
	for i := 0; i < 4; i++ {
		num1, _ := strconv.Atoi(ip1Separada[i])
		num2, _ := strconv.Atoi(ip2Separada[i])
		if num1 == num2 {
			continue
		} else if num1 > num2 {
			return 1
		} else {
			return -1
		}
	}
	return 0
}

func main() {
	diccionarioAbbIps := TDADiccionario.CrearABB[Funciones.IP, []Funciones.DatoLog](cmpIps)
	sitios := TDADiccionario.CrearHash[string, int]()
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		linea := scanner.Text()
		entradas := strings.Split(linea, " ")
		huboError := false
		if huboError {
			break
		}

		switch entradas[0] {

		case "agregar_archivo":
			err := Funciones.AgregarArchivo(diccionarioAbbIps, sitios, entradas[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, MSJ_ERROR, "agregar_archivo")
				fmt.Fprintf(os.Stderr, "\n")
				huboError = true
			}

		case "ver_visitantes":
			if len(entradas) != 3 {
				fmt.Fprintf(os.Stderr, MSJ_ERROR, "ver_visitantes")
				fmt.Fprintf(os.Stderr, "\n")
				huboError = true
				break
			}
			Funciones.VerVisitantes(diccionarioAbbIps, Funciones.IP(entradas[1]), Funciones.IP(entradas[2]))

		case "ver_mas_visitados":
			if len(entradas) != 2 {
				fmt.Fprintf(os.Stderr, MSJ_ERROR, "ver_mas_visitados")
				fmt.Fprintf(os.Stderr, "\n")
				huboError = true
				break
			}

			cantidad, _ := strconv.Atoi(entradas[1])
			Funciones.VerMasVisitados(sitios, cantidad)
		}
	}
}
