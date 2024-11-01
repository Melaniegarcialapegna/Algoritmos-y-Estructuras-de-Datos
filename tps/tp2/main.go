package tp2

import (
	"os"
	TDADiccionario "tdas/diccionario"
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
	fecha      string
	hora       string
	metodoHttp string
	metodoUrl  string
}

func cmpIps(ip1, ip2 IP) int {
	//cosas
	// return ip1 - ip2
	return 0
}

func main() {
	diccionarioAbbIps := TDADiccionario.CrearABB[IP, []DatoLog](cmpIps)
	arregloSitios := []Sitio{}

	switch os.Args[0] {
	case "agregar_archivo":
		AgregarArchivo(diccionarioAbbIps, arregloSitios)
	case "ver_visitantes":
		VerVisitantes(diccionarioAbbIps)
	case "ver_mas_visitados":
		VerMasVisitados(arregloSitios)
	}
}
