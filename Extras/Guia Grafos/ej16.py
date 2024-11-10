def convertir_a_matriz(grafo):
    'Devolver la Matriz construida'
    'El arreglo de mapeo debe contener a todos los vértices, en el mismo orden en el que están representados en la matriz.'
    matriz = []
    arreglo_mapeo = []
    
    for vertice in grafo:
        arreglo_mapeo.append(vertice)
        
    for vertice in arreglo_mapeo:
        fila = []
        for vertice2 in arreglo_mapeo:
            if grafo.estan_unidos(vertice, vertice2) and vertice != vertice2:
                fila.append(grafo.peso_arista(vertice, vertice2))
            else:
                fila.append(0)
        matriz.append(fila)
    return matriz, arreglo_mapeo


# [ [    A , B , C , D]
# [  A , 0 , C , D    ]
# [  B ,  , 0, -     ]
# [  C , - ,  ,     ]
# [  D , - , - ,   0]
# ]