import collections

def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'
    visitados = {}
    padres = {}

    grados_de_entrada = grados_entrada(grafo)
    grados_de_salida = grados_salida(grafo)
    
    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            vertices_criticos = recorrido_bfs(grafo, vertice, padres, visitados, grados_de_entrada, grados_de_salida)
            return vertices_criticos

    return []

def recorrido_bfs(grafo,vertice,padre,visitados,grados_de_entrada, grados_de_salida):

    cola = collections.deque()
    cola.append(vertice)
    criticos = {}
    lista_criticos = []

    while len(cola) != 0:
        vertice = cola.popleft()
        for adyacente in grafo.adyacentes(vertice):
            cola.append(adyacente)
            padre[adyacente] = vertice
            visitados[adyacente] = True
            
            #Ahora empezamos a chequear si es o no punto critico
            if grados_de_entrada[adyacente] ==1 or grados_de_salida[adyacente] == 1:
                criticos[vertice] = True

    for elemento in criticos.keys():
        lista_criticos.append(elemento)


def grados_entrada(g):
    return buscar_grados(g, tipo = "entrada")

def grados_salida(g):
    return buscar_grados(g, tipo = "salida")

def buscar_grados(grafo, tipo):
    gradoVertices = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            #Si el vertice NO esta visitado es pq es le primero de una componente
            padres = {}
            padres[vertice] = None
            visitados[vertice]=  True
            gradoVertices[vertice] = 0
            recorrido_dfs_grados(grafo,vertice,padres,gradoVertices,visitados,tipo)
    return gradoVertices

def recorrido_dfs_grados(grafo, vertice, padres, gradoVertices, visitados, tipo):
    ##Tenemos q inicializar si es de salida
    if tipo =="salida":
        gradoVertices[vertice] = 0        

    for adyacente in grafo.adyacentes(vertice):
        if not adyacente in visitados:
            visitados[adyacente] = True
            padres[adyacente] = vertice
            recorrido_dfs_grados(grafo, adyacente, padres, gradoVertices, visitados,tipo)
                        
        if tipo == "salida":
            gradoVertices[vertice] += 1
        elif tipo == "entrada":
            if adyacente not in gradoVertices:
                gradoVertices[adyacente] =0         
            gradoVertices[adyacente] += 1