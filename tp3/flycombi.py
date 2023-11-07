#!/usr/bin/python3

import sys
from grafo import grafo
import biblioteca

COMANDO = 0
PARAMETRO = 1


def main():
    g = grafo(False)
    ciudades = {}  # ciudad : {codigo_aeropuerto}
    cantidadVuelos = {}  # aeropuerto : {destino: cantVuelos}
    coordenadas = {}  # codigo_aeropuerto : ((latitud,longitud),ciudad)
    biblioteca.cargarAeropuertos(sys.argv[1], g, ciudades, coordenadas)
    biblioteca.cargarVuelos(sys.argv[2], g, cantidadVuelos)
    ultima_ruta = []

    try:
        while True:
            linea = input()
            if linea == "":
                break

            parametros = linea.split(" ", maxsplit=1)
            if parametros[COMANDO] == "camino_mas":
                modo, origen, destino = parametros[PARAMETRO].split(",")
                if not biblioteca.perteneneCiudad(ciudades, origen, destino): continue
                ultima_ruta = biblioteca.minimos_generalizado(g, origen, destino, ciudades, modo, parametros[COMANDO])
            if parametros[COMANDO] == "camino_escalas":
                origen, destino = parametros[PARAMETRO].split(",")
                if not biblioteca.perteneneCiudad(ciudades, origen, destino): continue
                ultima_ruta = biblioteca.minimos_generalizado(g, origen, destino, ciudades, 0, parametros[COMANDO])
            if parametros[COMANDO] == "centralidad":
                biblioteca.mas_central(g, int(parametros[PARAMETRO]), cantidadVuelos)
            if parametros[COMANDO] == "nueva_aerolinea":
                arbol = biblioteca.mst_prim(g)
                biblioteca.nuevo_aeropuerto(arbol, parametros[PARAMETRO], cantidadVuelos)
            if parametros[COMANDO] == "itinerario":
                itinerarios = biblioteca.cargarItinerario(parametros[PARAMETRO])
                orden = biblioteca.topologia(itinerarios)
                biblioteca.camino_minimo_por_orden(orden, g, ciudades)
            if parametros[COMANDO] == "exportar_kml":
                biblioteca.exportar_KML(parametros[PARAMETRO], ultima_ruta, coordenadas)

    except EOFError:
        pass


if __name__ == "__main__":
    main()
