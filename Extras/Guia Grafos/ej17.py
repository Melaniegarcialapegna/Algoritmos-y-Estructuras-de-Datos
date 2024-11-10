def comprobar_teorema(grafo):
    contador = 0
    for vertice in grafo:
        cantAdyacentes = 0
        for adyacentes in grafo.adyacentes(vertice):
            cantAdyacentes += 1
        if cantAdyacentes %2 != 0:
            contador +=1

    return contador%2 == 0 