import collections

def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'
    visitados = {}
    padres = {}

    grados_de_entrada = grados_entrada(grafo)
    grados_de_salida = grados_salida(grafo)
    lista_vertices_criticos = set()  # Asi no tenemos duplicados

    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            recorrido_bfs(grafo, vertice, padres, visitados, grados_de_entrada, grados_de_salida, lista_vertices_criticos)
    
    return list(lista_vertices_criticos)  #Lo hago lista

def recorrido_bfs(grafo, vertice, padre, visitados, grados_de_entrada, grados_de_salida, lista_vertices_criticos):
    cola = collections.deque()
    cola.append(vertice)

    while len(cola) != 0:
        vertice = cola.popleft()
        for adyacente in grafo.adyacentes(vertice):
            if adyacente not in visitados:
                cola.append(adyacente)
                padre[adyacente] = vertice
                visitados[adyacente] = True

                # Ahora verificamos si el vertice es critico
                if grados_de_entrada[adyacente] == 1 or grados_de_salida[adyacente] == 1:
                    lista_vertices_criticos.add(vertice)  # Para no tener duplicadosxd

def grados_entrada(g):
    return buscar_grados(g, tipo="entrada")

def grados_salida(g):
    return buscar_grados(g, tipo="salida")

def buscar_grados(grafo, tipo):
    gradoVertices = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            padres = {}
            padres[vertice] = None
            visitados[vertice] = True
            gradoVertices[vertice] = 0
            recorrido_dfs_grados(grafo, vertice, padres, gradoVertices, visitados, tipo)
    return gradoVertices

def recorrido_dfs_grados(grafo, vertice, padres, gradoVertices, visitados, tipo):
    
    if tipo == "salida":
        gradoVertices[vertice] = 0        

    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            visitados[adyacente] = True
            padres[adyacente] = vertice
            recorrido_dfs_grados(grafo, adyacente, padres, gradoVertices, visitados, tipo)

        if tipo == "salida":
            gradoVertices[vertice] += 1
        elif tipo == "entrada":
            if adyacente not in gradoVertices:
                gradoVertices[adyacente] = 0
            gradoVertices[adyacente] += 1
            