def es_arbol(g):
    return es_conexo(g) and es_aciclico(g)

def es_conexo(grafo):
    padres={}
    visitados = {}
    orden = {}
    contador = 0
    for vertice in grafo:
        if vertice not in visitados:
            contador +=1
            padres[vertice] = None
            visitados[vertice] = True
            orden[vertice] = 0
            ciclo_dfs(grafo, vertice,orden ,padres, visitados)
    return contador == 1


def es_aciclico(g):
    padres={}
    visitados = {}
    orden = {}
    for vertice in g:
        if vertice not in visitados:
            visitados[vertice] = True
            padres[vertice] = None
            orden[vertice] =0
            ciclo= ciclo_dfs(g, vertice, orden ,padres, visitados)
            if ciclo:
                return False
    return True   

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
            hayCiclo = ciclo_dfs(grafo, adyacente, orden, padres, visitados)
            if hayCiclo:
                return True
    return False
