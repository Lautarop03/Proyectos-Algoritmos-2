package main

import (
	Aeropuerto "algueiza/aeropuerto"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	aeropuerto := Aeropuerto.CrearAeropuerto()

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		parametros := strings.Split(s.Text(), " ")
		switch parametros[0] {
		case "agregar_archivo":
			if cantidadParametros(parametros, "agregar_archivo", 2) {
				aeropuerto.AgregarArchivo(parametros[1])
			}
		case "ver_tablero":
			if cantidadParametros(parametros, "ver_tablero", 5) {
				aeropuerto.VerTablero(parametros[1], parametros[2], parametros[3], parametros[4])
			}
		case "info_vuelo":
			if cantidadParametros(parametros, "info_vuelo", 2) {
				aeropuerto.InfoVuelo(parametros[1])
			}
		case "prioridad_vuelos":
			if cantidadParametros(parametros, "prioridad_vuelos", 2) {
				aeropuerto.PrioridadVuelos(parametros[1])
			}
		case "siguiente_vuelo":
			if cantidadParametros(parametros, "siguiente_vuelo", 4) {
				aeropuerto.SiguienteVuelo(parametros[1], parametros[2], parametros[3])
			}
		case "borrar":
			if cantidadParametros(parametros, "borrar", 3) {
				aeropuerto.Borrar(parametros[1], parametros[2])
			}
		}
	}
}

func cantidadParametros(param []string, funcion string, cantidad int) bool {
	if len(param) < cantidad {
		fmt.Fprintf(os.Stderr, "Error en comando %s\n", funcion)
		return false
	}
	return true
}
