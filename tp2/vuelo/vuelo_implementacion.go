package Vuelo

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vuelo struct {
	numeroVuelo string
	aeroLinea   string
	origen      string
	destino     string
	numAvion    string
	prioridad   int
	fecha       string
	demora      int
	tiempoVuelo int
	cancelado   int
}

// CrearVuelo Recibe un string con todos los datos de un vuelo y se guardan en el struct
func CrearVuelo(datos string) Vuelo {
	vuelo := new(vuelo)
	listaDatos := strings.Split(datos, ",")
	vuelo.numeroVuelo = listaDatos[0]
	vuelo.aeroLinea = listaDatos[1]
	vuelo.origen = listaDatos[2]
	vuelo.destino = listaDatos[3]
	vuelo.numAvion = listaDatos[4]
	vuelo.prioridad, _ = strconv.Atoi(listaDatos[5])
	vuelo.fecha = listaDatos[6]
	vuelo.demora, _ = strconv.Atoi(listaDatos[7])
	vuelo.tiempoVuelo, _ = strconv.Atoi(listaDatos[8])
	vuelo.cancelado, _ = strconv.Atoi(listaDatos[9])
	return vuelo
}

func (vuelo *vuelo) VerFecha() string {
	return vuelo.fecha
}

func (vuelo *vuelo) VerNumVuelo() string {
	return vuelo.numeroVuelo
}

func (vuelo *vuelo) VerInfo() {
	fmt.Fprintf(os.Stdout, "%s %s %s %s %s %d %s %d %d %d\n",
		vuelo.numeroVuelo, vuelo.aeroLinea,
		vuelo.origen, vuelo.destino,
		vuelo.numAvion, vuelo.prioridad,
		vuelo.fecha, vuelo.demora,
		vuelo.tiempoVuelo, vuelo.cancelado)
}

func (vuelo *vuelo) VerPrioridad() int {
	return vuelo.prioridad
}

func (vuelo *vuelo) VerConexion() string {
	conexion := fmt.Sprintf("%s %s", vuelo.origen, vuelo.destino)
	return conexion
}
