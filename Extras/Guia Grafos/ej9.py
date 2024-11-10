def noche_museos(grafo):
    padres = {}
    visitados = {}
    recorrido = []
    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            ciclos =recorrido_dfs_clairo(grafo,vertice,padres,visitados, recorrido)
            if ciclos != None:
                return ciclos


def recorrido_dfs_clairo(grafo,vertice,padres,visitados, recorrido):
    recorrido.append(vertice)
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados:
                return recorrido
        else:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            recorrido = recorrido_dfs_clairo(grafo, adyacente, padres, visitados,recorrido)
            if recorrido is not None:
                return recorrido
    return recorrido