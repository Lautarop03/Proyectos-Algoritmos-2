package ordenamientos

import (
	"rerepolez/votos"
)

func RadixSort(arr []int) []int {
	ordenado := countingSort(arr, 99, func(n int) int {
		return n % 100 // xx.xxx.x00
	})
	ordenado = countingSort(ordenado, 99, func(n int) int {
		return (n % 10000) / 100 // xx.xx0.0xx
	})
	ordenado = countingSort(ordenado, 99, func(n int) int {
		return (n % 1000000) / 10000 // xx.00x.xxx
	})
	ordenado = countingSort(ordenado, 99, func(n int) int {
		return n / 1000000 // 00.xxx.xxx
	})

	return ordenado
}

func countingSort(arreglo []int, max int, f func(int) int) []int {

	frequencias := make([]int, max+1)

	for _, par := range arreglo {
		frequencias[f(par)]++
	}

	suma := 0
	sumas_acumuladas := make([]int, max+1)

	for i, freq := range frequencias {
		sumas_acumuladas[i] += suma
		suma += freq
	}

	ordenado := make([]int, len(arreglo))
	for _, elem := range arreglo {
		ordenado[sumas_acumuladas[f(elem)]] = elem
		sumas_acumuladas[f(elem)]++
	}
	return ordenado
}

func BusquedaBinaria(dni int, listaVotantes []votos.Votante) votos.Votante {
	return busqueda(0, len(listaVotantes), dni, listaVotantes)
}

func busqueda(inicio, fin, dni int, lista []votos.Votante) votos.Votante {
	if inicio >= fin {
		return nil
	}
	medio := (inicio + fin) / 2

	if lista[medio].LeerDNI() == dni {
		votante := lista[medio]
		return votante
	}
	if lista[medio].LeerDNI() > dni {
		return busqueda(inicio, medio, dni, lista)
	} else {
		return busqueda(medio+1, fin, dni, lista)
	}
}
