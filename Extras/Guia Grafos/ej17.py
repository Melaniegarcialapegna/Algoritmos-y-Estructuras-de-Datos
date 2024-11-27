def comprobar_teorema(grafo):
    contador = 0
    for vertice in grafo:
        cantAdyacentes = len(grafo.adyacentes(vertice))
        if cantAdyacentes %2 != 0:
            contador +=1

    return contador%2 == 0 