package votos

import "fmt"

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_VOTACION]string
	resultados [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	resultados [CANT_VOTACION]int
}

func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	return &partidoImplementacion{nombre: nombre, candidatos: candidatos}
}

func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.resultados[tipo]++
}

func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	var votos string
	if partido.resultados[tipo] == 1 {
		votos = "voto"
	} else {
		votos = "votos"
	}
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], partido.resultados[tipo], votos)
}

func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.resultados[tipo]++
}

func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	var votos string
	if blanco.resultados[tipo] == 1 {
		votos = "voto"
	} else {
		votos = "votos"
	}
	return fmt.Sprintf("Votos en Blanco: %d %s", blanco.resultados[tipo], votos)
}
