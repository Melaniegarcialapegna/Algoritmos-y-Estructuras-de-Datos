import collections

def encontrarCaminoMinimo(x,y):
    nuevosVertices = {} #NuevosVertices va a tener el numero multiplicado/restado y como dato la distancia
    if x == y:
        return 0
    
    cola= collections.deque()
    nuevosVertices[x] = 0
    cola.append(x)

    while len(cola) != 0:
        vertice = cola.popleft()

        #Si multiplico x2
        proxMult = vertice * 2
        if  proxMult not in nuevosVertices:
            nuevosVertices[proxMult] = nuevosVertices[vertice] +1
            cola.append(proxMult)

        #Si resto 1
        if vertice > 1: # si no me rompe xd
            proxResta = vertice -1
            if  proxResta not in nuevosVertices:
                nuevosVertices[proxResta] = nuevosVertices[vertice] +1
                cola.append(proxResta)

        #Me fijo si encontre a Y
        if y in nuevosVertices:
            return nuevosVertices[y]
        
    return -1