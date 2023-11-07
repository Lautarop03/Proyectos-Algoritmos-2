package cola_prioridad_test

import (
	"fmt"
	"math/rand"
	TDAHeap "tdas/cola_prioridad"
	"testing"

	"github.com/stretchr/testify/require"
)

func comp(a, b int) int {
	return a - b
}

func TestHeapVacio(t *testing.T) {
	heap := TDAHeap.CrearHeap(comp)
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestEncolarDesencolar(t *testing.T) {
	heap := TDAHeap.CrearHeap(comp)
	heap.Encolar(10)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 10, heap.VerMax())
	require.EqualValues(t, 1, heap.Cantidad())
	require.EqualValues(t, 10, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestCondicionHeap(t *testing.T) {
	heap := TDAHeap.CrearHeap(comp) // Maximos
	heap.Encolar(5)
	heap.Encolar(10)
	heap.Encolar(50)
	heap.Encolar(1000)
	heap.Encolar(60)
	heap.Encolar(1)
	require.EqualValues(t, 1000, heap.Desencolar())
	require.EqualValues(t, 60, heap.Desencolar())
	require.EqualValues(t, 50, heap.Desencolar())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestManejoCantidad(t *testing.T) {
	heap := TDAHeap.CrearHeap(comp)
	for i := 0; i < 100; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(i)
	}
	require.EqualValues(t, 100, heap.Cantidad())
	for i := 100; i > 0; i-- {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Desencolar()
	}
	require.EqualValues(t, 0, heap.Cantidad())
}

func TestCrearHeapArr(t *testing.T) {
	arr := []int{10, 20, 5, 6, 7, 15}
	heap := TDAHeap.CrearHeapArr(arr, comp)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
}
func TestCrearHeapArrEncolarDesencolar(t *testing.T) {
	arr := []int{10, 20, 5, 6, 7, 15}
	heap := TDAHeap.CrearHeapArr(arr, comp)
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 6, heap.Cantidad())
	heap.Encolar(12)
	heap.Encolar(1)
	require.EqualValues(t, 20, heap.VerMax())
	require.EqualValues(t, 8, heap.Cantidad())
	require.EqualValues(t, 20, heap.Desencolar())
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 12, heap.Desencolar())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 7, heap.Desencolar())
	require.EqualValues(t, 6, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, 1, heap.Desencolar())
	require.True(t, heap.EstaVacia())
	require.EqualValues(t, 0, heap.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { heap.Desencolar() })
}

func TestPruebaDeVolumen(t *testing.T) {
	tam := 100000
	heap := TDAHeap.CrearHeap(comp)
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Encolar(rand.Intn(99999))
	}
	require.False(t, heap.EstaVacia())
	require.EqualValues(t, 100000, heap.Cantidad())
	for i := tam; i > 0; i-- {
		require.EqualValues(t, i, heap.Cantidad())
		heap.Desencolar()
	}
	require.EqualValues(t, 0, heap.Cantidad())
	require.True(t, heap.EstaVacia())
}

func TestPruebaDeVolumen2(t *testing.T) {
	tam := 12500
	heap := TDAHeap.CrearHeap(comp)
	for i := 0; i < tam; i++ {
		heap.Encolar(i)
	}
	require.EqualValues(t, 12500, heap.Cantidad())
	for i := tam; i > 0; i-- {
		elem := heap.Desencolar()
		require.EqualValues(t, elem, i-1)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{10, 20, 30, 0, 50, 60, 1, 2, 3, 5, 9}
	TDAHeap.HeapSort(arr, comp)
	res := []int{0, 1, 2, 3, 5, 9, 10, 20, 30, 50, 60}
	require.EqualValues(t, res, arr)
}

// nuevo test
func TestNumerosNegativos(t *testing.T) {
	heap := TDAHeap.CrearHeap(comp)
	heap.Encolar(-1)
	heap.Encolar(10)
	heap.Encolar(5)
	heap.Encolar(-5)
	heap.Encolar(15)
	require.EqualValues(t, 15, heap.Desencolar())
	require.EqualValues(t, 10, heap.Desencolar())
	require.EqualValues(t, 5, heap.Desencolar())
	require.EqualValues(t, -1, heap.Desencolar())
	require.EqualValues(t, -5, heap.Desencolar())
}

type Vuelo struct {
	prioridad int
	nro       int
}

func TestNumerosNegatsivos(t *testing.T) {
	vuelos := Vuelo{26, 6391}
	vuelo1 := Vuelo{15, 1086}
	vuelo2 := Vuelo{15, 4701}
	vuelo3 := Vuelo{11, 654}
	vuelo4 := Vuelo{10, 5460}
	vuelo5 := Vuelo{10, 5948}
	arr := []Vuelo{}
	arr = append(arr, vuelos)
	arr = append(arr, vuelo1)
	arr = append(arr, vuelo2)
	arr = append(arr, vuelo3)
	arr = append(arr, vuelo4)
	arr = append(arr, vuelo5)
	heap := TDAHeap.CrearHeapArr(arr, func(vuelo Vuelo, vuelo2 Vuelo) int {
		if vuelo.prioridad == vuelo2.prioridad {
			return vuelo2.nro - vuelo.nro
		}
		return vuelo.prioridad - vuelo2.prioridad
	})
	for i := heap.Cantidad(); i > 0; i-- {
		elem := heap.Desencolar()
		fmt.Println(elem.prioridad, " - ", elem.nro)
	}
}
