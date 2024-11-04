package tp2

import (
	"bufio"
	"os"
	"strings"
	TDADiccionario "tdas/diccionario"
	"time"
)

func AgregarArchivo(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog], sitios TDADiccionario.Diccionario[Sitio, int], rutaArchivo string) {
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

		if !sitios.Pertenece(Sitio(datoLog.url)) {
			sitios.Guardar(Sitio(datoLog.url), 1)
		} else {
			visitas := sitios.Obtener(Sitio(datoLog.url))
			visitas++
			sitios.Guardar(Sitio(datoLog.url), visitas)
		}
	}

	for iter := diccArchivo.Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		ip, logs := iter.VerActual()
		for _, log := range logs {

		}
	}

	return

}

func VerVisitantes(diccOrdenado TDADiccionario.DiccionarioOrdenado[IP, []DatoLog]) {
	return
}

func VerMasVisitados(diccionario TDADiccionario.Diccionario[Sitio, int]) {
	return
}

func parsearLog(linea string) DatoLog {
	campos := strings.Fields("\t")
	fecha, _ := time.Parse("2006-01-02T15:04:05-07:00", campos[1])

	return DatoLog{
		ip:         IP(campos[0]),
		fecha:      fecha,
		metodoHttp: campos[2],
		url:        campos[3],
	}
}

func crearSitioVisitantes(dato DatoLog) SitioVisitantes {
	return SitioVisitantes{
		sitio:   Sitio(dato.url),
		visitas: 0,
	}
}

/*

 */
