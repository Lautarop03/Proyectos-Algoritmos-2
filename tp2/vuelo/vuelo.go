package Vuelo

type Vuelo interface {

	// Devuelve la fecha de vuelo
	VerFecha() string

	// Devuelve el numero del vuelo
	VerNumVuelo() string

	// Imprime toda la informacion sobre el vuelo
	VerInfo()

	// Devuelve la prioridad del vuelo
	VerPrioridad() int

	// Devuelve un string con el origen y destino del vuelo
	VerConexion() string
}
