package lista_test

import (
	TDALista "tdas/lista"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestInsertarBorrar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 1, lista.VerPrimero())
	require.EqualValues(t, 1, lista.VerUltimo())
	require.EqualValues(t, 1, lista.Largo())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestInvarianteDePila(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	lista.InsertarPrimero(1)
	lista.InsertarUltimo(2)
	lista.InsertarPrimero(3)
	lista.InsertarUltimo(4)
	lista.InsertarPrimero(5)
	require.EqualValues(t, 5, lista.BorrarPrimero())
	require.EqualValues(t, 3, lista.BorrarPrimero())
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
}

func TestPruebaDeVolumen(t *testing.T) {
	tam := 10000
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= tam; i++ {
		lista.InsertarUltimo(i)
	}
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, lista.VerPrimero())
		require.EqualValues(t, i, lista.BorrarPrimero())
		require.EqualValues(t, tam-i, lista.Largo())
		require.EqualValues(t, i+1, lista.VerPrimero())
	}
	require.EqualValues(t, 10000, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())

}

func TestInsertarBorrarListaNueva(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		lista.InsertarPrimero(i)
	}
	for i := 10; i >= 0; i-- {
		lista.BorrarPrimero()
	}
	require.True(t, lista.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.EqualValues(t, 0, lista.Largo())
}

func TestListaStr(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	lista.InsertarPrimero("hola")
	lista.InsertarPrimero("Vaso")
	lista.InsertarPrimero("reloj")
	require.False(t, lista.EstaVacia())
	require.EqualValues(t, 3, lista.Largo())
	require.EqualValues(t, "reloj", lista.BorrarPrimero())
	require.EqualValues(t, "Vaso", lista.BorrarPrimero())
	require.EqualValues(t, "hola", lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestListaBool(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[bool]()
	lista.InsertarPrimero(true)
	lista.InsertarPrimero(false)
	require.EqualValues(t, false, lista.BorrarPrimero())
	require.EqualValues(t, true, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

func TestColaFloat(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[float32]()
	lista.InsertarPrimero(133.15)
	lista.InsertarPrimero(253.66)
	require.EqualValues(t, 253.66, lista.BorrarPrimero())
	require.EqualValues(t, 133.15, lista.BorrarPrimero())
	require.True(t, lista.EstaVacia())
}

// Al insertar un elemento en la posición en la que se crea el iterador, efectivamente se inserta al principio.
func TestInsertarIterador(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	iterador := lista.Iterador()
	iterador.Insertar(10)
	require.EqualValues(t, 10, lista.VerPrimero())
}

// Insertar un elemento cuando el iterador está al final efectivamente es equivalente a insertar al final.
func TestIteradorInsertarAlFinal(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	iterador.Insertar(10)
	require.EqualValues(t, 10, lista.VerUltimo())
}

// Insertar un elemento en el medio se hace en la posición correcta.
func TestInsertarEnElMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Insertar(10)
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 10, lista.BorrarPrimero())
}

// Al remover el elemento cuando se crea el iterador, cambia el primer elemento de la lista.
func TestRemoverPrimerElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	iterador.Borrar()
	require.EqualValues(t, 2, lista.VerPrimero())
}

// Remover el último elemento con el iterador cambia el último de la lista.
func TestRemoverUltimoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Borrar()
	require.EqualValues(t, 3, lista.VerUltimo())
}

// Verificar que al remover un elemento del medio, este no está.
func TestRemoverElementoDelMedio(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	iterador := lista.Iterador()
	iterador.Siguiente()
	iterador.Siguiente()
	iterador.Borrar()
	require.EqualValues(t, 1, lista.BorrarPrimero())
	require.EqualValues(t, 2, lista.BorrarPrimero())
	require.EqualValues(t, 4, lista.BorrarPrimero())
	require.EqualValues(t, 5, lista.BorrarPrimero())
}

func TestIteradorTerminoIterar(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iterador := lista.Iterador()
	for iterador.HaySiguiente() {
		iterador.Siguiente()
	}
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
}

func TestIteradorSinElementos(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	iterador := lista.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Borrar() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iterador.Siguiente() })
	require.False(t, iterador.HaySiguiente())
}

func TestIteradorActualizaLargo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(1)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	iterador := lista.Iterador()
	iterador.Insertar(10)
	iterador.Insertar(20)
	iterador.Borrar()
	require.EqualValues(t, 6, lista.Largo())
}

func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(5)
	lista.InsertarUltimo(7)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += 1
		return *contador_ptr < 4
	})
	require.EqualValues(t, 4, contador)
}

func TestIteradorInterno2(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(10)
	lista.InsertarUltimo(4)
	lista.InsertarUltimo(6)
	lista.InsertarUltimo(20)
	contador := 0
	contador_ptr := &contador
	lista.Iterar(func(v int) bool {
		*contador_ptr += v
		return true
	})
	require.EqualValues(t, 44, contador)
}
