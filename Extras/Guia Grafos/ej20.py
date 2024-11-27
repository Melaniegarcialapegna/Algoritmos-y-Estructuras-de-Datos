def cantidad_componentes_debiles(grafo):
    cant_comp = 0
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            visitados[vertice] = True
            recorrido_dfs(grafo,vertice,visitados)
            cant_comp += 1
    return cant_comp


def recorrido_dfs(grafo, vertice, visitados):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            visitados[adyacente] = True
            recorrido_dfs(grafo,adyacente, visitados)
    return None

