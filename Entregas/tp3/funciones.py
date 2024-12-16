import biblioteca , random , heapq, collections

FLECHA = " --> "
MSJ_CANCION_APARECE_PLAYLIST = "aparece en playlist"
MSJ_CONECTOR = "de"
MSJ_PLAYLIST_CONTIENE = "donde aparece"
MSJ_USUARIO_TIENE_PLAYLIST = "tiene una playlist"
CANT_RANDOM_WALKS = 500
CANT_PAGE_RANKS = 300
LARGO_RANDOM_WALK = 100
COEFICIENTE_AMORTIGUACION = 0.85

#CAMINO MAS CORTO
def camino_mas_corto(grafo,origen,destino, usuarios):
    padres, distancia = biblioteca.bfs(grafo, origen)
    
    if not destino in distancia:
        print("No se encontro recorrido")
        return 
    
    camino = biblioteca.reconstruir_camino(padres, origen, destino)

    i = 0

    while i < len(camino) - 2:
        cancion = camino[i]
        cancionStr = " - ".join(cancion)
        usuario = camino[i+1]
        
        cancionSig = camino[i+2]
        cancionSigStr = " - ".join(cancionSig)

        playlistActual = obtener_playlist(usuarios, cancion, usuario)
        playlistSig = obtener_playlist(usuarios, cancionSig, usuario)
        
        print(cancionStr, end="")
        print(FLECHA, end="")
        print(MSJ_CANCION_APARECE_PLAYLIST, end="")
        print(FLECHA, end="")
        print(playlistActual, end="")
        print(FLECHA, end="")
        print(MSJ_CONECTOR, end="")
        print(FLECHA, end="")
        print(usuario, end="")
        print(FLECHA, end="")
        print(MSJ_USUARIO_TIENE_PLAYLIST, end="")
        print(FLECHA, end="")
        print(playlistSig, end="")
        print(FLECHA, end="")
        print(MSJ_PLAYLIST_CONTIENE, end="")
        print(FLECHA, end="")
        
        i+=2
        
        if i == len(camino) - 1:
            print(cancionSigStr)
    

def obtener_playlist(usuarios, cancion, usuario):
    return list(usuarios[usuario][cancion].values())[0] #Devuelve alguna playlist del usuario en la que esta la cancion


#RECOMENCACIONES(CANCIONES Y USUARIOS)
def recomendacion(grafo, tipo, vertices, cantidad): 
    probabilidades = {}

    for vertice in grafo.obtener_vertices():
        tipo_actual = "canciones" if es_cancion(vertice) else "usuarios"
        probabilidades[tipo_actual] = probabilidades.get(tipo_actual, {})
        probabilidades[tipo_actual][vertice] = 1

    for vertice in vertices:
        for i in range(500):
            random_walk(grafo, vertice, 1, probabilidades, 200, True)
    
    return heapq.nlargest(cantidad, probabilidades[tipo].items(), compararPageRank)


def compararPageRank(elemento):
    return elemento[1]

def random_walk(grafo, vertice, probabilidad, probabilidades, largoMax,primeraIteracion):
    if largoMax == 0:
        return

    largoMax -= 1
    tipo = "canciones" if es_cancion(vertice) else "usuarios"
    
    if not primeraIteracion:
        probabilidades[tipo] = probabilidades.get(tipo, {})
        probabilidades[tipo][vertice] += probabilidad
    
    adyacentes = grafo.adyacentes(vertice)
    probabilidad_sig = probabilidad/len(adyacentes)

    siguiente = random.choice(adyacentes)
    random_walk(grafo, siguiente, probabilidad_sig, probabilidades, largoMax,False)


#CANCIONES MAS IMPORTANTES
def canciones_mas_importantes(grafo,n):
    pageranks = {}

    vertices = grafo.obtener_vertices()

    for vertice in vertices:
        tipo = "canciones" if es_cancion(vertice) else "usuarios"
        if tipo not in pageranks:
            pageranks[tipo] = {}
        pageranks[tipo][vertice] = (1-COEFICIENTE_AMORTIGUACION)/len(vertices)

    for i in range(CANT_PAGE_RANKS):
        page_rank(grafo, pageranks)

    #Parte de devolver las n mas importantes
    return heapq.nlargest(n, pageranks["canciones"].items(), compararPageRank)

def page_rank(grafo, dicc_pageranks):
    
    for vertice in grafo.obtener_vertices():
        tipo_actual = "canciones" if es_cancion(vertice) else "usuarios"
        if tipo_actual == "canciones":
            pageRank_articulo = dicc_pageranks[tipo_actual][vertice]
            for adyacente in grafo.adyacentes(vertice):
                tipo_ady = "canciones" if es_cancion(adyacente) else "usuarios"
                pagerank_adyacente = dicc_pageranks[tipo_ady].get(adyacente,0)
                cant_adyacentes = len(grafo.adyacentes(adyacente))
                pageRank_articulo += COEFICIENTE_AMORTIGUACION * (pagerank_adyacente / cant_adyacentes) 
            dicc_pageranks[tipo_actual][vertice] = pageRank_articulo

def es_cancion(vertice):
    return type(vertice) == tuple

def ciclo_n_canciones(grafo, cancion, n):
    '''Esta funcion se encarga de buscar un ciclo de n vertices que comience y termine en el vertice pasado por parametro'''
    camino = []
    camino.append(cancion)
    return backtracking(grafo, cancion, cancion, n, camino)

#CICLO DE N CANCIONES
def backtracking(grafo, origen, vertice, n, camino):
    '''Algoritmo llamado en buscar_ciclo'''
    if((len(camino) == n+1) and (vertice == origen)):
        return True, camino
    for adyacente in grafo.adyacentes(vertice):
        if (adyacente != origen):
            if adyacente in camino:
                continue
        if (adyacente == origen):
            if len(camino) != n:
                continue
        if len(camino) > n:
            return False, camino
        camino.append(adyacente)
        booleano, camino = backtracking(grafo, origen, adyacente, n, camino)
        if booleano == False: 
            camino.remove(adyacente)
        else:
            return True, camino
    return False, camino

#TODAS EN RANGO
def todas_en_rango(grafo,cancion,n):
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