package lista

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados,
	// false en caso contrario
	EstaVacia() bool

	// InsertarPrimero inserta el dato al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacía, entra el pánico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan más
	// elementos, o la función en cuestión devuelva false.
	Iterar(visitar func(T) bool)

	//
	Iterador() IteradorLista[T]
}

type IteradorLista[T any] interface {
	// Devuelve el valor del elemento actual, si ya todos los elementos fueron iterados devuelve un panic "El iterador termino de iterar"
	VerActual() T

	// Devuelve un "True" si hay un nodo siguiente al actual, "False" en caso contrario
	HaySiguiente() bool

	// Avanza el iterador, si ya todos los elementos fueron iterados devuelve un panic "El iterador termino de iterar"
	Siguiente()

	// Inserta un elemento en donde esta parado el iterador
	Insertar(T)

	// Borra el elemento donde esta parado el iterador, avanza el iterador
	// y devuelve el valor borrado. Si ya todos los elementos fueron iterados devuelve un panic "El iterador termino de iterar"
	Borrar() T
}
