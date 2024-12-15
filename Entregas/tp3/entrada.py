
import sys,funciones,recomendify

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
    
    for linea in archivo:
        data = parsear_linea(linea)
        id_linea, id_usuario, cancion, artista, playlist_id, playlist, generos = data
        
        usuarios[id_usuario] = usuarios.get(id_usuario, {})

        usuarios[id_usuario][(cancion, artista)] = usuarios[id_usuario].get((cancion, artista), {})
        usuarios[id_usuario][(cancion, artista)][playlist_id] = playlist
    
    return usuarios

def parsear_linea(linea):
    linea = linea.strip()
    return linea.split("\t")

def salida(resultado,separador):
    for i in range(len(resultado)):
        if separador == PUNTO_COMA:
            print(" - ".join(resultado[i][0]), end="")
        else:
            print(" - ".join(resultado[i]), end="")
        if i < len(resultado)-1:
            print(f"{separador} ", end="")



def main():
    ruta = sys.argv[1]

    entrada = input()
    with open(ruta, "r") as archivo:
        dicc_usuarios = obtener_data(archivo)
        
    condicion1 = True
    condicion2 = True

    while entrada != "":
        argumentos = entrada.split()
        comando = argumentos[0]
        
        if comando== "ciclo" or comando== "rango":
            if condicion2:
                grafoCanciones, usuariosCanciones = recomendify.crear_grafo_canciones(ruta)
                condicion2 = False
        else:
            if condicion1:
                grafoConexiones = recomendify.crear_grafo_conexiones(dicc_usuarios)
                condicion1 = False

        if comando =="camino":
            origen = argumentos[1]
            destino = argumentos[2]
            
            canciones_list = obtener_cancion(argumentos[1:], ">>>>")
            print(canciones_list)
            origen, destino = [obtener_cancion(cancion, "-") for cancion in canciones_list]
            origen = (" ".join(origen[0]), " ".join(origen[1]))
            destino = (" ".join(destino[0]), " ".join(destino[1]))

            print(funciones.camino_mas_corto(grafoConexiones,origen,destino, dicc_usuarios))
            
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
            salida(resultado,PUNTO_COMA)

        elif comando == "ciclo":
            n = int(argumentos[1])
            lista = obtener_cancion(argumentos[2:], "-")
            cancion = sacar_cancion(lista)
            lista = funciones.ciclo_n_canciones(grafoCanciones, cancion,n)
            if lista is None:
                print("No se encontro recorrido.")
            print(salida(lista,FLECHA))

        elif comando == "rango":
            n = int(argumentos[1])
            lista = obtener_cancion(argumentos[2:], "-")
            cancion = sacar_cancion(lista)
            print(funciones.todas_en_rango(grafoCanciones,cancion,n))
            
        else:
            raise Exception("El comando ingresado es invalido")
        
        entrada = input()

    return

def sacar_cancion(lista_cancion):
    return (" ".join(lista_cancion[0]), " ".join(lista_cancion[1]))

main()