package aeropuerto

type Aeropuerto interface {
	AgregarArchivo(ruta string)

	// Muestra los K vuelos ordenados por fecha de forma ascendente o descendente.
	VerTablero(cantidadVuelos, modo, desde, hasta string)

	// Muestra toda la información posible sobre el vuelo que tiene el código pasado por parámetro.
	InfoVuelo(numeroVuelo string)

	// Muestra los codigos de los K vuelos con mayor prioridad
	PrioridadVuelos(K string)

	// Muestra la información del vuelo (tal cual en info_vuelo) del próximo vuelo directo que
	//conecte los aeropuertos de origen y destino, a partir de la fecha indicada (inclusive).
	SiguienteVuelo(origen, destino, fecha string)

	// Borra todos los vuelos cuya fecha de despegue estén dentro del intervalo <desde> <hasta>
	Borrar(desde, hasta string)
}
