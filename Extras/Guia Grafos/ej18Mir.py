def encontrarCamino(x, y, anterior, distancia):
	if x == y :
		return distancia

	izq = x - 1
	der = 2 * x

	if abs(y-izq) < abs(y-x):
		distIzq = encontrarCamino(izq, y, x, distancia+1) 
	if abs(y-der) < abs(y-x):
		distDer = encontrarCamino(der, y, x, distancia+1)

	if distIzq < distDer:
		return distIzq
        
	return distDer
