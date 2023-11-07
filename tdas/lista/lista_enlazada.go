package lista

type nodoLista[T any] struct {
	dato T
	prox *nodoLista[T]
}

func nodoCrear[T any](valor T) *nodoLista[T] {
	return &nodoLista[T]{dato: valor}
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

func (lista *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := nodoCrear(dato)
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	} else {
		nuevo.prox = lista.primero
	}
	lista.primero = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := nodoCrear(dato)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.prox = nuevo
	}
	lista.ultimo = nuevo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	dato := lista.primero.dato
	lista.primero = lista.primero.prox
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	act := lista.primero
	for act != nil && visitar(act.dato) {
		act = act.prox
	}
}

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, lista: lista}
}

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
}

func (iter *iterListaEnlazada[T]) VerActual() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.dato
}

func (iter *iterListaEnlazada[T]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterListaEnlazada[T]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.anterior = iter.actual
	iter.actual = iter.actual.prox
}

func (iter *iterListaEnlazada[T]) Borrar() T {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	siguiente := iter.actual.prox
	dato := iter.actual.dato
	if iter.anterior != nil {
		iter.anterior.prox = siguiente
	}
	iter.actual = siguiente
	if iter.anterior == nil {
		iter.lista.primero = iter.actual
	}
	if iter.actual == nil {
		iter.lista.ultimo = iter.anterior
	}
	iter.lista.largo--
	return dato
}

func (iter *iterListaEnlazada[T]) Insertar(dato T) {
	nuevo := nodoCrear(dato)
	if iter.anterior == nil && iter.actual == nil {
		iter.lista.primero = nuevo
		iter.lista.ultimo = nuevo
	} else if iter.anterior == nil {
		nuevo.prox = iter.actual
		iter.lista.primero = nuevo
	} else if iter.actual == nil {
		iter.anterior.prox = nuevo
		iter.lista.ultimo = nuevo
	} else {
		iter.anterior.prox = nuevo
		nuevo.prox = iter.actual
	}
	iter.actual = nuevo
	iter.lista.largo++
}
