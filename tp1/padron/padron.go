package padron

import "rerepolez/votos"

type Padron interface {
	// Devuelve true si el votante esta en el padron
	EstaEnPadron(int) bool

	// Si existe devuelve el Votante, sino nil
	BuscarVotante(int) votos.Votante
}
