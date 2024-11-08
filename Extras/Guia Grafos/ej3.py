def grados(g):
    return buscar_grados(g, tipo="salida")

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

def recorrido_dfs_grados(grafo, vertice, padres, gradoVertices , visitados, tipo):
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