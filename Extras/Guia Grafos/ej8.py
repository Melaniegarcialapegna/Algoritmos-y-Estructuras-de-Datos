import collections

def seis_grados(grafo):
    for vertice in grafo:
        mas_de_6 = recorrido_bfs(grafo, vertice)
        if not mas_de_6:
            return False
    return True

def recorrido_bfs(grafo, vertice):
    padres = {}
    orden = {}
    visitados = {}
    
    padres[vertice] = None
    orden[vertice] = 0
    visitados[vertice] = True
    
    cola = collections.deque()
    cola.append(vertice)
    
    while len(cola) != 0:
        vertice = cola.popleft()
        if orden[vertice] > 6: #Si en algun momento la distancia es mayor a 6
            return False
        for adyacente in grafo.adyacentes(vertice):
            if adyacente not in visitados:
                cola.append(adyacente)
                visitados[adyacente] = True
                padres[adyacente] = vertice
                orden[adyacente] = orden[vertice] + 1
    return True #La distancia nunca fue mayor a 6