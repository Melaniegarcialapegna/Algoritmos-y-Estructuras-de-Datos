def encontrar_ciclo(grafo):
    '''
    Devuelve una lista de vertices que conforman el ciclo. 
    Si no hay ciclo, debe devolver None. 
    '''

    padres = {}
    visitados = set()

    for vertice in grafo:
        if vertice not in visitados:
            iteracion_actual = set()
            iteracion_actual.add(vertice)
            visitados.add(vertice)
            padres[vertice] = None

            ciclo = dfs(grafo, vertice, visitados, padres, iteracion_actual)
            if ciclo is not None:
                return ciclo

    return None

def dfs(grafo, vertice, visitados, padres, iteracion_actual):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente in visitados and adyacente in iteracion_actual:
            return reconstruir_ciclo(padres, adyacente, vertice)

        visitados.add(adyacente)
        iteracion_actual.add(adyacente)
        padres[adyacente] = vertice
        ciclo = dfs(grafo, adyacente, visitados, padres, iteracion_actual)
        if ciclo is not None:
            return ciclo
    
    return None

def reconstruir_ciclo(padres, origen, final):
    actual = final
    ciclo = []

    while actual != origen:
        ciclo.append(actual)
        actual = padres[actual]
    ciclo.append(origen)

    ciclo = ciclo[::-1]

    return ciclo