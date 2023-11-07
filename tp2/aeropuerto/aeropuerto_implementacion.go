package aeropuerto

import (
	"algueiza/vuelo"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	TDAHeap "tdas/cola_prioridad"
	TDADiccionario "tdas/diccionario"
	TDALista "tdas/lista"
)

type aeropuerto struct {
	datosOrdenados TDADiccionario.DiccionarioOrdenado[string, Vuelo.Vuelo]
	datos          TDADiccionario.Diccionario[string, Vuelo.Vuelo]
	conexiones     TDADiccionario.Diccionario[string, TDADiccionario.DiccionarioOrdenado[string, Vuelo.Vuelo]]
}

func CrearAeropuerto() Aeropuerto {
	aeropuerto := new(aeropuerto)
	aeropuerto.datos = TDADiccionario.CrearHash[string, Vuelo.Vuelo]()
	aeropuerto.datosOrdenados = TDADiccionario.CrearABB[string, Vuelo.Vuelo](comparacion)
	aeropuerto.conexiones = TDADiccionario.CrearHash[string, TDADiccionario.DiccionarioOrdenado[string, Vuelo.Vuelo]]()
	return aeropuerto
}

func (aeropuerto aeropuerto) AgregarArchivo(ruta string) {
	listaVuelos, err := cargarArchivo(ruta)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", "Error en comando agregar_archivo\n")
		return
	}
	for _, elem := range listaVuelos {
		vuelo := Vuelo.CrearVuelo(elem)
		conexion := vuelo.VerConexion()
		numVuelo := vuelo.VerNumVuelo()
		fecha := vuelo.VerFecha()
		clave := fechaAComparable(fecha, numVuelo)

		if !aeropuerto.conexiones.Pertenece(conexion) { // Si no existe la conexión, la creo
			arbol := TDADiccionario.CrearABB[string, Vuelo.Vuelo](comparacion)
			aeropuerto.conexiones.Guardar(conexion, arbol)
		}
		if aeropuerto.datos.Pertenece(numVuelo) { // Si existe el vuelo, borro sus datos para luego guardarlos actualizados
			antiguo := aeropuerto.datos.Obtener(numVuelo)
			arbol := aeropuerto.conexiones.Obtener(antiguo.VerConexion())
			aeropuerto.datosOrdenados.Borrar(fechaAComparable(antiguo.VerFecha(), numVuelo))
			arbol.Borrar(fechaAComparable(antiguo.VerFecha(), numVuelo))
		}

		arbol := aeropuerto.conexiones.Obtener(conexion)
		arbol.Guardar(clave, vuelo)
		aeropuerto.datosOrdenados.Guardar(clave, vuelo)
		aeropuerto.datos.Guardar(numVuelo, vuelo)
	}
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

func (aeropuerto aeropuerto) VerTablero(cantidadVuelos, modo, desde, hasta string) {
	cantVuelos, err := strconv.Atoi(cantidadVuelos)
	desdeComp := fechaAComparable(desde, "")
	hastaComp := fechaAComparable(hasta, "")
	if err != nil || cantVuelos < 0 || strings.Compare(desdeComp, hastaComp) > 1 || (modo != "asc" && modo != "desc") {
		fmt.Fprintf(os.Stderr, "%s", "Error en comando ver_tablero\n")
		return
	}
	iterRango := aeropuerto.datosOrdenados.IteradorRango(&desdeComp, &hastaComp)
	if modo == "asc" {
		for i := 0; iterRango.HaySiguiente() && i < cantVuelos; i++ {
			_, vuelo := iterRango.VerActual()
			fmt.Fprintf(os.Stdout, "%s - %s\n", vuelo.VerFecha(), vuelo.VerNumVuelo())
			iterRango.Siguiente()
		}
	} else {
		lista := TDALista.CrearListaEnlazada[Vuelo.Vuelo]()
		for iterRango.HaySiguiente() {
			_, vuelo := iterRango.VerActual()
			lista.InsertarPrimero(vuelo)
			iterRango.Siguiente()
		}
		for i := 0; i < cantVuelos && !lista.EstaVacia(); i++ {
			vuelo := lista.BorrarPrimero()
			fmt.Fprintf(os.Stdout, "%s - %s\n", vuelo.VerFecha(), vuelo.VerNumVuelo())
		}
	}
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

func (aeropuerto aeropuerto) InfoVuelo(NumVuelo string) {
	if !aeropuerto.datos.Pertenece(NumVuelo) {
		fmt.Fprintf(os.Stderr, "%s", "Error en comando info_vuelo\n")
		return
	}
	vuelo := aeropuerto.datos.Obtener(NumVuelo)
	vuelo.VerInfo()
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

// Uso un Algoritmo de tipo Top-K para obtener las PrioridadVuelos
func (aeropuerto aeropuerto) PrioridadVuelos(num string) {
	K, err := strconv.Atoi(num)
	if err != nil || K <= 0 {
		fmt.Fprintf(os.Stderr, "%s", "Error en comando prioridad_vuelos\n")
		return
	}
	var arr []Vuelo.Vuelo
	// Creo un arr de vuelos
	aeropuerto.datos.Iterar(func(_ string, dato Vuelo.Vuelo) bool {
		arr = append(arr, dato)
		return true
	})
	// Creo un heap a partir de Arr ordenado por prioridad
	heap := TDAHeap.CrearHeapArr(arr, func(a, b Vuelo.Vuelo) int {
		if a.VerPrioridad() == b.VerPrioridad() {
			return strings.Compare(b.VerNumVuelo(), a.VerNumVuelo())
		}
		return a.VerPrioridad() - b.VerPrioridad()
	})
	// Desencolo K elementos
	for i := 0; i < K && !heap.EstaVacia(); i++ {
		elem := heap.Desencolar()
		fmt.Fprintf(os.Stdout, "%d - %s\n", elem.VerPrioridad(), elem.VerNumVuelo())
	}
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

func (aeropuerto aeropuerto) SiguienteVuelo(origen, destino, fecha string) {
	conexion := fmt.Sprintf("%s %s", origen, destino)
	arbol := aeropuerto.conexiones.Obtener(conexion)
	fechaComp := fechaAComparable(fecha, "")
	iter := arbol.IteradorRango(&fechaComp, nil)
	if iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		vuelo.VerInfo()
	} else {
		fmt.Fprintf(os.Stdout, "No hay vuelo registrado desde %s hacia %s desde %s\n", origen, destino, fecha)
	}
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

func (aeropuerto aeropuerto) Borrar(desde, hasta string) {
	desdeComp := fechaAComparable(desde, "")
	hastaComp := fechaAComparable(hasta, "")
	if comparacion(desdeComp, hastaComp) > 1 {
		fmt.Fprintf(os.Stderr, "error en comando borrar\n")
		return
	}
	iter := aeropuerto.datosOrdenados.IteradorRango(&desdeComp, &hastaComp)
	lista := TDALista.CrearListaEnlazada[Vuelo.Vuelo]()
	// Guardo en una lista todos mis elementos a borrar
	for iter.HaySiguiente() {
		_, vuelo := iter.VerActual()
		lista.InsertarUltimo(vuelo)
		iter.Siguiente()
	}
	// Recorro los elementos y los borro de mis estructuras de datos
	lista.Iterar(func(vuelo Vuelo.Vuelo) bool {
		clave := fechaAComparable(vuelo.VerFecha(), vuelo.VerNumVuelo())
		vuelo.VerInfo()
		arbol := aeropuerto.conexiones.Obtener(vuelo.VerConexion())
		aeropuerto.datosOrdenados.Borrar(clave)
		arbol.Borrar(fechaAComparable(vuelo.VerFecha(), vuelo.VerNumVuelo()))
		aeropuerto.datos.Borrar(vuelo.VerNumVuelo())
		return true
	})
	fmt.Fprintf(os.Stdout, "%s", "OK\n")
}

// Funciones AUX

func cargarArchivo(ruta string) ([]string, error) {
	var res []string
	archivo, err := os.Open(ruta)

	defer archivo.Close()

	s := bufio.NewScanner(archivo)
	for s.Scan() {
		res = append(res, s.Text())
	}
	err = s.Err()

	return res, err
}

// fechaAInt Convierte una fecha en formato string a un formato adecuado para comparar.
// Recibe una fecha con el formato "2018-10-10T08:51:32" y devuelve "20181010085132-0000".
func fechaAComparable(fecha string, nroVuelo string) string {
	fecha = strings.ReplaceAll(fecha, "T", "")
	fecha = strings.ReplaceAll(fecha, "-", "")
	fecha = strings.ReplaceAll(fecha, ":", "")
	str := fecha + "-" + nroVuelo
	return str
}

// Funcion de comparacion utilizada para los arboles
func comparacion(a, b string) int {
	a1 := strings.Split(a, "-")
	b1 := strings.Split(b, "-")
	if strings.Compare(a1[0], b1[0]) == 0 {
		if a1[1] == "" || b1[1] == "" {
			return strings.Compare(a1[0], b1[0])
		}
		return strings.Compare(a1[1], b1[1])
	}
	return strings.Compare(a1[0], b1[0])
}
