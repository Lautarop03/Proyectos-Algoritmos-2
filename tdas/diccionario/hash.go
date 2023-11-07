package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	TAMANIO_INICIAL    = 13
	FACTOR_REDIMENSION = 2
	CARGA_MAXIMA       = 3    // a > 3  (a = cant/tam)
	CARGA_MINIMA       = 0.25 // a < 0.25
)

type parClaveValor[K comparable, V any] struct {
	clave K
	dato  V
}

type hashAbierto[K comparable, V any] struct {
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	tam      int
	cantidad int
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashAbierto[K, V]{tabla: crearTabla[K, V](TAMANIO_INICIAL), tam: TAMANIO_INICIAL}
}

func (diccionario *hashAbierto[K, V]) Guardar(clave K, dato V) {
	if (diccionario.cantidad / diccionario.tam) > CARGA_MAXIMA {
		redimension(diccionario.tam*FACTOR_REDIMENSION, diccionario)
	}
	iter := buscarPar(clave, diccionario)
	if iter != nil {
		elem := iter.VerActual()
		elem.dato = dato
	} else {
		diccionario.tabla[f_hash(clave, diccionario.tam)].InsertarUltimo(&parClaveValor[K, V]{clave, dato})
		diccionario.cantidad++
	}
}

func (diccionario *hashAbierto[K, V]) Pertenece(clave K) bool {
	iter := buscarPar(clave, diccionario)
	return iter != nil
}

func (diccionario *hashAbierto[K, V]) Obtener(clave K) V {
	iter := buscarPar(clave, diccionario)
	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}
	return iter.VerActual().dato

}

func (diccionario *hashAbierto[K, V]) Borrar(clave K) V {
	iter := buscarPar(clave, diccionario)
	if iter == nil {
		panic("La clave no pertenece al diccionario")
	}
	elem := iter.VerActual()
	iter.Borrar()
	diccionario.cantidad--
	if float32(diccionario.cantidad/diccionario.tam) < CARGA_MINIMA && (diccionario.tam/FACTOR_REDIMENSION) > TAMANIO_INICIAL {
		redimension(diccionario.tam/FACTOR_REDIMENSION, diccionario)
	}
	return elem.dato
}

func (diccionario *hashAbierto[K, V]) Cantidad() int {
	return diccionario.cantidad
}

func (diccionario *hashAbierto[K, V]) Iterar(f func(clave K, dato V) bool) {
	for _, elem := range diccionario.tabla {
		iter := elem.Iterador()
		for iter.HaySiguiente() {
			if !f(iter.VerActual().clave, iter.VerActual().dato) {
				return
			}
			iter.Siguiente()
		}
	}
}

type iterDiccionario[K comparable, V any] struct {
	actual   *parClaveValor[K, V]
	iterador TDALista.IteradorLista[*parClaveValor[K, V]]
	tabla    []TDALista.Lista[*parClaveValor[K, V]]
	posicion int
}

func (diccionario *hashAbierto[K, V]) Iterador() IterDiccionario[K, V] {
	iter := new(iterDiccionario[K, V])
	iter.tabla = diccionario.tabla
	proximaPosicionValida(iter)
	return iter
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.iterador.VerActual().clave, iter.iterador.VerActual().dato
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.iterador.Siguiente()
	if !iter.iterador.HaySiguiente() {
		iter.posicion += 1
		proximaPosicionValida(iter)
	}
}

// FUNCIONES AUX

func redimension[K comparable, V any](tam int, hash *hashAbierto[K, V]) {
	hashAux := &hashAbierto[K, V]{tabla: crearTabla[K, V](tam), tam: tam}
	hash.Iterar(func(clave K, dato V) bool {
		hashAux.Guardar(clave, dato)
		return true
	})
	hash.tabla = hashAux.tabla
	hash.tam = hashAux.tam
}

func buscarPar[K comparable, V any](clave K, hash *hashAbierto[K, V]) TDALista.IteradorLista[*parClaveValor[K, V]] {
	iter := hash.tabla[f_hash(clave, hash.tam)].Iterador()
	for iter.HaySiguiente() {
		if clave == iter.VerActual().clave {
			return iter
		}
		iter.Siguiente()
	}
	return nil
}

func proximaPosicionValida[K comparable, V any](iter *iterDiccionario[K, V]) {
	for i := iter.posicion; i <= len(iter.tabla); i++ {
		if i == len(iter.tabla) {
			iter.actual = nil
			return
		}
		if iter.tabla[i].EstaVacia() {
			continue
		}
		iterador := iter.tabla[i].Iterador()
		iter.posicion = i
		iter.actual = iterador.VerActual()
		iter.iterador = iterador
		return
	}
}

func crearTabla[K comparable, V any](tam int) []TDALista.Lista[*parClaveValor[K, V]] {
	var tabla []TDALista.Lista[*parClaveValor[K, V]]
	for i := 0; i < tam; i++ {
		lista := TDALista.CrearListaEnlazada[*parClaveValor[K, V]]()
		tabla = append(tabla, lista)
	}
	return tabla
}

// https://golangprojectstructure.com/hash-functions-go-code/
// Uso funcion de hashing de esta pagina
const (
	uint64Offset uint64 = 0xcbf29ce484222325
	uint64Prime  uint64 = 0x00000100000001b3
)

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fvnHash(data []byte) (hash uint64) {
	hash = uint64Offset

	for _, b := range data {
		hash ^= uint64(b)
		hash *= uint64Prime
	}

	return
}

func f_hash[K comparable](clave K, largo int) uint64 {
	return fvnHash(convertirABytes(clave)) % uint64(largo)
}
