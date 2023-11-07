package mesaelectoral

import (
	"fmt"
	"os"
	"reflect"
	"rerepolez/errores"
	"rerepolez/padron"
	"rerepolez/votos"
	"strconv"
	"strings"
	"tdas/cola"
)

type mesaelectoral struct {
	nombrePartidos     []string
	candidatosPartidos [][]string
	listaPartidos      []votos.Partido
	padron             padron.Padron
	fila               cola.Cola[int]
}

func CrearMesa(ruta string, TDApadron padron.Padron, fila cola.Cola[int]) MesaElectoral {
	nombrePartidos, candidatosPartidos := prepararPartido(padron.CargarArchivo(ruta))
	listaPartidos := crearPartidos(nombrePartidos, candidatosPartidos)
	return mesaelectoral{nombrePartidos, candidatosPartidos, listaPartidos, TDApadron, fila}
}

func (mesa mesaelectoral) Ingresar(dniStr string) {
	dni, err := strconv.Atoi(dniStr)
	if err != nil || dni <= 0 {
		fmt.Fprintf(os.Stdout, "%s\n", errores.DNIError{}.Error())
	} else if !mesa.padron.EstaEnPadron(dni) {
		fmt.Fprintf(os.Stdout, "%s\n", errores.DNIFueraPadron{}.Error())
	} else {
		mesa.fila.Encolar(dni)
		fmt.Fprintf(os.Stdout, "OK\n")
	}
}

func (mesa mesaelectoral) Votar(nLista, candidato string) {
	numeroLista, err := strconv.Atoi(nLista)
	var votante votos.Votante
	if !mesa.fila.EstaVacia() {
		votante = mesa.padron.BuscarVotante(mesa.fila.VerPrimero())
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	}

	if candidato != "Presidente" && candidato != "Gobernador" && candidato != "Intendente" {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorTipoVoto{}.Error())
	} else if err != nil || numeroLista > len(mesa.nombrePartidos) {
		fmt.Fprintf(os.Stdout, "%s\n", errores.ErrorAlternativaInvalida{}.Error())
	} else {
		var tipoVoto votos.TipoVoto
		switch candidato {
		case "Presidente":
			tipoVoto = votos.PRESIDENTE
		case "Intendente":
			tipoVoto = votos.INTENDENTE
		case "Gobernador":
			tipoVoto = votos.GOBERNADOR
		}
		voto := votante.Votar(tipoVoto, numeroLista)
		if voto != nil {
			fmt.Fprintf(os.Stdout, "%s\n", voto.Error())
			mesa.fila.Desencolar()
		} else {
			fmt.Fprintf(os.Stdout, "OK\n")
		}
	}
}

func (mesa mesaelectoral) DesHacer() {
	var votante votos.Votante
	if !mesa.fila.EstaVacia() {
		votante = mesa.padron.BuscarVotante(mesa.fila.VerPrimero())
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	}
	err := votante.Deshacer()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
		if reflect.TypeOf(err) == reflect.TypeOf(errores.ErrorVotanteFraudulento{}) {
			mesa.fila.Desencolar()
		}
	} else {
		fmt.Fprintf(os.Stdout, "OK\n")
	}
}

func (mesa mesaelectoral) FinVotar(impugnados []int) {
	var votante votos.Votante
	if !mesa.fila.EstaVacia() {
		votante = mesa.padron.BuscarVotante(mesa.fila.VerPrimero())
	} else {
		fmt.Fprintf(os.Stdout, "%s\n", errores.FilaVacia{}.Error())
		return
	}
	voto, err := votante.FinVoto()
	if err != nil {
		fmt.Fprintf(os.Stdout, "%s\n", err.Error())
	} else if voto.Impugnado {
		impugnados[0]++
	} else {
		mesa.listaPartidos[voto.VotoPorTipo[0]].VotadoPara(votos.PRESIDENTE)
		mesa.listaPartidos[voto.VotoPorTipo[1]].VotadoPara(votos.GOBERNADOR)
		mesa.listaPartidos[voto.VotoPorTipo[2]].VotadoPara(votos.INTENDENTE)
	}
	fmt.Fprintf(os.Stdout, "OK\n")
	mesa.fila.Desencolar()
}

func prepararPartido(arr []string) ([]string, [][]string) {
	var candidatos [][]string
	var nombres []string
	for _, elemento := range arr {
		split := strings.Split(elemento, ",")
		nombres = append(nombres, split[0])
		candidatos = append(candidatos, split[1:])
	}
	return nombres, candidatos
}

func crearPartidos(nombres []string, candidatos [][]string) []votos.Partido {
	var res []votos.Partido
	var partido votos.Partido
	var candidatosPartido [3]string
	blanco := votos.CrearVotosEnBlanco()
	res = append(res, blanco)
	for i := 0; i < len(nombres); i++ {
		candidatosPartido = [3]string{candidatos[i][0], candidatos[i][1], candidatos[i][2]}
		partido = votos.CrearPartido(nombres[i], candidatosPartido)
		res = append(res, partido)
	}
	return res
}

func (mesa mesaelectoral) ImprimirResultados(impugnados []int) {
	fmt.Printf("Presidente:\n")
	for i := 0; i < len(mesa.listaPartidos); i++ {
		fmt.Println(mesa.listaPartidos[i].ObtenerResultado(votos.PRESIDENTE))
	}
	fmt.Printf("\n")
	fmt.Printf("Gobernador:\n")
	for i := 0; i < len(mesa.listaPartidos); i++ {
		fmt.Println(mesa.listaPartidos[i].ObtenerResultado(votos.GOBERNADOR))
	}
	fmt.Printf("\n")
	fmt.Printf("Intendente:\n")
	for i := 0; i < len(mesa.listaPartidos); i++ {
		fmt.Println(mesa.listaPartidos[i].ObtenerResultado(votos.INTENDENTE))
	}
	fmt.Printf("\n")
	var votos string
	if impugnados[0] == 1 {
		votos = "voto"
	} else {
		votos = "votos"
	}
	fmt.Printf("Votos Impugnados: %d %s\n", impugnados[0], votos)
}
