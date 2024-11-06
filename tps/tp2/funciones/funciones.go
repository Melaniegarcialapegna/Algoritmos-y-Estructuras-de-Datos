package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

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

func AgregarArchivo(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], sitios TDADiccionario.Diccionario[string, int], rutaArchivo string) {
	// abrimos el archivo -> O(n)
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		panic("No se pudo leer el archivo")
	}
	scanner := bufio.NewScanner(archivo)
	diccArchivo := TDADiccionario.CrearHash[IP, []DatoLog]()
	for scanner.Scan() {
		datoLog := parsearLog(scanner.Text())
		if !diccArchivo.Pertenece(datoLog.ip) {
			arreglo := []DatoLog{datoLog}
			diccArchivo.Guardar(datoLog.ip, arreglo)
		} else {
			arreglo := diccArchivo.Obtener(datoLog.ip)
			arreglo = append(arreglo, datoLog)
			diccArchivo.Guardar(datoLog.ip, arreglo)
		}

		if !diccOrdenado.Pertenece(datoLog.ip) {
			arreglo := []DatoLog{datoLog}
			diccOrdenado.Guardar(datoLog.ip, arreglo)
		} else {
			arreglo := diccOrdenado.Obtener(datoLog.ip)
			arreglo = append(arreglo, datoLog)
			diccArchivo.Guardar(datoLog.ip, arreglo)
		}
		if !sitios.Pertenece(datoLog.url) {
			sitios.Guardar(datoLog.url, 1)
		} else {
			visitas := sitios.Obtener(datoLog.url)
			visitas++
			sitios.Guardar(datoLog.url, visitas)
		}
	}

	for iter := diccArchivo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, logs := iter.VerActual()
		contadorLogs := 0
		inicio := logs[0].fecha
		for _, log := range logs {
			diferencia := log.fecha.Sub(inicio) * 1000000000
			if diferencia < 2 {
				contadorLogs++
			}
			if contadorLogs >= 5 {
				fmt.Printf("DoS: %s\n", ip)
				break
			}
		}
	}

}

func VerVisitantes(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], desde IP, hasta IP) {
	fmt.Printf("Visitantes:")
	for iter := diccOrdenado.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		ip, _ := iter.VerActual()
		fmt.Printf("\t %s \n", ip)
	}
	fmt.Printf("OK")
}

func VerMasVisitados(diccionario TDADiccionario.Diccionario[string, int], cantidad int) {
	sitios := []Sitio{}
	for iter := diccionario.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		url, visitas := iter.VerActual()
		sitio := Sitio{url, visitas}
		sitios = append(sitios, sitio)
	}
	heap := TDAHeap.CrearHeapArr(sitios, compararSitios)
	fmt.Printf("OK\n")
	fmt.Printf("Sitios más visitados:\n")
	for i := 0; i < cantidad; i++ {
		sitio := heap.Desencolar()
		fmt.Printf("\t%s - %d\n", sitio.url, sitio.visitas)
	}
	fmt.Printf("OK\n")
}

func parsearLog(linea string) DatoLog {
	campos := strings.Split(linea, "\t")
	fecha, _ := time.Parse("2006-01-02T15:04:05-07:00", campos[1])

	return DatoLog{
		ip:         IP(campos[0]),
		fecha:      fecha,
		metodoHttp: campos[2],
		url:        campos[3],
	}
}

func compararSitios(s1, s2 Sitio) int {
	if s1.visitas > s2.visitas {
		return 1
	} else if s1.visitas < s2.visitas {
		return -1
	}
	return 0
}
