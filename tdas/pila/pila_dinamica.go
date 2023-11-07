package pila

/* Definición del struct pila proporcionado por la cátedra. */

const (
	TAMANIO_INICIAL_ARREGLO = 10
	FACTOR_REDIMENSION      = 2
	CARGA_MINIMA            = 4
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, TAMANIO_INICIAL_ARREGLO)}
}

func (pila pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila pilaDinamica[T]) VerTope() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	return pila.datos[pila.cantidad-1]
}

func (pila *pilaDinamica[T]) Apilar(dato T) {
	if pila.cantidad == len(pila.datos) {
		redimensionar(pila, len(pila.datos)*FACTOR_REDIMENSION)
	}
	pila.datos[pila.cantidad] = dato
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if pila.EstaVacia() {
		panic("La pila esta vacia")
	}
	pila.cantidad--
	if pila.cantidad <= len(pila.datos)/CARGA_MINIMA && (len(pila.datos)/FACTOR_REDIMENSION) > TAMANIO_INICIAL_ARREGLO {
		redimensionar(pila, len(pila.datos)/FACTOR_REDIMENSION)
	}
	return pila.datos[pila.cantidad]
}

func redimensionar[T any](pila *pilaDinamica[T], tam_nuevo int) {
	destino := make([]T, tam_nuevo)
	copy(destino, pila.datos)
	pila.datos = destino
}
