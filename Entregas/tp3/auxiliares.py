
from collections import deque

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

def reconstruir_camino(padres, origen, destino):
    camino = []
    actual = destino
    while actual != origen:
        camino.append(actual)
        actual = padres[actual]
    camino.append(origen)

    return camino[::-1]

def parsear_linea(linea):
    linea = linea.strip()
    return linea.split("\t")

def separar_listas(argumentos, separador):
    """Separa una lista en sublistas segun el separador"""
    resultado = []
    actual = []
    i = 0
    while i < len(argumentos):
        if argumentos[i] == separador:
            resultado.append(actual)
            i += 1
            actual = []
        else:
            actual.append(argumentos[i])
            i += 1
    if len(actual) != 0:
        resultado.append(actual)
    return resultado
        
def crear_cancion(lista_cancion):
    """Crea una cancion con el formato '(cancion,artista)' a partir de una lista"""
    return (" ".join(lista_cancion[0]), " ".join(lista_cancion[1]))

def obtener_playlist(usuarios, cancion, usuario):
    """Devuelve una playlist del usuario donde esta la cancion"""
    return list(usuarios[usuario][cancion].values())[0] 

def es_cancion(vertice):
    """Determina si un vertice es una cancion"""
    return type(vertice) == tuple