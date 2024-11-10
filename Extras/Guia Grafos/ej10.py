def es_bipartito(grafo):
    padres = {}
    visitados = {}
    for vertice in grafo:
        padres[vertice] = None
        visitados[vertice] = 0
        ciclos_par = recorrido_dfs(grafo,vertice,padres,visitados)
        if not ciclos_par:
            return False
    return True
    
def recorrido_dfs(grafo, vertice, padres, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            if padres[adyacente]!= vertice:
                es_par = contar_vertices_ciclo(vertice,adyacente,padres)
                if not es_par:
                    return False
                    
        else:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            ciclo = recorrido_dfs(grafo, adyacente, padres, visitados)
            if ciclo is not None:
                return ciclo
    return True

def contar_vertices_ciclo(vertice,adyacente,padres):
    contador = 0
    while vertice != adyacente:
        contador += 1
        vertice = padres[vertice]
    contador += 1
    return contador % 2 == 0

