package padron

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/errores"
	ordenamientos "rerepolez/ordenamientos"
	"rerepolez/votos"
	"strconv"
)

type padron struct {
	dniValidos    []int
	listaVotantes []votos.Votante
}

func CrearPadron(ruta string) Padron {
	arrDni := stringAInt(CargarArchivo(ruta))
	dniOrdenados := ordenamientos.RadixSort(arrDni)
	listaVotantes := crearVotantes(dniOrdenados)
	return padron{dniOrdenados, listaVotantes}
}

func (p padron) BuscarVotante(dni int) votos.Votante {
	return ordenamientos.BusquedaBinaria(dni, p.listaVotantes)
}

func (p padron) EstaEnPadron(dni int) bool {
	votante := ordenamientos.BusquedaBinaria(dni, p.listaVotantes)
	return votante != nil
}

func CargarArchivo(ruta string) []string {
	var res []string
	archivo, err := os.Open(ruta)
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorLeerArchivo{}.Error())
		os.Exit(1)
	}
	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		res = append(res, s.Text())
	}
	err = s.Err()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorLeerArchivo{}.Error())
		os.Exit(1)
	}
	return res
}

func stringAInt(arr []string) []int {
	res := make([]int, 0)
	for i := 0; i < len(arr); i++ {
		num, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorLeerArchivo{}.Error())
		} else {
			res = append(res, num)
		}
	}
	return res
}

func crearVotantes(slice []int) []votos.Votante {
	var votantes []votos.Votante
	for i := 0; i < len(slice); i++ {
		votante := votos.CrearVotante(slice[i])
		votantes = append(votantes, votante)
	}
	return votantes
}
