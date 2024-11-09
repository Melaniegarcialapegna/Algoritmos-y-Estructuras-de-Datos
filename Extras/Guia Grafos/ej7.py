from grafo import Grafo

def grafo_traspuesto(grafo):
    #Devolver el nuevo Grafo
    nuevoGrafo = Grafo(es_dirigido=True, vertices_init = grafo.obtener_vertices())
    padres = {}
    visitados = {}
    for vertice in grafo:
        if vertice not in visitados:
            padres[vertice] = None
            visitados[vertice] = True
            recorrido_dfs_trasponer(grafo,nuevoGrafo,vertice,visitados,padres)
    return nuevoGrafo

def recorrido_dfs_trasponer(grafo, nuevoGrafo, vertice, visitados,padres ):
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            padres[adyacente] = vertice
            visitados[adyacente] = True
            nuevoGrafo.agregar_arista(adyacente, vertice)
            recorrido_dfs_trasponer(grafo,nuevoGrafo,adyacente,visitados,padres)
        else:
            nuevoGrafo.agregar_arista(adyacente,vertice)