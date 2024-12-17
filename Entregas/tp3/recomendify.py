#!/usr/bin/python3
from grafo import Grafo
import auxiliares ,biblioteca, salida, sys, heapq

def crear_grafo_conexiones(usuarios):
    """Crea un grafo bipartito en el cual los vertices son los usuarios y las canciones, las aristas son las conexiones entre ellos"""
    grafo = Grafo(False)

    for id_usuario, canciones in usuarios.items():
        if not id_usuario in grafo:
            grafo.agregar_vertice(id_usuario)
        for cancion in canciones.keys():
            if not cancion in grafo:
                grafo.agregar_vertice(cancion)
            grafo.agregar_arista(id_usuario, cancion)
    
    return grafo

def crear_grafo_canciones(usuarios):
    """Crea un grafo en el cual los vertices son las canciones y las aristas son las conexiones entre ellas mediante un usuario"""
    grafo = Grafo(False)

    for id_usuario in usuarios:
        for cancion in usuarios[id_usuario]:
            if not cancion in grafo:
                grafo.agregar_vertice(cancion)
        
        canciones_usuario = list(usuarios[id_usuario].keys())
        for i in range(len(canciones_usuario)):
            for j in range(i+1, len(canciones_usuario)):
                if not grafo.estan_unidos(canciones_usuario[i], canciones_usuario[j]):
                    grafo.agregar_arista(canciones_usuario[i], canciones_usuario[j])

    return grafo

def obtener_data(archivo):
    """Crea un diccionario con la informacion de los usuarios"""
    usuarios = {}
    
    es_primera_linea = True
    for linea in archivo:
        if es_primera_linea:
            es_primera_linea = False
            continue
        id_linea, id_usuario, cancion, artista, playlist_id, playlist, generos = auxiliares.parsear_linea(linea)
        
        usuarios[id_usuario] = usuarios.get(id_usuario, {})

        usuarios[id_usuario][(cancion, artista)] = usuarios[id_usuario].get((cancion, artista), {})
        usuarios[id_usuario][(cancion, artista)][playlist_id] = playlist
    
    return usuarios


def main():
    ruta = sys.argv[1]

    entrada = sys.stdin.readline()
    with open(ruta, "r") as archivo:
        dicc_usuarios = obtener_data(archivo)
        
    grafoCanciones = None
    grafoConexiones = None
    dicc_pagerank = None

    while entrada != "":
        argumentos = entrada.split()
        comando = argumentos[0]
        
        if (comando== "ciclo" or comando== "rango") and grafoCanciones == None:
                grafoCanciones = crear_grafo_canciones(dicc_usuarios)
        elif grafoConexiones == None:
                grafoConexiones = crear_grafo_conexiones(dicc_usuarios)

        if comando =="camino":
            origen = argumentos[1]
            destino = argumentos[2]
            
            canciones_list = auxiliares.separar_listas(argumentos[1:], ">>>>")
            origen_lista, destino_lista = [auxiliares.separar_listas(cancion, "-") for cancion in canciones_list]
            if len(origen_lista) >= 2 and len(destino_lista) >= 2:
                origen = auxiliares.crear_cancion(origen_lista)
                destino = auxiliares.crear_cancion(destino_lista)
                camino = biblioteca.camino_mas_corto(grafoConexiones,origen,destino)
                salida.salida_camino_mas_corto(camino,dicc_usuarios)

            else:
                print("Tanto el origen como el destino deben ser canciones")

        elif comando == "mas_importantes":
            cantidad = int(argumentos[1])
            if dicc_pagerank is None:
                dicc_pagerank = biblioteca.canciones_mas_importantes(grafoConexiones, cantidad)
            
            resultado = heapq.nlargest(cantidad, dicc_pagerank["canciones"].items(), biblioteca.compararPageRank)
            
            salida.salida_canciones(resultado,auxiliares.PUNTO_COMA)

        elif comando == "recomendacion":
            tipo = argumentos[1]
            n = int(argumentos[2])
            canciones_list = auxiliares.separar_listas(argumentos[3:], ">>>>")
            canciones_pre = [auxiliares.separar_listas(cancion, "-") for cancion in canciones_list]
            canciones = [auxiliares.crear_cancion(cancion) for cancion in canciones_pre] 
            
            cancion = auxiliares.separar_listas(argumentos[2:], "-")
            resultado = biblioteca.recomendacion(grafoConexiones, tipo, canciones, n)
            if tipo == "canciones":
                salida.salida_canciones(resultado,auxiliares.PUNTO_COMA)
            else:
                salida.salida_usuarios(resultado)

        elif comando == "ciclo":
            n = int(argumentos[1])
            lista = auxiliares.separar_listas(argumentos[2:], "-")
            cancion = auxiliares.crear_cancion(lista)
            lista = biblioteca.ciclo_n_canciones(grafoCanciones, cancion,n)
            if lista is None:
                print("No se encontro recorrido")
            else:
                salida.salida_canciones(lista,auxiliares.FLECHA)

        elif comando == "rango":
            n = int(argumentos[1])
            lista = auxiliares.separar_listas(argumentos[2:], "-")
            cancion = auxiliares.crear_cancion(lista)
            print(biblioteca.todas_en_rango(grafoCanciones,cancion,n))
            
        else:
            raise Exception("El comando ingresado es invalido")
        
        entrada = sys.stdin.readline()
    return

main()