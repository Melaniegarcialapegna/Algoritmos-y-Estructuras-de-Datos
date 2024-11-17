package funciones

import (
	"strconv"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
)

const (
	MIN_PETICIONES_DOS = 5
)

// parsearLog convierte una linea de log en un DatoLog
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

// compararSitios compara dos sitios basandose en la cantidad de visitas
func compararSitios(s1, s2 Sitio) int {
	if s1.visitas > s2.visitas {
		return 1
	} else if s1.visitas < s2.visitas {
		return -1
	}
	return 0
}

// SepararIp convierte una IP en un array
func separarIp(ip IP) []int {
	ipSeparada := strings.Split(string(ip), ".")
	numeros := []int{}
	for i := 0; i < 4; i++ {
		num, _ := strconv.Atoi(ipSeparada[i])
		numeros = append(numeros, num)
	}
	return numeros
}

// juntarIp convierte un array en una IP
func juntarIp(ipInt []int) IP {
	ipArrStr := []string{}
	for i := 0; i < 4; i++ {
		numStr := strconv.Itoa(ipInt[i])
		ipArrStr = append(ipArrStr, numStr)
	}
	ipStr := strings.Join(ipArrStr, ".")
	return IP(ipStr)
}

// agregarDatoLog agrega un DatoLog a un diccionario de logs
func agregarDatoLog(dicc TDADiccionario.Diccionario[IP, []DatoLog], ip IP, datoLog DatoLog) {
	if !dicc.Pertenece(ip) {
		arreglo := []DatoLog{datoLog}
		dicc.Guardar(ip, arreglo)
	} else {
		arreglo := dicc.Obtener(ip)
		arreglo = append(arreglo, datoLog)
		dicc.Guardar(ip, arreglo)
	}
}

// agregarVisita incrementa el contador de visitas de una URL
func agregarVisita(sitios TDADiccionario.Diccionario[string, int], url string) {
	if !sitios.Pertenece(url) {
		sitios.Guardar(url, 1)
	} else {
		visitas := sitios.Obtener(url)
		visitas++
		sitios.Guardar(url, visitas)
	}
}

// detectarDoS verifica si alguna IP que ha realizado mas de 5 peticiones en un lapso de tiempo
func detectarDoS(dicc TDADiccionario.Diccionario[IP, []DatoLog]) []IP {
	ipsDoS := []IP{}
	for iter := dicc.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, logs := iter.VerActual()
		contadorLogs := 1
		inicio := logs[0].fecha
		for i := 1; i < len(logs); i++ {
			diferencia := logs[i].fecha.Sub(inicio).Seconds()
			if diferencia < DETECTOR_DOS {
				contadorLogs++
			} else {
				inicio = logs[i].fecha
				contadorLogs = 1
			}
			if contadorLogs >= MIN_PETICIONES_DOS {
				ipsDoS = append(ipsDoS, ip)
				break
			}
		}
	}
	return ipsDoS
}
