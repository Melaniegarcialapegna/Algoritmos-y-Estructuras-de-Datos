def es_conexo(grafo):
    ya_entro = False
    visitados = {}
    padres = {}
    for vertice in grafo:
        if vertice not in visitados:
            if ya_entro:
                return False
            ya_entro = True
            padres[vertice] = None
            visitados[vertice] = True
            recorrido_dfs(grafo,vertice,padres,visitados)
    return True

def recorrido_dfs(grafo,vertice,padres,visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            recorrido_dfs(grafo,adyacente,padres,visitados)
            