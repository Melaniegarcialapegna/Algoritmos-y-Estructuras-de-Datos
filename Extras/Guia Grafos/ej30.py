from grafo import Grafo
import collections

# El grafo recibido tiene si o si pesos 1 o 2
def camino_minimo(grafo, origen):
    # Creo grafo auxiliar -> O(v)
    grafoAux = Grafo(es_dirigido = False,vertices_init=grafo.obtener_vertices()) #

    #Recorro grafo original y agrego aristas a grafoAux
    # -> si la arista tiene peso 2, meto un vertice en el medio
    vertices_intermedios = set() 
    agregar_aristas(grafo, origen, grafoAux, vertices_intermedios)

    padres_con_intermedios = recorrido_bfs(grafoAux, origen)

    padres = sacar_intermedios(padres_con_intermedios)

    return padres

def agregar_aristas(grafo, vertice, grafoAux, vertices_agregados):
    for adyacente in grafo.adyacentes(vertice):
        intermedio = f"_{vertice}-{adyacente}"
        if grafoAux.estan_unidos(vertice, adyacente) or intermedio in vertices_agregados:
            continue

        if grafo.peso_arista(vertice, adyacente) == 1:
            grafoAux.agregar_arista(vertice, adyacente, peso=1)
        else:
            grafoAux.agregar_vertice(intermedio)
            vertices_agregados.add(intermedio)
            grafoAux.agregar_arista(vertice, intermedio, peso=1)
            grafoAux.agregar_arista(intermedio, adyacente, peso=1)
        agregar_aristas(grafo, adyacente, grafoAux, vertices_agregados)

def recorrido_bfs(grafo, origen):
    cola = collections.deque()
    visitados = set()
    padres = {}
    padres[origen] = None
    visitados.add(origen)
    cola.append(origen)
    while len(cola) != 0:
        actual = cola.popleft()
        for adyacente in grafo.adyacentes(actual):
            if not adyacente in visitados:
                visitados.add(adyacente)
                padres[adyacente] = actual
                cola.append(adyacente)
    
    return padres

def sacar_intermedios(padres):
    resultado = {}
    for vertice in padres:
        if vertice[0] == "_":
            continue
        padre = padres[vertice]
        if padre != None and padre[0] == "_":
            padre = padres[padre]
        resultado[vertice] = padre
    return resultado