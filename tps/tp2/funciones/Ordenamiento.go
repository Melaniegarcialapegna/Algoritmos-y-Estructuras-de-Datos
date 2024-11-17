package funciones

func radixSort(ips []IP) []IP {

	arrIps := [][]int{}
	for _, ip := range ips {
		arrIps = append(arrIps, SepararIp(ip))
	}

	arrIps = countingSort(arrIps, criterioCuarto)
	arrIps = countingSort(arrIps, criterioTercero)
	arrIps = countingSort(arrIps, criterioSegundo)
	arrIps = countingSort(arrIps, criterioPrimero)

	ipsOrdenada := []IP{}
	for _, ip := range arrIps {
		ipsOrdenada = append(ipsOrdenada, juntarIp(ip))
	}
	return ipsOrdenada
}

func countingSort(arrIps [][]int, criterio func([]int) int) [][]int {

	//Primer paso -> FRECUENCIAS
	frecuencias := make([]int, 256) // 256 valores posibles
	for _, ip := range arrIps {
		frecuencias[criterio(ip)]++
	}

	//Segundo paso -> INDICES - SUMAR FRECUENCIAS
	sumaFrecuencias := make([]int, 256)
	for i := 1; i < 256; i++ {
		sumaFrecuencias[i] = sumaFrecuencias[i-1] + frecuencias[i-1] // frecuencia[i-1]:cantidad q habia del anterior
	}

	//Tercer paso -> ACOMODAR XD
	arrIpsOrdenado := make([][]int, len(arrIps))
	for _, ip := range arrIps {
		valor := criterio(ip)
		indice := sumaFrecuencias[valor]
		arrIpsOrdenado[indice] = ip
		sumaFrecuencias[valor]++
	}
	return arrIpsOrdenado

}

func criterioCuarto(ip []int) int {
	return ip[3]
}
func criterioTercero(ip []int) int {
	return ip[2]
}
func criterioSegundo(ip []int) int {
	return ip[1]
}
func criterioPrimero(ip []int) int {
	return ip[0]
}
