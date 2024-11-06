package tp2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

func AgregarArchivo(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], sitios TDADiccionario.Diccionario[string, int], rutaArchivo string) {
	// abrimos el archivo -> O(n)
	if len(os.Args) < 2 {
		panic("No se eviaron los argumentos suficientes")
	}
	archivo, err := os.Open(rutaArchivo)
	scanner := bufio.NewScanner(archivo)
	if err != nil {
		panic("No se pudo leer el archivo")
	}
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
	return

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
	fmt.Printf("Visitantes: \n")
	for i := 0; i < cantidad; i++ {
		sitio := heap.Desencolar()
		fmt.Printf("%s \n", sitio.url)
	}
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
	return 1
}

/*

 */
