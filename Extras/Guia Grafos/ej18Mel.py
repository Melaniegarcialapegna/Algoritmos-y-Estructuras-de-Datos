import collections

def encontrar_minimos(x,y):
    distanciaCaminoMinimo = 0
    distancia = {}
    if x == y:
        return distanciaCaminoMinimo
    
    cola= collections.deque()
    cola.append(x)

    while len(cola) != 0:
        vertice = cola.popleft()

        #Si multiplico x2
        proxMult = vertice * 2
        if  proxMult not in distancia:
            distancia[proxMult] = distancia[vertice] +1
            cola.append(proxMult)

        #Si resto 1
        proxResta = vertice -1
        if  proxResta not in distancia:
            distancia[proxResta] = distancia[vertice] +1
            cola.append(proxResta)


