import random 

class Grafo:
    def __init__(self,es_dirigido,vertices=None):
        self.es_dirigido = es_dirigido
        self.vertices = {}
        if vertices != None:
            for vertice in vertices:
                self.agregar_vertice(vertice)

    #Iterador
    def __iter__(self):
        self.lista = self.obtener_vertices()
        self.indice = 0
        return self

    def __next__(self):
        if self.indice >= len(self.lista):
            raise StopIteration
        
        elemento = self.lista[self.indice]
        self.indice += 1
        return elemento

    def __contains__(self, vertice):
        return vertice in self.vertices

    # Primitivas
    def agregar_vertice(self,vertice):
        if vertice in self.vertices:
            raise Exception("ERROR: El vertice ya pertenece al grafo")
        self.vertices[vertice]={}

    def borrar_vertice(self,vertice):
        if vertice not in self.vertices:
            raise Exception("ERROR: El vertice no pertenece al grafo")
        
        for verticeActual in self.vertices :
            if vertice in verticeActual:
                del verticeActual[vertice]

        del self.vertices[vertice]

    def agregar_arista(self,vertice,adyacente,peso=1):
        if vertice not in self.vertices or adyacente not in self.vertices:
            raise Exception("ERROR: Uno de los vertices no esta en el grafo")
        
        self.vertices[vertice][adyacente] = peso
        
        if not self.es_dirigido: 
            self.vertices[adyacente][vertice] = peso
        
    def borrar_arista(self,vertice,adyacente):
        if vertice not in self.vertices or adyacente not in self.vertices:
            raise Exception("ERROR: Uno de los vertices no esta en el grafo")
        
        del self.vertices[vertice][adyacente] 
        
        if not self.es_dirigido: 
            del self.vertices[adyacente][vertice]

    def estan_unidos(self,vertice,adyacente):
        return adyacente in self.vertices[vertice]

    def peso_arista(self,vertice,adyacente):
        if vertice not in self.vertices or adyacente not in self.vertices:
            raise Exception("ERROR: Uno de los vertices no esta en el grafo")

        if not self.estan_unidos(vertice, adyacente):
            raise Exception("ERROR: Los ver")

        return self.vertices[vertice][adyacente]        

    def obtener_vertices(self): 
        return self.vertices.keys()

    def vertice_aleatorio(self):
        return random.choice(list(self.vertices.keys()))

    def adyacentes(self, vertice):
        if vertice not in self.vertices:
            raise Exception("ERROR: El vertice no pertenece al grafo")
        
        return list(self.vertices[vertice].keys())