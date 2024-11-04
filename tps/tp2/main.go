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

type Sitio string
type IP string
type DatoLog struct {
	ip         IP
	fecha      time.Time
	metodoHttp string
	url        string
}

type SitioVisitantes struct {
	sitio   Sitio
	visitas int
}

func cmpIps(ip1, ip2 IP) int {
	//cosas
	// return ip1 - ip2
	return 0
}

func main() {
	diccionarioAbbIps := TDADiccionario.CrearABB[IP, []DatoLog](cmpIps)
	sitios := TDADiccionario.CrearHash[Sitio, int]()

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		linea := scanner.Text()
		entradas := strings.Fields(" ")

	}

	switch entradas[0] {
	case "agregar_archivo":
		AgregarArchivo(diccionarioAbbIps, sitios, entradas[1])
	case "ver_visitantes":
		VerVisitantes(diccionarioAbbIps)
	case "ver_mas_visitados":
		VerMasVisitados(sitios)
	}
}
