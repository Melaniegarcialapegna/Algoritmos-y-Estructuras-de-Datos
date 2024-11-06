package tp2

import (
	"bufio"
	"os"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
	//TDAColaPrioridad "tps/tdas/cola_prioridad"
)

// import(
// 	"os"
// 	"fmt"
// 	"time"
// )

type IP string
type DatoLog struct {
	ip         IP
	fecha      time.Time
	metodoHttp string
	url        string
}

type Sitio struct {
	url     string
	visitas int
}

func cmpIps(ip1, ip2 IP) int {
	ip1Separada := strings.Fields(".") //separar en campos (buscar)
	ip2Separada := strings.Fields(".")
	for i := 0; i < 4; i++ {
		if ip1Separada[i] == ip2Separada[i] {
			continue
		}
		if ip1Separada[i] > ip2Separada[i] {
			return 1
		}
		return -1
	}
	return 0
}

func main() {
	diccionarioAbbIps := TDADiccionario.CrearABB[IP, []DatoLog](cmpIps)
	sitios := TDADiccionario.CrearHash[Sitio, int]()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		entradas := strings.Split(linea, "")

	}

	switch entradas[0] {
	case "agregar_archivo":
		AgregarArchivo(diccionarioAbbIps, sitios, entradas[1])
	case "ver_visitantes":
		VerVisitantes(diccionarioAbbIps, desde, hasta)
	case "ver_mas_visitados":
		VerMasVisitados(sitios)
	}
}
