package tp0

// Swap intercambia dos valores enteros.
func Swap(x *int, y *int) {
	*x, *y = *y, *x
}

// Maximo devuelve la posición del mayor elemento del arreglo, o -1 si el el arreglo es de largo 0. Si el máximo
// elemento aparece más de una vez, se debe devolver la primera posición en que ocurre.
func Maximo(vector []int) int {
	if len(vector) == 0 {
		return -1
	}
	var max int
	for i := 0; i < len(vector); i++ {
		if vector[i] > vector[max] {
			max = i
		}
	}
	return max
}

// Comparar compara dos arreglos de longitud especificada.
// Devuelve -1 si el primer arreglo es menor que el segundo; 0 si son iguales; o 1 si el primero es el mayor.
// Un arreglo es menor a otro cuando al compararlos elemento a elemento, el primer elemento en el que difieren
// no existe o es menor.
func Comparar(vector1 []int, vector2 []int) int {
	for i := 0; i < len(vector1) && i < len(vector2); i++ {
		if vector1[i] < vector2[i] {
			return -1
		} else if vector1[i] > vector2[i] {
			return 1
		}
	}
	if len(vector1) == len(vector2) {
		return 0
	} else if len(vector1) > len(vector2) {
		return 1
	}
	return -1
}

// Seleccion ordena el arreglo recibido mediante el algoritmo de selección.
func Seleccion(vector []int) {
	for i := 0; i < len(vector); i++ {
		max := Maximo(vector[:len(vector)-i])
		Swap(&vector[max], &vector[len(vector)-i-1])
	}
}

// Suma devuelve la suma de los elementos de un arreglo. En caso de no tener elementos, debe devolver 0.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func Suma(vector []int) int {
	total := 0
	return suma(vector, total)
}

func suma(vector []int, total int) int {
	if len(vector) == 0 {
		return total
	}
	total += vector[0]
	return suma(vector[1:], total)
}

// EsCadenaCapicua devuelve si la cadena es un palíndromo. Es decir, si se lee igual al derecho que al revés.
// Esta función debe implementarse de forma RECURSIVA. Se puede usar una función auxiliar (que sea
// la recursiva).
func EsCadenaCapicua(cadena string) bool {
	if len(cadena) <= 1 {
		return true
	}
	if cadena[0] == cadena[len(cadena)-1] {
		return EsCadenaCapicua(cadena[1 : len(cadena)-1])
	}
	return false
}
