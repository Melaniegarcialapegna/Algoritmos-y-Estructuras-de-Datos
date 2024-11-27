def amenazados(grafo):
    'Devolver una lista con los vértices que cumplen la condición'
    visitados = {}
    
    grados_de_entrada = grados_entrada(grafo)
    grados_de_salida = grados_salida(grafo)
    lista_vertices_criticos = set()  # Asi no tenemos duplicados

    for depredador in grafo:
        for presa in grafo.adyacentes(depredador):
            if grados_de_entrada[presa]==1:
                lista_vertices_criticos.add(depredador)
            if grados_de_salida[depredador]==1:
                lista_vertices_criticos.add(presa)
           
    return list(lista_vertices_criticos)  #Lo hago lista

def grados_entrada(grafo):
    grados = {}
    vertices = grafo.obtener_vertices()
    for vertice in vertices:
        grados[vertice] = 0
    for vertice in vertices:
        for ady in grafo.adyacentes(vertice):
            grados[ady]= grados.get(ady,0)+1
    return grados

def grados_salida(grafo):
    grados = {}
    for vertice in grafo:
        grados[vertice] = len(grafo.adyacentes(vertice))
    return grados

    