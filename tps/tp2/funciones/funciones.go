package funciones

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func AgregarArchivo(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], sitios TDADiccionario.Diccionario[string, int], rutaArchivo string) string {
	// abrimos el archivo -> O(n)
	archivo, err := os.Open(rutaArchivo)
	if err != nil {
		return "Error en comando agregar_archivo"
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
			diccOrdenado.Guardar(datoLog.ip, arreglo)
		}
		if !sitios.Pertenece(datoLog.url) {
			sitios.Guardar(datoLog.url, 1)
		} else {
			visitas := sitios.Obtener(datoLog.url)
			visitas++
			sitios.Guardar(datoLog.url, visitas)
		}
	}

	ipsDoS := []IP{}
	for iter := diccArchivo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, logs := iter.VerActual()
		contadorLogs := 1
		inicio := logs[0].fecha
		for i := 1; i < len(logs); i++ {
			diferencia := logs[i].fecha.Sub(inicio).Seconds()
			if diferencia < 2 {
				contadorLogs++
			} else {
				inicio = logs[i].fecha
				contadorLogs = 1
			}
			if contadorLogs >= 5 {
				ipsDoS = append(ipsDoS, ip)
				break
			}
		}
	}
	ipsDoSOrdenadas := radixSort(ipsDoS)
	for _, ip := range ipsDoSOrdenadas {
		fmt.Printf("DoS: ")
		fmt.Printf("%s", string(ip))
		fmt.Printf("\n")
	}
	fmt.Println("OK")

	return ""
}

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

func SepararIp(ip IP) []int {
	ipSeparada := strings.Split(string(ip), ".")
	numeros := []int{}
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(ipSeparada[i])
		numeros = append(numeros, num)
	}
	return numeros
}

func juntarIp(ipInt []int) IP {
	ipArrStr := []string{}
	for i := 0; i < 4; i++ {
		numStr := strconv.Itoa(ipInt[i])
		ipArrStr = append(ipArrStr, numStr)
	}
	ipStr := strings.Join(ipArrStr, ".")
	return IP(ipStr)
}
