package funciones

// radixSort ordena las IPS de menor a mayor
func radixSort(ips []IP) []IP {
	arrIps := [][]int{}
	for _, ip := range ips {
		arrIps = append(arrIps, separarIp(ip))
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

// countingSort ordena IPS basandose en un criterio
func countingSort(arrIps [][]int, criterio func([]int) int) [][]int {

	frecuencias := make([]int, 256)
	for _, ip := range arrIps {
		frecuencias[criterio(ip)]++
	}

	sumaFrecuencias := make([]int, 256)
	for i := 1; i < 256; i++ {
		sumaFrecuencias[i] = sumaFrecuencias[i-1] + frecuencias[i-1]
	}

	arrIpsOrdenado := make([][]int, len(arrIps))
	for _, ip := range arrIps {
		valor := criterio(ip)
		indice := sumaFrecuencias[valor]
		arrIpsOrdenado[indice] = ip
		sumaFrecuencias[valor]++
	}
	return arrIpsOrdenado

}

// Criterios
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
