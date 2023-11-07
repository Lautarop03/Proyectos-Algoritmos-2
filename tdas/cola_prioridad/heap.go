package cola_prioridad

const (
	TAMANIO_INICIAL_ARREGLO = 10
	FACTOR_REDIMENSION      = 2
	CARGA_MINIMA            = 4
)

type heap[T any] struct {
	datos    []T
	cantidad int
	cmp      func(T, T) int
}

func CrearHeap[T any](funcionCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, TAMANIO_INICIAL_ARREGLO)
	heap.cmp = funcionCmp
	return heap
}

func CrearHeapArr[T any](arreglo []T, funcionCmp func(T, T) int) ColaPrioridad[T] {
	heap := new(heap[T])
	heap.datos = make([]T, len(arreglo))
	copy(heap.datos, arreglo)
	heapify(heap.datos, funcionCmp)
	heap.cmp = funcionCmp
	heap.cantidad = len(arreglo)
	return heap
}

func (cola *heap[T]) EstaVacia() bool {
	return cola.cantidad == 0
}

func (cola *heap[T]) Encolar(dato T) {
	if cola.cantidad == len(cola.datos) {
		tam := len(cola.datos)
		if tam == 0 {
			tam = 5
		}
		redimensionar(cola, tam*FACTOR_REDIMENSION)
	}
	cola.datos[cola.cantidad] = dato
	upHeap(cola.datos, cola.cantidad, cola.cmp)
	cola.cantidad++
}

func (cola *heap[T]) VerMax() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	return cola.datos[0]
}

func (cola *heap[T]) Desencolar() T {
	if cola.EstaVacia() {
		panic("La cola esta vacia")
	}
	dato := cola.datos[0]
	cola.cantidad--
	swap(&cola.datos[0], &cola.datos[cola.cantidad])
	downHeap(cola.datos, 0, cola.cantidad, cola.cmp)
	if cola.cantidad <= len(cola.datos)/CARGA_MINIMA && (len(cola.datos)/FACTOR_REDIMENSION) >= TAMANIO_INICIAL_ARREGLO {
		redimensionar(cola, len(cola.datos)/FACTOR_REDIMENSION)
	}
	return dato
}

func (cola *heap[T]) Cantidad() int {
	return cola.cantidad
}

func upHeap[T any](arr []T, posHijo int, cmp func(T, T) int) {
	posPadre := (posHijo - 1) / 2
	if cmp(arr[posHijo], arr[posPadre]) <= 0 {
		return
	}
	swap(&arr[posHijo], &arr[posPadre])
	upHeap(arr, posPadre, cmp)
}

func swap[T any](a, b *T) {
	*a, *b = *b, *a
}

func downHeap[T any](arr []T, posPadre, cantHeap int, cmp func(T, T) int) {
	posHijoIzq := 2*posPadre + 1
	if posHijoIzq >= cantHeap {
		return
	}
	posHijoDer := 2*posPadre + 2
	posMayor := posPadre
	if posHijoIzq < cantHeap && posHijoDer < cantHeap {
		posMayor = max(arr, posHijoDer, posHijoIzq, cmp)
	} else if posHijoDer >= cantHeap && posHijoIzq < cantHeap {
		posMayor = posHijoIzq
	} else if posHijoIzq >= cantHeap && posHijoDer < cantHeap {
		posMayor = posHijoDer
	}
	if cmp(arr[posMayor], arr[posPadre]) <= 0 {
		return
	}
	swap(&arr[posPadre], &arr[posMayor])
	downHeap(arr, posMayor, cantHeap, cmp)
}

func heapify[T any](arr []T, cmp func(T, T) int) {
	for i := len(arr) - 1; i >= 0; i-- {
		downHeap(arr, i, len(arr), cmp)
	}
}

func HeapSort[T any](elementos []T, funcionCmp func(T, T) int) {
	heapify(elementos, funcionCmp)
	heapSort(elementos, funcionCmp)
}
func heapSort[T any](elementos []T, f func(T, T) int) {
	if len(elementos) == 0 {
		return
	}
	ultimo := len(elementos) - 1
	swap(&elementos[0], &elementos[ultimo])
	downHeap(elementos[:ultimo], 0, len(elementos[:ultimo]), f)
	heapSort(elementos[:ultimo], f)
}

func redimensionar[T any](cola *heap[T], tamNuevo int) {
	destino := make([]T, tamNuevo)
	copy(destino, cola.datos)
	cola.datos = destino
}

// func aux downheap
func max[T any](arr []T, posHijoDer, posHijoIzq int, cmp func(T, T) int) int {
	if cmp(arr[posHijoIzq], arr[posHijoDer]) >= 0 {
		return posHijoIzq
	}
	return posHijoDer
}
