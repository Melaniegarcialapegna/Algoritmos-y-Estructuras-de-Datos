"""
a.  Lo planteamos como una matriz de adyacencia , en la cual los pesos varian entre : (agua=3,arbol=2,libre=1), en cada celda de la matriz pondriamos el peso correspondiente y luego usando el algoritmo de Dijkstra hallaremos el camino minimo para llegar a nuestro oponente 

"""
""" 
Matriz de incidencia : vertice * arista
Matriz de adyacencia : vertice * vertice 

Hallar caminos minimos : devuelve la distincia(orden) y los padres( para poder armar lista con el camino minimo)
Hallar MST (arbol de tendido minimo) : devuelve un grafo

"""
 
 #Averiguar -> incidencia !! como hacerla
##Armamos grafo con la matriz de adyacencia
def obtener_grafo(matriz):
    #grafo = {"1": {diccionario con vertices adyacentes}, "2": {}}
    grafo = {}
    #Inicializamos
    for i in range (0,len(matriz)):
        grafo[f"{i}"] = {} # Lo ponemos como cadena xd
        
    for indice_linea in range(0, len(matriz)):
        linea = matriz[indice_linea] 
        for indice_celda in range(0, len(linea)):
            if linea[indice_celda] != 0:
                grafo[f"{indice_linea}"][f"{indice_celda}"] = True

    return grafo

def dijkstra(grafo, desde, hasta):
    distancia = {}
    visitados = {}
    padres = {}
    heap = CrearHeap(compararDistancia)

    for vertice in grafo:
        distancia[vertice] = "inf"
    
    visitados[desde] = True
    distancia[desde] = 0
    padres[desde] = None
    vertice_distancia = crear_vertice_distancia(desde, distancia[desde])
    heap.encolar(vertice_distancia)

    while len(heap)!=0:
        vertice = heap.desencolar() #El heap ya agarra el mas chico
        
        if vertice == hasta:          #esto es para este casoxd
            return padres,distancia    #xd x2
                                       #xd x3
        for adyacente in grafo.adyacentes(vertice):
            if adyacente in visitados: #Tenemos que ver si le mejoramos o no la distancia
                condicion1 = distancia[adyacente] == "inf"
                condicion2 = (distancia[vertice] + grafo.peso_arista(vertice,adyacente)) < distancia[adyacente]
                if condicion1 or condicion2:
                    distancia[adyacente] = distancia[vertice] + grafo.peso_arista(vertice,adyacente)
            else:
                visitados[adyacente] = True
                padres[adyacente] = vertice
                distancia[adyacente] = distancia[vertice] + grafo.peso_arista(vertice,adyacente)

        return padres,distancia

def crear_vertice_distancia(vertice, distancia):
    return f"{vertice}-{distancia}"