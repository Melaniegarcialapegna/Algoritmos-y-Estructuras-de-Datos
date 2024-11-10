func encontrarCamino(x, y int, distancia int) int {
	if x == y {
		return distancia
	}

	izq := x - 1
	der := 2 * x

	distIzq := encontrarCamino(izq, y, distancia+1)
	distDer := encontrarCamino(der, y, distancia+1)

	if distIzq < distDer {
		return distIzq
	}
	return distDer

}
