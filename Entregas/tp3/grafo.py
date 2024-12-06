
#Implementacion de TDA Grafo
import random

class Grafo:
    def __init__(self,es_dirigido,vertices):
        self.es_dirigido = es_dirigido
        #Lo hacemos con dicc de dicc
        self.grafo = {}
        if vertices != None:
            for vertice in vertices:
                self.agregar_vertice(vertice)

    # dic -> { Vertice:{ Ady:Peso , Ady:Peso}, Vertice:{ Ady:Peso , Ady:Peso},etc.. }

    # Primitivas
    def agregar_vertice(self,vertice):
        self.grafo[vertice]={}
        
    def borrar_vertice(self,vertice):


    def agregar_arista(self,vertice,adyacente,peso)

    def borrar_arista(self,vertice,adyacente)

    def estan_unidos(self,vertice,adyacente):
        if adyacente in self.grafo.get(vertice,{}):
            return True
        return False

    def peso_arista(self,vertice,adyacente):
        return True

    def obtener_vertices(self): # -> devuelve una lista con todos los vertices del grafo
        return True

    def vertice_aleatorio(self):
        return random.choice(list(self.grafo.keys()))

    def adyacentes(self,vertice)