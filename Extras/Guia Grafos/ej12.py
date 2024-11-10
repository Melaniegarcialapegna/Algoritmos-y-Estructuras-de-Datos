def puede_ser_no_dirigido(grafo):
    visitados = {}
    padres = {}
    
    for vertice in grafo:
        padres[vertice] = None
        visitados[vertice] = True
        tiene_ida_y_vuelta = recorrido_dfs(grafo, vertice, padres, visitados)
        if not tiene_ida_y_vuelta:
            return False
    
    return True

def recorrido_dfs(grafo,vertice,padres,visitados):
    for adyacente in grafo.adyacentes(vertice):
        #Si tienen ida y vuelta -> son no dirigidos
        if not grafo.estan_unidos(adyacente, vertice):
            return False
        if adyacente not in visitados:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            tiene_ida_y_vuelta = recorrido_dfs(grafo,adyacente,padres,visitados)
            if not tiene_ida_y_vuelta:
                return False
    return True