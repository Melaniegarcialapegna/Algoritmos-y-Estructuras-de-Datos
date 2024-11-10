import collections

def obtener_orden(grafo):
    'Devolver una lista con un posible orden válido'
    
    visitados = {}
    padres = {}
    orden = {}
    cola = collections.deque()
    orden_cosas = []
    
    grados_entrada = buscar_grados_entrada(grafo)
    for vertice in grafo:
        if grados_entrada[vertice] == 0:
            cola.append(vertice)
    
    while len(cola) != 0:
        vertice = cola.popleft()
        orden_cosas.append(vertice)

        for adyacente in grafo.adyacentes(vertice):
                grados_entrada[adyacente] -= 1
                if grados_entrada[adyacente] == 0:
                    cola.append(adyacente)

    return orden_cosas

def buscar_grados_entrada(grafo):
    gradoVertices = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            #Si el vertice NO esta visitado es pq es le primero de una componente
            padres = {}
            padres[vertice] = None
            visitados[vertice]=  True
            gradoVertices[vertice] = 0
            recorrido_dfs_grados(grafo,vertice,padres,gradoVertices,visitados)
    return gradoVertices

def recorrido_dfs_grados(grafo, vertice, padres, gradoVertices , visitados):
    for adyacente in grafo.adyacentes(vertice):
        if not adyacente in visitados:
            visitados[adyacente] = True
            padres[adyacente] = vertice
            recorrido_dfs_grados(grafo, adyacente, padres, gradoVertices, visitados)
            if adyacente not in gradoVertices:
                gradoVertices[adyacente] =0         
            gradoVertices[adyacente] += 1