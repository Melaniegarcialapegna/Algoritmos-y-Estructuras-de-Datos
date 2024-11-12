    #Asumimos que el grafo tiene el peso sin red
    # A -> B : p=4

    # agregar_arista(origen, destino) -> primitiva de ellos
    # arista.destino = vertice Y arista.origen = vertice Y ADEMAS el pesoRed
    # A -> B : p=2
import collections

def mst_cambiar_arista(grafo, mst, arista):
    #bfs
    padres = {}
    visitados = {}

    #Con el grafo original(sin modificar)
    peso_ciclo, vertices_original = recorrido_bfs(grafo, arista.origen, arista.destino, padres, visitados)
    
    #Actualizamos el grafo cambiandole el peso con el menor

    grafo.borrar_arista(arista.origen, arista.destino)    # Aca unicamente borramos la arista
    grafo.agregar_arista(arista.origen, arista.destino, arista.peso)    # Aca unicamente ponemos la arista con el nuevo peso red

    #bfs 2
    peso_ciclo_nuevo, vertices_nuevos = recorrido_bfs(grafo, arista.origen, arista.destino, padres, visitados)
    
    if peso_ciclo_nuevo == None or peso_ciclo_nuevo > peso_ciclo:
        return mst

    #hacer cosas
    # En caso de que con la arista red nos quede un ARBOL TENDIDO MINIMO distinto:
    return reconstruir_mst(grafo,mst,vertices_original, vertices_nuevos) # Le pasamos el grafo para consultar los pesos
    
def recorrido_bfs(grafo,vertice_origen,vertice_destino,padre,visitados):#devuelve los vertices del ciclo:
    cola = collections.deque()
    
    cola.append(vertice_origen)
    padre[vertice_origen] = None
    visitados[vertice_origen] = True
    
    cola.append(vertice_destino)
    padre[vertice_destino] = vertice_origen
    visitados[vertice_destino] = True
    
    
    while len(cola) != 0:
        vertice = cola.popleft()
        for adyacente in grafo.adyacentes(vertice):
            if adyacente in visitados:
                if padre[adyacente] != vertice:
                    return reconstruir_ciclo(grafo, padre, vertice, adyacente)
            else:
                padre[adyacente] = vertice
                visitados[adyacente] = True
                cola.append(adyacente)
    
    return None, []

def reconstruir_ciclo(grafo,padre,vertice,adyacente):#Devuelve el peso y los vertices
    vertices_ciclo = []
    peso = 0

    while vertice != adyacente:
        vertices_ciclo.append(vertice)
        peso += grafo.peso_arista(vertice, padre[vertice])
        vertice = padre[vertice]
    vertices_ciclo.append(vertice)
    # NO SE PONE LO DE ABAJO pq esa arista ya la recorrimos en el anterior
    # peso += grafo.peso_arista(vertice, padre[vertice])

    return peso , vertices_ciclo
    
def reconstruir_mst(grafo, mst,vertices_original, vertices_nuevo):
    ##Borramos todas las aristas
    for i in range(0, len(vertices_original)-1): #Pq usamos este y el sign
        origen = vertices_original[i]
        destino = vertices_original[i+1] #  Es de un ciclo q armamos
        mst.borrar_arista(origen, destino)
    
    ##Ponemos nuevas aristas integrando la nueva para el nuevo mst
    for i in range(0, len(vertices_nuevo)-1):
        origen = vertices_nuevo[i]
        destino = vertices_nuevo[i+1] #  Es de un ciclo q armamos
        nuevos_pesos = grafo.pesos_aristas(origen,destino)
        mst.agregar_arista(origen,destino,nuevos_pesos)
    return mst