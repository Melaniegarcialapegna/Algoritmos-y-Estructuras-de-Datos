import biblioteca 
import random

FLECHA = " --> "
MSJ_CANCION_APARECE_PLAYLIST = "aparece en playlist"
MSJ_CONECTOR = "de"
MSJ_PLAYLIST_CONTIENE = "donde aparece"
MSJ_USUARIO_TIENE_PLAYLIST = "tiene una playlist"

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

def canciones_mas_importantes(grafo,n):
    pass

def random_walk(grafo, vertice, probabilidad, probabilidades, largoMax, cantidad):
    if largoMax == 0:
        return
    
    largoMax -= 1
    probabilidades[vertice] = probabilidades.get(vertice, 0)
    probabilidades[vertice] += probabilidad/cantidad
    
    adyacentes = grafo.adyacentes(vertice)
    factor_probabilidad = 1/len(adyacentes)
    probabilidad_sig = probabilidad*factor_probabilidad

    siguiente = random.choice(adyacentes)
    random_walk(grafo, siguiente, probabilidad_sig, probabilidades, largoMax, cantidad)