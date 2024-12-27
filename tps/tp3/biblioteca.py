import auxiliares ,random , heapq , collections

CANT_RANDOM_WALKS = 500
CANT_PAGE_RANKS = 100
LARGO_RANDOM_WALK = 100
COEFICIENTE_AMORTIGUACION = 0.85

#Camino mas Corto
def camino_mas_corto(grafo,origen,destino):
    """Devuelve el camino mas corto entre dos canciones"""
    padres, distancia = auxiliares.bfs(grafo, origen)

    if not destino in distancia:
        return None

    return auxiliares.reconstruir_camino(padres, origen, destino)

#Mas importantes
def canciones_mas_importantes(grafo):
    """Devuelve un diccionario con las canciones mas importantes obtenidas utilizando el algoritmo de pageranks"""
    pageranks = {}
    for i in range(CANT_PAGE_RANKS):
        page_rank(grafo, pageranks)
    return pageranks

def page_rank(grafo, dicc_pageranks):
    """Calcula el page rank de cada vertice del grafo"""
    vertices = grafo.obtener_vertices()
    for vertice in vertices:
        tipo_actual = "canciones" if auxiliares.es_cancion(vertice) else "usuarios"
        pageRank_articulo = (1-COEFICIENTE_AMORTIGUACION)/len(vertices)
        for adyacente in grafo.adyacentes(vertice):
            tipo_ady = "canciones" if auxiliares.es_cancion(adyacente) else "usuarios"
            dicc_pageranks[tipo_ady] = dicc_pageranks.get(tipo_ady, {})
            pagerank_adyacente = dicc_pageranks[tipo_ady].get(adyacente, 0)
            cant_adyacentes = len(grafo.adyacentes(adyacente))
            pageRank_articulo += COEFICIENTE_AMORTIGUACION * (pagerank_adyacente / cant_adyacentes) 
        dicc_pageranks[tipo_actual] = dicc_pageranks.get(tipo_actual, {})
        dicc_pageranks[tipo_actual][vertice] = pageRank_articulo

#Recomendaciones
def recomendacion(grafo, tipo, vertices, cantidad): 
    """Devuelve las 'cantidad' canciones mas recomendadas"""
    probabilidades = {}

    for vertice in grafo.obtener_vertices():
        tipo_actual = "canciones" if auxiliares.es_cancion(vertice) else "usuarios"
        probabilidades[tipo_actual] = probabilidades.get(tipo_actual, {})
        probabilidades[tipo_actual][vertice] = 1

    for vertice in vertices:
        for i in range(CANT_RANDOM_WALKS):
            random_walk(grafo, vertice, 1, probabilidades, 200, True)
    
    return heapq.nlargest(cantidad, probabilidades[tipo].items(), compararPageRank)

def compararPageRank(elemento):
    return elemento[1]

def random_walk(grafo, vertice, probabilidad, probabilidades, largoMax,primeraIteracion):
    """Realiza un random walk de largoMax sobre el grafo"""
    if largoMax == 0:
        return

    largoMax -= 1
    tipo = "canciones" if auxiliares.es_cancion(vertice) else "usuarios"
    
    if not primeraIteracion:
        probabilidades[tipo] = probabilidades.get(tipo, {})
        probabilidades[tipo][vertice] += probabilidad
    
    adyacentes = grafo.adyacentes(vertice)
    probabilidad_sig = probabilidad/len(adyacentes)

    siguiente = random.choice(adyacentes)
    random_walk(grafo, siguiente, probabilidad_sig, probabilidades, largoMax,False)


#Ciclo n Canciones
def ciclo_n_canciones(grafo, cancion, n):
    """Devuelve un ciclo de 'n' canciones"""
    padres = {}
    padres[cancion]= None
    visitados = set()
    visitados.add(cancion)
    ciclo = _ciclo_n_canciones(grafo, cancion, cancion, n, padres,visitados)
    if ciclo:
        ciclo.append(cancion)
    return ciclo

def _ciclo_n_canciones(grafo, cancion, cancion_actual, n, padres, visitados):
    for adyacente in grafo.adyacentes(cancion_actual):
        if n == 1:
            if adyacente == cancion:
                return auxiliares.reconstruir_camino(padres, cancion, cancion_actual)
        elif adyacente not in visitados:
            visitados.add(adyacente)
            padres[adyacente] = cancion_actual
            ciclo = _ciclo_n_canciones(grafo,cancion,adyacente,n-1,padres,visitados)

            if ciclo:
                return ciclo
            visitados.remove(adyacente)
            del padres[adyacente]
    return None

#Rango
def todas_en_rango(grafo,cancion,n):
    """Devuelve la cantidad de canciones que estan a 'n' vertices de la cancion"""
    cantidad = 0
    visitados = set()
    orden = {}

    cola = collections.deque()
    cola.append(cancion)
    visitados.add(cancion)
    orden[cancion] = 0

    while not len(cola) == 0:
        actual = cola.popleft()
        for adyacente in grafo.adyacentes(actual):
            if adyacente not in visitados:
                visitados.add(adyacente)
                orden[adyacente]= orden[actual] +1
                
                if orden[adyacente]== n:
                    cantidad += 1
                elif orden[adyacente] > n:
                    return cantidad

                cola.append(adyacente)
    
    return cantidad