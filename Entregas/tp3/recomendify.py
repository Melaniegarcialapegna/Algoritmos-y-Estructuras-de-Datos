from grafo import Grafo
import funciones , random

def crear_grafo_conexiones(usuarios):
    grafo = Grafo(False)

    for id_usuario, canciones in usuarios.items():
        if not id_usuario in grafo:
            grafo.agregar_vertice(id_usuario)
        for cancion in canciones.keys():
            if not cancion in grafo:
                grafo.agregar_vertice(cancion)
            grafo.agregar_arista(id_usuario, cancion)
    
    return grafo


    # usuarios = {}

    # with open(ruta, "r") as archivo:
    #     for linea in archivo:
    #         data = parsear_linea(linea)
    #         id_linea, id_usuario, cancion, artista, playlist_id, playlist, generos = data
            
    #         usuarios[id_usuario] = usuarios.get(id_usuario, {})

    #         usuarios[id_usuario][(cancion, artista)] = usuarios[id_usuario].get((cancion, artista), {})
    #         usuarios[id_usuario][(cancion, artista)][playlist_id] = playlist

    #         if not (cancion, artista) in grafo:
    #             grafo.agregar_vertice((cancion,artista))
    #         if not id_usuario in grafo:
    #             grafo.agregar_vertice(id_usuario)
    #         grafo.agregar_arista((cancion, artista), id_usuario)
    
    return grafo, usuarios

def crear_grafo_canciones(ruta):
    grafo = Grafo(False)
    usuarios = {}
    
    #Esto se hace de vuelta arriba. Poner en una funcion tipo obtenerData
    with open(ruta, "r") as archivo:
        for linea in archivo:
            data = parsear_linea(linea)
            id_linea, id_usuario, cancion, artista, playlist_id, playlist, generos = data
            
            usuarios[id_usuario] = usuarios.get(id_usuario, {})
            usuarios[id_usuario][(cancion, artista)] = usuarios[id_usuario].get((cancion, artista), {})
            usuarios[id_usuario][(cancion, artista)][playlist_id] = playlist
        
            if not (cancion, artista) in grafo:
                grafo.agregar_vertice((cancion,artista))
            
            for cancionUsuario in usuarios[id_usuario].keys():
                grafo.agregar_arista((cancion, artista), cancionUsuario)

    return grafo, usuarios

def parsear_linea(linea):
    linea = linea.strip()
    return linea.split("\t")

