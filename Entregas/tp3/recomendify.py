#!/usr/bin/python3
from grafo import Grafo
import funciones , sys, csv

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

FLECHA = " --> "
PUNTO_COMA = ";"

def obtener_cancion(argumentos, separador):
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
        
def obtener_canciones(argumentos, separador):
    resultado = []
    actual = []

    i = 0
    while argumentos[i] != separador:
        actual.append(argumentos[i])
        i += 1
    
    resultado.append(actual)
    i += 1
    actual = []

    while i < len(argumentos):
        actual

def verificar_entrada():
    pass

def obtener_data(archivo):
    usuarios = {}
    
    es_primera_linea = True
    for linea in archivo:
        if es_primera_linea:
            es_primera_linea = False
            continue
        id_linea, id_usuario, cancion, artista, playlist_id, playlist, generos = parsear_linea(linea)
        
        usuarios[id_usuario] = usuarios.get(id_usuario, {})

        usuarios[id_usuario][(cancion, artista)] = usuarios[id_usuario].get((cancion, artista), {})
        usuarios[id_usuario][(cancion, artista)][playlist_id] = playlist
    
    return usuarios



def salida(resultado,separador):
    for i in range(len(resultado)):
        if separador == PUNTO_COMA:
            print(" - ".join(resultado[i][0]), end="")
        else:
            print(" - ".join(resultado[i]), end="")
        if i < len(resultado)-1:
            print(f"{separador} ", end="")
    print("", end="\n")
    return



def main():
    ruta = sys.argv[1]

    entrada = sys.stdin.readline()
    with open(ruta, "r") as archivo:
        dicc_usuarios = obtener_data(archivo)
        
    condicion1 = True
    condicion2 = True

    while entrada != "":
        argumentos = entrada.split()
        comando = argumentos[0]
        
        if comando== "ciclo" or comando== "rango":
            if condicion2:
                grafoCanciones, usuariosCanciones = crear_grafo_canciones(ruta)
                condicion2 = False
        else:
            if condicion1:
                grafoConexiones = crear_grafo_conexiones(dicc_usuarios)
                condicion1 = False

        if comando =="camino":
            origen = argumentos[1]
            destino = argumentos[2]
            
            canciones_list = obtener_cancion(argumentos[1:], ">>>>")
            origen_lista, destino_lista = [obtener_cancion(cancion, "-") for cancion in canciones_list]
            if len(origen_lista) >= 2 and len(destino_lista) >= 2:
                origen = sacar_cancion(origen_lista)
                destino = sacar_cancion(destino_lista)
                funciones.camino_mas_corto(grafoConexiones,origen,destino, dicc_usuarios)
            else:
                print("Tanto el origen como el destino deben ser canciones")

        elif comando == "mas_importantes":
            cantidad = int(argumentos[1])
            resultado = funciones.canciones_mas_importantes(grafoConexiones, cantidad)
            #Devuelve una lista de tuplas
            salida(resultado,PUNTO_COMA)

        elif comando == "recomendacion":
            tipo = argumentos[1]
            n = int(argumentos[2])
            canciones_list = obtener_cancion(argumentos[3:], ">>>>")
            canciones_pre = [obtener_cancion(cancion, "-") for cancion in canciones_list]
            canciones = [sacar_cancion(cancion) for cancion in canciones_pre] 
            
            cancion = obtener_cancion(argumentos[2:], "-")
            resultado = funciones.recomendacion(grafoConexiones, tipo, canciones, n)
            if tipo == "canciones":
                salida(resultado,PUNTO_COMA)
            else:
                for i in range(len(resultado)):
                    print(resultado[i][0], end="")
                    if i < len(resultado)-1:
                        print("; ", end="")

        elif comando == "ciclo":
            n = int(argumentos[1])
            lista = obtener_cancion(argumentos[2:], "-")
            cancion = sacar_cancion(lista)
            lista = funciones.ciclo_n_canciones(grafoCanciones, cancion,n)
            if lista is None:
                print("No se encontro recorrido.")
            else:
                salida(lista,FLECHA)

        elif comando == "rango":
            n = int(argumentos[1])
            lista = obtener_cancion(argumentos[2:], "-")
            cancion = sacar_cancion(lista)
            print(funciones.todas_en_rango(grafoCanciones,cancion,n))
            
        else:
            raise Exception("El comando ingresado es invalido")
        
        entrada = sys.stdin.readline()

    return

def sacar_cancion(lista_cancion):
    return (" ".join(lista_cancion[0]), " ".join(lista_cancion[1]))

main()