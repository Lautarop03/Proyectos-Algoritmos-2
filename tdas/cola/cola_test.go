package cola_test

import (
	TDACola "tdas/cola"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.False(t, cola == nil)
}

func TestEncolarDesencolar(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestInvarianteDeCola(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	cola.Encolar(1)
	cola.Encolar(2)
	cola.Encolar(3)
	cola.Encolar(4)
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
	require.EqualValues(t, 4, cola.Desencolar())
	require.True(t, cola.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func TestPruebaDeVolumen(t *testing.T) {
	tam := 10000
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i <= tam; i++ {
		cola.Encolar(i)
	}
	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, cola.VerPrimero())
		require.EqualValues(t, i, cola.Desencolar())
		require.EqualValues(t, i+1, cola.VerPrimero())
	}
	require.EqualValues(t, 10000, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestEncolarDesencolarColaNueva(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	for i := 0; i <= 10; i++ {
		cola.Encolar(i)
	}
	for i := 10; i >= 0; i-- {
		cola.Desencolar()
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
}

func TestColastr(t *testing.T) {
	colastr := TDACola.CrearColaEnlazada[string]()
	colastr.Encolar("hola")
	colastr.Encolar("Vaso")
	colastr.Encolar("reloj")
	require.False(t, colastr.EstaVacia())
	require.EqualValues(t, "hola", colastr.Desencolar())
	require.EqualValues(t, "Vaso", colastr.Desencolar())
	require.EqualValues(t, "reloj", colastr.Desencolar())
	require.True(t, colastr.EstaVacia())
}

func TestColaBool(t *testing.T) {
	colabool := TDACola.CrearColaEnlazada[bool]()
	colabool.Encolar(true)
	colabool.Encolar(false)
	require.EqualValues(t, true, colabool.Desencolar())
	require.EqualValues(t, false, colabool.Desencolar())
	require.True(t, colabool.EstaVacia())
}

func TestColaFloat(t *testing.T) {
	colafloat := TDACola.CrearColaEnlazada[float32]()
	colafloat.Encolar(133.15)
	colafloat.Encolar(253.66)
	require.EqualValues(t, 133.15, colafloat.Desencolar())
	require.EqualValues(t, 253.66, colafloat.Desencolar())
	require.True(t, colafloat.EstaVacia())
}
