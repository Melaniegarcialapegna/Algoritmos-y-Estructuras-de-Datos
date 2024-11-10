import collections

def a_n_aristas(grafo, v, n):
    'Devolver una lista con los vértices que cumplen la propiedad'
    padres={}
    orden={}
    visitados = {}
    a_n_de_distancia = recorrido_bfs(grafo, padres, v, visitados, orden,n)
    
    return a_n_de_distancia

def recorrido_bfs(grafo,padres,vertice,visitados, orden, n):
    a_n_de_distancia = []

    padres[vertice] = None
    visitados[vertice] = True
    orden[vertice] = 0

    cola = collections.deque()
    cola.append(vertice)
    while len(cola) != 0:
        vertice = cola.popleft()
        if orden[vertice] == n:
            a_n_de_distancia.append(vertice)
            continue
        for adyacente in grafo.adyacentes(vertice):
            if adyacente not in visitados:
                cola.append(adyacente)
                padres[adyacente] = vertice
                visitados[adyacente] = True
                orden[adyacente] = orden[vertice] +1


    return a_n_de_distancia 