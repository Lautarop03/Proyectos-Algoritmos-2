package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) { // Se pueda crear una Pila vacía, y esta se comporta como tal.
	pila := TDAPila.CrearPilaDinamica[int]()
	require.False(t, pila == nil)
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// Se puedan apilar elementos, que al desapilarlos se mantenga el invariante de pila (que esta es LIFO). Probar con elementos diferentes,
// y ver que salgan en el orden deseado.
func TestInvarianteDePila(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	pila.Apilar(1)
	pila.Apilar(2)
	pila.Apilar(3)
	pila.Apilar(4)
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}

// Prueba de volumen: Se pueden apilar muchos elementos (1000, 10000 elementos, o el volumen que corresponda): hacer crecer la pila,
// y desapilar elementos hasta que esté vacía, comprobando que siempre cumpla el invariante. Recordar no apilar siempre lo mismo,
// validar que se cumpla siempre que el tope de la pila sea el correcto paso a paso, y que el nuevo tope después de cada desapilar también sea el correcto.
func TestPruebaDeVolumen(t *testing.T) {
	tam := 10000
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= tam; i++ {
		pila.Apilar(i)
	}
	for i := tam; i > 0; i-- {
		require.EqualValues(t, i, pila.VerTope())
		require.EqualValues(t, i, pila.Desapilar())
		require.EqualValues(t, i-1, pila.VerTope())
	}
	require.EqualValues(t, 0, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

// Condición de borde: comprobar que al desapilar hasta que está vacía hace que la pila se comporte como recién creada.
// Condición de borde: las acciones de desapilar y ver_tope en una pila a la que se le apiló y desapiló hasta estar vacía son inválidas.
func TestApilarDesapilarPilaNueva(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i <= 10; i++ {
		pila.Apilar(i)
	}
	for i := 10; i >= 0; i-- {
		pila.Desapilar()
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

// Probar apilar diferentes tipos de datos: probar con una pila de enteros, con una pila de cadenas, etc…
func TestPilaStr(t *testing.T) {
	pilastr := TDAPila.CrearPilaDinamica[string]()
	pilastr.Apilar("hola")
	pilastr.Apilar("Vaso")
	pilastr.Apilar("reloj")
	require.False(t, pilastr.EstaVacia())
	require.EqualValues(t, "reloj", pilastr.Desapilar())
	require.EqualValues(t, "Vaso", pilastr.Desapilar())
	require.EqualValues(t, "hola", pilastr.Desapilar())
	require.True(t, pilastr.EstaVacia())
}

func TestPilaInt(t *testing.T) {
	pilaint := TDAPila.CrearPilaDinamica[int]()
	pilaint.Apilar(100)
	pilaint.Apilar(666)
	pilaint.Apilar(399)
	require.EqualValues(t, 399, pilaint.Desapilar())
	require.EqualValues(t, 666, pilaint.Desapilar())
	require.EqualValues(t, 100, pilaint.Desapilar())
	require.True(t, pilaint.EstaVacia())
}

func TestPilaBool(t *testing.T) {
	pilabool := TDAPila.CrearPilaDinamica[bool]()
	pilabool.Apilar(true)
	pilabool.Apilar(false)
	require.EqualValues(t, false, pilabool.Desapilar())
	require.EqualValues(t, true, pilabool.Desapilar())
	require.True(t, pilabool.EstaVacia())
}

func TestPilaFloat(t *testing.T) {
	pilafloat := TDAPila.CrearPilaDinamica[float32]()
	pilafloat.Apilar(133.15)
	pilafloat.Apilar(253.66)
	require.EqualValues(t, 253.66, pilafloat.Desapilar())
	require.EqualValues(t, 133.15, pilafloat.Desapilar())
	require.True(t, pilafloat.EstaVacia())
}
