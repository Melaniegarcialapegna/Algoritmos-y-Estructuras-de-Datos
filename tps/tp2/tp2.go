package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	Funciones "tp2/funciones"
	//TDAColaPrioridad "tps/tdas/cola_prioridad"
)

func cmpIps(ip1, ip2 Funciones.IP) int {
	ip1Separada := strings.Split(string(ip1), ".") //separar en campos (buscar)
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
	//file, _ := os.Open("04_in")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		entradas := strings.Split(linea, " ")

		switch entradas[0] {
		case "agregar_archivo":
			Funciones.AgregarArchivo(diccionarioAbbIps, sitios, entradas[1])
		case "ver_visitantes":
			Funciones.VerVisitantes(diccionarioAbbIps, Funciones.IP(entradas[1]), Funciones.IP(entradas[2]))
		case "ver_mas_visitados":
			cantidad, _ := strconv.Atoi(entradas[1])
			Funciones.VerMasVisitados(sitios, cantidad)
		}
	}
}
