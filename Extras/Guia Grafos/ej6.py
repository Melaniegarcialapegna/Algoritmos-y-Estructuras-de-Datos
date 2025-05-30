def es_arbol(g):
    return es_conexo(g) and  es_aciclico(g)


def es_conexo(grafo):
    padres={}
    visitados = {}
    orden = {}
    ya_entro = False
    for vertice in grafo:
        if vertice not in visitados:
            if ya_entro:
                return False
            padres[vertice] = None
            visitados[vertice] = True
            ya_entro = True
            ciclo_dfs(grafo, vertice, orden, padres, visitados)
    return True   


def es_aciclico(g):
    padres={}
    visitados = {}
    orden = {}
    for vertice in g:
        if vertice not in visitados:
            visitados[vertice] = True
            padres[vertice] = None
            ciclo= ciclo_dfs(g, vertice, orden, padres, visitados)
            if ciclo:
                return False
    return True   

def ciclo_dfs(grafo, vertice, padres, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados:
            if adyacente != padres[vertice]:
                return True 
        else:
            padres[adyacente]= vertice
            ciclo = ciclo_dfs(grafo,adyacente,padres,visitados)
            if ciclo:
                return True
    return False







def ciclo_dfs(grafo, vertice, orden, padres, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados:
            mayor = adyacente
            menor = vertice
            if orden[vertice] > orden[adyacente]:
                mayor = vertice
                menor = adyacente
            if orden[mayor] - orden[menor] > 1:
                return True
        else: 
            padres[adyacente] = vertice
            visitados[adyacente] = True
            orden[adyacente] = orden[vertice]+1
            hayCiclo = ciclo_dfs(grafo, adyacente, padres, visitados)
            if hayCiclo:
                return True
    return False