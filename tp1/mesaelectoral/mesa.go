package mesaelectoral

type MesaElectoral interface {
	Ingresar(candidato string)

	Votar(numLista, candidato string)

	DesHacer()

	FinVotar(impugnados []int)

	ImprimirResultados(impugnados []int)
}
