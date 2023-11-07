package diccionario

import (
	TDAPila "tdas/pila"
)

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type iterDiccionarioOrdenado[K comparable, V any] struct {
	pila  TDAPila.Pila[nodoAbb[K, V]]
	abb   abb[K, V]
	hasta *K
	desde *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave: clave, dato: dato}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	if abb.raiz == nil { // si el arbol no existe defino la raiz
		abb.raiz = crearNodo(clave, dato)
		abb.cantidad++
	} else {
		elem, padre := abb.buscarElemento(clave)
		if elem != nil { // si es true existe, reemplazo el dato
			elem.dato = dato
		} else { // no existe lo tengo que guardar
			valor := abb.cmp(clave, padre.clave)
			if valor < 0 {
				padre.izquierdo = crearNodo(clave, dato)
			} else {
				padre.derecho = crearNodo(clave, dato)
			}
			abb.cantidad++
		}
	}
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	elem, _ := abb.buscarElemento(clave)
	return elem != nil
}

func (abb *abb[K, V]) Obtener(clave K) V {
	elem, _ := abb.buscarElemento(clave)
	if elem == nil {
		panic("La clave no pertenece al diccionario")
	}
	return elem.dato
}

func (abb *abb[K, V]) Borrar(clave K) V {
	elem, padre := abb.buscarElemento(clave)
	var valor int
	if elem == nil {
		panic("La clave no pertenece al diccionario")
	}
	if padre != nil {
		valor = abb.cmp(elem.clave, padre.clave)
	}
	dato := elem.dato

	if elem.izquierdo == nil || elem.derecho == nil { // caso 1 y 2 (sin hijos),(1 solo hijo)
		var hijo *nodoAbb[K, V]

		if elem.izquierdo != nil {
			hijo = elem.izquierdo
		} else if elem.derecho != nil {
			hijo = elem.derecho
		}

		if padre == nil { // borro la raiz
			abb.raiz = hijo
		} else if valor < 0 {
			padre.izquierdo = hijo
		} else {
			padre.derecho = hijo
		}
		abb.cantidad--
	} else { // caso 3 (dos hijos)
		reemplazo := abb.buscarReemplazo(elem.derecho, nil, "", nil)
		abb.Borrar(reemplazo.clave)
		elem.clave, elem.dato = reemplazo.clave, reemplazo.dato
	}
	return dato
}

func (abb *abb[K, V]) buscarReemplazo(actual *nodoAbb[K, V], padre *nodoAbb[K, V], accion string, pila TDAPila.Pila[nodoAbb[K, V]]) *nodoAbb[K, V] {
	if actual == nil {
		return padre
	}
	if accion == "apilar" {
		pila.Apilar(*actual)
	}
	return abb.buscarReemplazo(actual.izquierdo, actual, accion, pila)
}

func (abb *abb[K, V]) buscarElemento(clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	return abb.buscarElementoRecursivo(abb.raiz, nil, clave)
}

func (abb *abb[K, V]) buscarElementoRecursivo(actual *nodoAbb[K, V], padre *nodoAbb[K, V], clave K) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if actual == nil { // No se encontró el elemento, se devuelve el padre
		return nil, padre
	}
	valor := abb.cmp(clave, actual.clave)
	if valor == 0 { // Se encontró el elemento, se devuelve el elemento y su padre
		return actual, padre
	} else if valor < 0 {
		return abb.buscarElementoRecursivo(actual.izquierdo, actual, clave)
	} else {
		return abb.buscarElementoRecursivo(actual.derecho, actual, clave)
	}
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.iterar(abb.raiz, f, nil, nil)
}

func (abb *abb[K, V]) iterar(nodo *nodoAbb[K, V], f func(clave K, dato V) bool, desde, hasta *K) bool {
	if nodo == nil {
		return true
	}

	if desde == nil || abb.cmp(nodo.clave, *desde) >= 0 {
		if !abb.iterar(nodo.izquierdo, f, desde, hasta) {
			return false
		}
	}
	if (desde == nil || abb.cmp(nodo.clave, *desde) >= 0) && (hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0) {
		if !f(nodo.clave, nodo.dato) {
			return false
		}
	}
	if hasta == nil || abb.cmp(nodo.clave, *hasta) <= 0 {
		if !abb.iterar(nodo.derecho, f, desde, hasta) {
			return false
		}
	}
	return true
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	abb.buscarReemplazo(abb.raiz, nil, "apilar", pila)
	return &iterDiccionarioOrdenado[K, V]{pila: pila, abb: *abb}
}

func (iter *iterDiccionarioOrdenado[K, V]) HaySiguiente() bool {
	var valor int
	if !iter.pila.EstaVacia() {
		if iter.hasta != nil {
			valor = iter.abb.cmp(iter.pila.VerTope().clave, *iter.hasta)
		}
	}
	return !iter.pila.EstaVacia() && (valor <= 0)
}

func (iter *iterDiccionarioOrdenado[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elem := iter.pila.VerTope()
	return elem.clave, elem.dato
}

func (iter *iterDiccionarioOrdenado[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elem := iter.pila.Desapilar()
	if elem.derecho != nil {
		iter.abb.buscarReemplazo(elem.derecho, nil, "apilar", iter.pila)
	}
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.iterar(abb.raiz, visitar, desde, hasta)
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[nodoAbb[K, V]]()
	if desde == nil {
		abb.buscarReemplazo(abb.raiz, nil, "apilar", pila)
	} else {
		abb.buscar(abb.raiz, pila, desde)
	}
	return &iterDiccionarioOrdenado[K, V]{pila, *abb, hasta, desde}
}
func (abb *abb[K, V]) buscar(nodo *nodoAbb[K, V], pila TDAPila.Pila[nodoAbb[K, V]], desde *K) {
	if nodo == nil {
		return
	}
	valor := abb.cmp(nodo.clave, *desde)
	if valor == 0 {
		pila.Apilar(*nodo)
		return
	} else if valor < 0 {
		abb.buscar(nodo.derecho, pila, desde)
	} else {
		pila.Apilar(*nodo)
		abb.buscar(nodo.izquierdo, pila, desde)
	}
}
