package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	errores "rerepolez/errores"
	TDAMesa "rerepolez/mesaelectoral"
	TDAPadron "rerepolez/padron"
	TDACola "tdas/cola"
)

func main() {

	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorParametros{}.Error())
		os.Exit(1)
	}

	rutaLista := os.Args[1]
	rutaPadron := os.Args[2]

	fila := TDACola.CrearColaEnlazada[int]()

	padron := TDAPadron.CrearPadron(rutaPadron)
	mesaelectoral := TDAMesa.CrearMesa(rutaLista, padron, fila)
	impugnados := make([]int, 1)

	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		if s.Text() == "" {
			break
		}
		parametros := strings.Split(s.Text(), " ")
		switch parametros[0] {
		case "ingresar":
			mesaelectoral.Ingresar(parametros[1])
		case "votar":
			mesaelectoral.Votar(parametros[2], parametros[1])
		case "deshacer":
			mesaelectoral.DesHacer()
		case "fin-votar":
			mesaelectoral.FinVotar(impugnados)
		}
	}

	if !fila.EstaVacia() {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorCiudadanosSinVotar{}.Error())
	}

	mesaelectoral.ImprimirResultados(impugnados)
}
