# Métodos del grafo:
# Grafo(es_dirigido = False, vertices_init = []) para crear un grafo no dirigido (hacer 'from grafo import Grafo')
# Grafo(es_dirigido = True, vertices_init = []) para crear un grafo dirigido (hacer 'from grafo import Grafo')
# agregar_vertice(self, v)
# borrar_vertice(self, v)
# agregar_arista(self, v, w, peso = 1)
# borrar_arista(self, v, w)
# estan_unidos(self, v, w)
# peso_arista(self, v, w)
# obtener_vertices(self)
# Devuelve una lista con todos los vértices del grafo
# vertice_aleatorio(self)
# adyacentes(self, v)
# str

####################################### EJ 0 #######################################

#Implementar un algoritmo que, dado un grafo no dirigido, nos devuelva un ciclo dentro del mismo, si es que los tiene. Indicar el orden del algoritmo.

# Una vez que nos topemos con un vertice ya visitado, ahi tenemos un posible ciclo. Esto es, si estoy viendo los adyacentes a un vértice dado, y dicho vértice está visitado, uno se apresuraría a decir que ahí se cierra un ciclo. Esto es cierto, salvo un caso: que dicho vértice visitado sea el antecesor a nuestro vértice en el recorrido (BFS o DFS)


#CON BFS
import collections

def encontrar_ciclo_bfs(grafo):
    q = collections.deque()
    visitados = {}
    padres = {}
    orden = {}
    primero = grafo.vertice_aleatorio()
    q.append(primero)
    ##Ponemos el primero en padre:nil
    
    padres[primero] = None
    orden[primero] = 0
    visitados[vertice] = True
    
    while not q.esta_vacia():
        vertice = q.pop()
        for adyacente in grafo.adyacentes(vertice):
            if adyacente in visitados:
                if padres[adyacente] != vertice:
                    return reconstruir_ciclo(padres ,vertice, adyacente)
            else:
                q.encolar(adyacente)
                padres[adyacente] = vertice
                visitados[adyacente] = True
                orden[adyacente] = orden[vertice] + 1



def reconstruir_ciclo(padres,vertice, adyacente):
    ciclo = []
    while vertice != adyacente:
        ciclo.append(vertice)
        vertice = padres[vertice]
    ciclo.append(adyacente)
    return ciclo[::-1]    


#CON DFS

def encontrar_ciclo(grafo):
    visitados = {}
    padre = {}
    for v in grafo:
        if v not in visitados:
            padre[v] = None  # Inicializar el nodo raíz sin padre
            ciclo = dfs_ciclo(grafo, v, visitados, padre)
            if ciclo is not None:
                return ciclo
    return None

def dfs_ciclo(grafo, v, visitados, padre):
    visitados[v] = True
    for w in grafo.adyacentes(v):
        if w in visitados:
            # Si w fue visitado y no es el padre de v, se detecta un ciclo
            if w != padre[v]:
                return reconstruir_ciclo(padre, w, v)
        else:
            padre[w] = v
            ciclo = dfs_ciclo(grafo, w, visitados, padre)
            if ciclo is not None:
                return ciclo

    return None

def reconstruir_ciclo(padre, inicio, fin):
    v = fin
    camino = []
    while v != inicio:
        camino.append(v)
        v = padre[v]
    camino.append(inicio)
    return camino[::-1]  