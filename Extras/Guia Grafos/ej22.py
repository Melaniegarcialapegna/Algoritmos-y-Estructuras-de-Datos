def no_es_impar(grafo):
    padres = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            cantidad_es_impar = ciclo_dfs(grafo, vertice, padres, visitados)
            if cantidad_es_impar:
                return False
    return True

def ciclo_dfs(grafo,vertice,padre,visitados):
    
    for adyacente in grafo.adycentes(vertice):
        if adyacente in visitados:
            if padre[adyacente] != vertice:
                return detector_impar(vertice,adyacente,padre)
        else:
            padre[adyacente] = vertice
            visitados[adyacente] =  True
            detector = ciclo_dfs(grafo,adyacente,padre,visitados)
            if detector:
                return True
    return False

def detector_impar(vertice,adyacente,padre):
    ciclo= []
    while vertice != adyacente:
        ciclo.append(vertice)
        vertice = padre[vertice]
    ciclo.append(vertice)
    return len(ciclo) % 2 != 0
