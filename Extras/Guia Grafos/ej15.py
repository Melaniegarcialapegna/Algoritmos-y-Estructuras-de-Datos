import collections

def obtener_orden(grafo):
    'Devolver una lista con un posible orden válido'
    cola = collections.deque()
    orden_cosas = []
    
    grados_entrada = buscar_grados_entrada(grafo)
    for vertice in grafo:
        if grados_entrada[vertice] == 0:
            cola.append(vertice)
    
    while len(cola) != 0:
        vertice = cola.popleft()
        orden_cosas.append(vertice)

        for adyacente in grafo.adyacentes(vertice):
                grados_entrada[adyacente] -= 1
                if grados_entrada[adyacente] == 0:
                    cola.append(adyacente)

    return orden_cosas

def buscar_grados_entrada(grafo):
    grados_entrada={}
    vertices = grafo.obtener_vertices()
    for vertice in vertices:
        grados_entrada[vertice] = 0
    for vertice in vertices:
         for ady in grafo.adyacentes(vertice):
              grados_entrada[ady] = grados_entrada.get(ady,0) +1
    return grados_entrada

