def cantidad_componentes_debiles(grafo):
    cant_comp = 0

    visitados = {}
    padres = {}

    for vertice in grafo:
        if vertice not in visitados:
            visitados[vertice] = True
            padres[vertice] = None
            lista_adyacentes = []
            lista_adyacentes.append(vertice)
            recorrido_dfs(grafo,padres,vertice,visitados,lista_adyacentes)
            cant_comp += 1
            # if len(lista_adyacentes) == 1:
            #     cant_comp += 1
            # for verticeA in lista_adyacentes:
            #     for verticeB in lista_adyacentes:
            #         es_debilmente = False
            #         if verticeA != verticeB:
            #             if not (grafo.estan_unidos(verticeA, verticeB)) and not (grafo.estan_unidos(verticeB, verticeA)):
            #                 cant_comp += 1
            #                 es_debilmente = True  
            #         if es_debilmente:
            #             break
    return cant_comp


def recorrido_dfs(grafo, padres, vertice, visitados, vertices_actual):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            vertices_actual.append(adyacente)
            recorrido_dfs(grafo, padres, adyacente, visitados, vertices_actual)
    return None