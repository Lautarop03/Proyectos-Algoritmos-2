package votos

import (
	errores "rerepolez/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni          int
	votos        Voto
	estadoDeVoto bool
	historial    TDAPila.Pila[Voto]
}

func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni: dni, historial: TDAPila.CrearPilaDinamica[Voto]()}
}

func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.estadoDeVoto {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.historial.Apilar(votante.votos)
	votante.votos.VotoPorTipo[tipo] = alternativa
	if alternativa == 0 {
		votante.votos.Impugnado = true
	}
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.estadoDeVoto {
		return errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if votante.historial.EstaVacia() {
		return errores.ErrorNoHayVotosAnteriores{}
	}
	votoPrevio := votante.historial.Desapilar()
	votante.votos = votoPrevio
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.estadoDeVoto {
		return Voto{}, errores.ErrorVotanteFraudulento{}
	}

	votante.estadoDeVoto = true

	return votante.votos, nil
}
