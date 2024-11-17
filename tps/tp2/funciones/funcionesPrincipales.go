package funciones

import (
	"bufio"
	"fmt"
	"os"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	"time"
)

const (
	DETECTOR_DOS = 2
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

// AgregarArchivo procesa un archivo de logs , guarda los datos en diccionarios y detecta DoS
func AgregarArchivo(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], sitios TDADiccionario.Diccionario[string, int], rutaArchivo string) string {
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return "Error en comando agregar_archivo"
	}
	scanner := bufio.NewScanner(archivo)
	diccArchivo := TDADiccionario.CrearHash[IP, []DatoLog]()
	for scanner.Scan() {
		datoLog := parsearLog(scanner.Text())
		agregarDatoLog(diccArchivo, datoLog.ip, datoLog)
		agregarDatoLog(diccOrdenado, datoLog.ip, datoLog)
		agregarVisita(sitios, datoLog.url)
	}
	//Detectar si hay DoS
	ipsDoS := detectarDoS(diccArchivo)
	ipsDoSOrdenadas := radixSort(ipsDoS)
	for _, ip := range ipsDoSOrdenadas {
		fmt.Printf("DoS: ")
		fmt.Printf("%s", string(ip))
		fmt.Printf("\n")
	}

	fmt.Println("OK")
	return ""
}

// VerVisitantes muestra las IPS que hicieron solicitudes dentro de un rango
func VerVisitantes(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], desde IP, hasta IP) {
	fmt.Println("Visitantes:")
	for iter := diccOrdenado.IteradorRango(&desde, &hasta); iter.HaySiguiente(); iter.Siguiente() {
		ip, _ := iter.VerActual()
		fmt.Printf("\t")
		fmt.Printf("%s", string(ip))
		fmt.Printf("\n")
	}
	fmt.Println("OK")
}

// VerMasVisitados muestra los sitios mas visitados basandose en el numero de visitas
func VerMasVisitados(diccionario TDADiccionario.Diccionario[string, int], cantidad int) {
	sitios := []Sitio{}
	for iter := diccionario.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		url, visitas := iter.VerActual()
		sitio := Sitio{url, visitas}
		sitios = append(sitios, sitio)
	}

	if cantidad > diccionario.Cantidad() {
		cantidad = diccionario.Cantidad()
	}
	heap := TDAHeap.CrearHeapArr(sitios, compararSitios)
	fmt.Printf("Sitios más visitados:\n")
	for i := 0; i < cantidad; i++ {
		sitio := heap.Desencolar()
		fmt.Printf("\t")
		fmt.Printf("%s", sitio.url)
		fmt.Printf(" - ")
		fmt.Printf("%d", sitio.visitas)
		fmt.Printf("\n")
	}
	fmt.Println("OK")
}
