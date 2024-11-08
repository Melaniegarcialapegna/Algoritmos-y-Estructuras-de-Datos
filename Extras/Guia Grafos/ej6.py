def es_arbol(g):
    return es_conexo_y_aciclico(g)

def es_conexo_y_aciclico(g):
    padres={}
    visitados = {}
    ya_entro = False
    for vertice in g:
        if ya_entro:
            return False
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            ya_entro = True
            ciclo = ciclo_dfs(g, vertice, padres, visitados)
            if ciclo:
                return False    

def ciclo_dfs(grafo, vertice, padres, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados:
            if padres[adyacente] != vertice:
                return True #Devuelve True cuando hay un ciclo
        else: 
            padres[adyacente] = vertice
            visitados[adyacente] = True
            hayCiclo = ciclo_dfs(grafo, adyacente, padres, visitados)
            if hayCiclo:
                return True
    return False