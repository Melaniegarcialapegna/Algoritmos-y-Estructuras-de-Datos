import grafo
from collections import deque

#Recorridos

#DFS
def recorrido_dfs(grafo):
    visitados = set()
    padres ={}
    orden={}
    for vertice in grafo.obtener_vertices():
        if vertice not in visitados:
            padres[vertice]= None
            orden[vertice] = 0
            dfs(grafo,vertice,visitados,padres,orden)
    return padres,orden

def dfs (grafo,vertice,visitados,padres,orden):
    visitados.add(vertice)
    for ady in grafo.adyacentes(vertice):
        if ady not in visitados:
            padres[ady]= vertice
            orden[ady]= orden[vertice]+1
            dfs(grafo,vertice,visitados,padres,orden)


#Camino Minimo

#BFS
def bfs(grafo,origen):
    visitados=set()
    padres= {}
    orden={}
    cola = deque()
    visitados.add(origen)
    padres[origen]=None
    orden[origen]=0
    cola.append(origen)
    while len(cola)!=0:
        vertice = cola.popleft()
        for ady in grafo.adyacentes(vertice):
            if ady not in visitados:
                visitados.add(ady)
                padres[ady]= vertice
                orden[ady]= orden[vertice]+1
                cola.append(ady)
    return padres,orden

#Dijsktra
def dijsktra(grafo,origen):
    dist={}
    padres={}
    for vertice in grafo.obtener_vertices():
        dist[vertice]=float("inf")
    dist[origen]=0
    padres[origen]=None
    cola=deque()#VER COMO HACERLO HEAP MINIMOS
    cola.append(origen,0)
    while len(cola)!=0:
        vertice= cola.popleft()
        for ady in grafo.adyacentes():
            actual= dist[vertice]+grafo.peso_union(vertice,ady)
            if actual < dist[ady]:
                dist[ady]= actual
                padres[ady]= vertice
                cola.append(ady,actual)
    return padres, dist

#Bellman(?) Hace falta hacerlo??¿


