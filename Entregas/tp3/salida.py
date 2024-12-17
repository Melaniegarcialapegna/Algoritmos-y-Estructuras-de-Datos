import auxiliares

FLECHA = " --> "
PUNTO_COMA = ";"
MSJ_CANCION_APARECE_PLAYLIST = "aparece en playlist"
MSJ_CONECTOR = "de"
MSJ_PLAYLIST_CONTIENE = "donde aparece"
MSJ_USUARIO_TIENE_PLAYLIST = "tiene una playlist"

def salida_camino_mas_corto(camino,usuarios):
    """Imprime el camino mas corto en el formato solicitado"""
    if camino is None:
        print("No se encontro recorrido")
        return

    i = 0
    while i < len(camino) - 2:
        cancion = camino[i]
        cancionStr = " - ".join(cancion)
        usuario = camino[i+1]
        cancionSig = camino[i+2]
        cancionSigStr = " - ".join(cancionSig)

        playlistActual = auxiliares.obtener_playlist(usuarios, cancion, usuario)
        playlistSig = auxiliares.obtener_playlist(usuarios, cancionSig, usuario)
        
        print_sin_salto(cancionStr)
        print_sin_salto(FLECHA)
        print_sin_salto(MSJ_CANCION_APARECE_PLAYLIST)
        print_sin_salto(FLECHA)
        print_sin_salto(playlistActual)
        print_sin_salto(FLECHA)
        print_sin_salto(MSJ_CONECTOR)
        print_sin_salto(FLECHA)
        print_sin_salto(usuario)
        print_sin_salto(FLECHA)
        print_sin_salto(MSJ_USUARIO_TIENE_PLAYLIST)
        print_sin_salto(FLECHA)
        print_sin_salto(playlistSig)
        print_sin_salto(FLECHA)
        print_sin_salto(MSJ_PLAYLIST_CONTIENE)
        print_sin_salto(FLECHA)
        
        i+=2
        if i == len(camino) - 1:
            print(cancionSigStr)

def salida_canciones(resultado,separador):
    """Imprime las canciones en el formato solicitado"""
    cadena_separador = f"{separador} " if separador == PUNTO_COMA else f"{separador}"
    for i in range(len(resultado)):
        if separador == PUNTO_COMA:
            print_sin_salto(" - ".join(resultado[i][0]))
        else:
            print_sin_salto(" - ".join(resultado[i]))
        if i < len(resultado)-1:
            print_sin_salto(cadena_separador)
    print(" ")
    return

def salida_usuarios(resultado):
    """Imprime los usuarios en el formato solicitado"""
    for i in range(len(resultado)):
        print_sin_salto(resultado[i][0])
        if i < len(resultado)-1:
            print_sin_salto(f"{PUNTO_COMA} ")

def print_sin_salto(cadena):
    """Imprime la cadena sin salto de linea"""
    print(cadena, end="")