package models

import (
	"errors"
	"strings"
)

type Moeda struct {
	MoedaID uint64  `json:"moedaID,omitempty"`
	Nome    string  `json:"nome,omitempty"`
	CodISO  string  `json:"codISO,omitempty"`
	Cotacao float64 `json:"cotacao,omitempty"`
}

func (moeda *Moeda) Preparar() error {
	if err := moeda.validar(); err != nil {
		return err
	}

	moeda.formatar()

	return nil
}

func (moeda *Moeda) validar() error {
	if moeda.Nome == "" {
		return errors.New("o nome é obrigatório e não pode estar em branco")
	}

	return nil
}

func (moeda *Moeda) formatar() {
	moeda.Nome = strings.TrimSpace(moeda.Nome)
}
