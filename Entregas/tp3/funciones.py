import biblioteca , random , heapq, collections

FLECHA = " --> "
MSJ_CANCION_APARECE_PLAYLIST = "aparece en playlist"
MSJ_CONECTOR = "de"
MSJ_PLAYLIST_CONTIENE = "donde aparece"
MSJ_USUARIO_TIENE_PLAYLIST = "tiene una playlist"
CANT_RANDOM_WALKS = 500
CANT_PAGE_RANKS = 200
LARGO_RANDOM_WALK = 100

#CAMINO MAS CORTO
def camino_mas_corto(grafo,origen,destino, usuarios):
    padres, distancia = biblioteca.bfs(grafo, origen)
    camino = biblioteca.reconstruir_camino(padres, origen, destino)
    
    i = 0

    while i < len(camino) - 2:
        cancion = camino[i]
        usuario = camino[i+1]
        cancionSig = camino[i+2]

        playlistActual = obtener_playlist(usuarios, cancion, usuario)
        playlistSig = obtener_playlist(usuarios, cancionSig, usuario)
        
        print(cancion[0])
        print(FLECHA)
        print(MSJ_CANCION_APARECE_PLAYLIST)
        print(FLECHA)
        print(playlistActual)
        print(FLECHA)
        print(MSJ_CONECTOR)
        print(FLECHA)
        print(usuario)
        print(FLECHA)
        print(MSJ_USUARIO_TIENE_PLAYLIST)
        print(FLECHA)
        print(playlistSig)
        print(FLECHA)
        print(MSJ_PLAYLIST_CONTIENE)
        print(FLECHA)
        
        i+=2
        
        if i == len(camino) - 1:
            print(cancionSig[0])
    

def obtener_playlist(usuarios, cancion, usuario):
    return list(usuarios[usuario][cancion].values())[0] #Devuelve alguna playlist del usuario en la que esta la cancion


#RECOMENCACIONES(CANCIONES Y USUARIOS)
def recomendacion(grafo, tipo, vertices, cantidad): #Cambiar lo de las n !! (quedo viejo -> hay que hacer lo de las listas)
    probabilidades = {}
    for vertice in vertices:
        for i in range(len(vertices)):
            random_walk(grafo, vertice, 1, probabilidades, cantidad*cantidad, CANT_RANDOM_WALKS*len(vertices), True)
    
    return heapq.nlargest(cantidad, probabilidades[tipo].items(), compararPageRank)

def compararPageRank(elemento):
    return elemento[1]

def random_walk(grafo, vertice, probabilidad, probabilidades, largoMax, cantidad,primeraIteracion):
    if largoMax == 0:
        return

    largoMax -= 1
    tipo = "canciones" if es_cancion(vertice) else "usuarios"

    
    if not primeraIteracion:
        probabilidades[tipo] = probabilidades.get(tipo, {})
        probabilidades[tipo][vertice] = probabilidades.get(vertice, 0)
        probabilidades[tipo][vertice] += probabilidad/cantidad
    
    adyacentes = grafo.adyacentes(vertice)

    factor_probabilidad = 1/len(adyacentes)
    probabilidad_sig = probabilidad*factor_probabilidad

    siguiente = random.choice(adyacentes)
    random_walk(grafo, siguiente, probabilidad_sig, probabilidades, largoMax, cantidad,False)


#CANCIONES MAS IMPORTANTES
def canciones_mas_importantes(grafo,n):
    pageranks = {}

    vertices = grafo.obtener_vertices()

    for vertice in vertices:
        tipo = "canciones" if es_cancion(vertice) else "usuarios"
        pageranks[tipo] = pageranks.get(tipo, {})
        pageranks[tipo][vertice] = 1/len(vertices)

    for i in range(CANT_PAGE_RANKS):
        visitados = set()
        vertice_aleatorio = grafo.vertice_aleatorio()
        page_rank(grafo, pageranks ,vertice_aleatorio,visitados)

    #Parte de devolver las n mas importantes
    return heapq.nlargest(n, pageranks["canciones"].items(), compararPageRank)

def page_rank(grafo, dicc_pageranks, vertice, visitados):
    visitados.add(vertice)
    tipo_actual = "canciones" if es_cancion(vertice) else "usuarios"
    
    for adyacente in grafo.adyacentes(vertice):
        if adyacente not in visitados:
            tipo_ady = "canciones" if es_cancion(adyacente) else "usuarios"
            pagerank_adyacente = dicc_pageranks[tipo_ady][adyacente] 
            cant_adyacentes = len(grafo.adyacentes(adyacente))
            dicc_pageranks[tipo_actual][vertice] += pagerank_adyacente / cant_adyacentes
            page_rank(grafo, dicc_pageranks, adyacente, visitados)

def es_cancion(vertice):
    return type(vertice) == tuple

def ciclo_n_canciones(grafo, cancion, n):
    padres = {}
    padres[cancion]= None
    visitados = set()
    visitados.add(cancion)

    ciclo = _ciclo_n_canciones(grafo, cancion, cancion, n, padres,visitados)
    if ciclo is not None:
        ciclo.append(cancion)
    return ciclo

#CICLO DE N CANCIONES
def _ciclo_n_canciones(grafo, cancion, cancion_actual, n, padres, visitados):
    for adyacente in grafo.adyacentes(cancion_actual):
        if n == 1:
            if adyacente == cancion:
                return biblioteca.reconstruir_camino(padres, cancion, cancion_actual)
        elif adyacente not in visitados:
            visitados.add(adyacente)
            padres[adyacente] = cancion_actual
            ciclo = _ciclo_n_canciones(grafo,cancion,adyacente,n-1,padres,visitados)

            if ciclo is not None:
                return ciclo

    visitados.remove(cancion_actual)
    return None


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