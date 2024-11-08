def encontrar_ciclo(grafo):
    padres = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            ciclo = ciclo_dfs(grafo, vertice,padres,visitados)
            if ciclo is not None:
                return ciclo
    return None

def ciclo_dfs(grafo, vertice, padres, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados:
            return recrear_ciclo(padres,vertice,adyacente)            
        else:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            ciclo =  ciclo_dfs(grafo, adyacente, padres, visitados)
            if ciclo is not None:
                return ciclo
    return None

            
def recrear_ciclo(padres, vertice, adyacente):
    ciclo = []
    while adyacente != vertice:
        ciclo.append(vertice)
        vertice = padres[vertice]
    ciclo.append(vertice)
    return ciclo[::-1]  