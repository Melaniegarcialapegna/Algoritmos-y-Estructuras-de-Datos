
import sys,funciones,recomendify,biblioteca

def main():
    ruta = sys.argv[1]

    entrada = input()
    with open(ruta, "r") as archivo:
        dicc_usuarios = obtener_data(archivo)
        
        grafo , usuarios = recomendify.crear_grafo_canciones
        grafo , usuarios = recomendify.crear_grafo_conexiones


    while entrada != "":
        argumentos = entrada.split()
        comando = argumentos[0]
        
        if comando =="camino":
            origen = argumentos[1]
            destino = argumentos[2]
            funciones.camino_mas_corto(grafo,origen,destino, usuarios)
            
        elif comando == "mas_importantes":
            cantidad = argumentos[1]
            funciones.canciones_mas_importantes(grafo, cantidad)

        elif comando == "recomendacion":
            tipo = argumentos[1]
            n = argumentos[2]
            canciones = obtener_cancion(argumentos[2:], ">>>>")
            funciones.recomendacion(grafo, tipo, canciones, n)

        elif comando == "ciclo":
            cantidad = argumentos[1]
            cancion = obtener_cancion(argumentos[2:])
            print(funciones.ciclo_n_canciones(grafo, cancion))
            entrada = input()

        elif comando == "rango":
            n = argumentos[1]
            cancion = tuple(obtener_cancion(argumentos[2:], "-"))
            funciones.todas_en_rango(grafo,cancion,n)
            
        else:
            raise Exception("El comando ingresado es invalido")


    return


main()
    
        
def obtener_cancion(argumentos, separador):
    resultado = []
    actual = []

    i = 0
    while argumentos[i] != separador:
        actual.append(argumentos[i])
        i += 1

    actual_str = " ".join(actual)
    resultado.append(actual_str)
    i += 1
    actual = []

    while i < len(argumentos):
        resultado.append(argumentos[i])
        i += 1

    actual = " ".join(actual)
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