import collections

def diametro(grafo):
    diametro_maximo = 0

    for vertice in grafo:
        visitados = {}
        padres = {}
        orden = {}
        padres[vertice] = None
        visitados[vertice] = True
        orden[vertice] = 0
        
        diametro_actual = recorrido_bfs(grafo, vertice, padres, visitados, orden)
        if diametro_actual > diametro_maximo:
            diametro_maximo = diametro_actual
    
    return diametro_maximo

def recorrido_bfs(grafo,vertice,padres,visitados,orden):
    cola = collections.deque()
    orden_maximo = 0
    
    cola.append(vertice)
    while len(cola) != 0:
        vertice = cola.popleft()
        for adyacente in grafo.adyacentes(vertice):
            if adyacente not in visitados:
                cola.append(adyacente)
                padres[adyacente] = vertice
                visitados[adyacente] = True
                orden[adyacente] = orden[vertice]+1
                if orden[adyacente] > orden_maximo:
                    orden_maximo += 1
    return orden_maximo
            