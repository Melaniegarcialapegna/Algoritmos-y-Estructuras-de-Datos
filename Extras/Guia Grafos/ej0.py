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

# Una vez que nos topemos con un vértice ya visitado, ahí tenemos un posible ciclo. Esto es, si estoy viendo los adyacentes a un vértice dado, y dicho vértice está visitado, uno se apresuraría a decir que ahí se cierra un ciclo. Esto es cierto, salvo un caso: que dicho vértice visitado sea el antecesor a nuestro vértice en el recorrido (BFS o DFS)
