package models

import (
	"errors"
	"math"
	"time"
)

type Cambio struct {
	CambioID    uint64    `json:"cambioID,omitempty"`
	DataCambio  time.Time `json:"dataCambio,omitempty"`
	ValorOrigem float64   `json:"valorOrigem,omitempty"`
	ValorFinal  float64   `json:"valorFinal,omitempty"`
	UsuarioID   uint64    `json:"usuarioID,omitempty"`
	UsuarioNome string    `json:"usuarioNome,omitempty"`
	MoedaID     uint64    `json:"moedaID,omitempty"`
	MoedaNome   string    `json:"moedaNome,omitempty"`
	SaldoMoeda  float64   `json:"saldoMoeda,omitempty"`
}

func (cambio *Cambio) Preparar() error {
	if err := cambio.validar(); err != nil {
		return err
	}

	if err := cambio.formatar(); err != nil {
		return err
	}

	return nil
}

func (cambio *Cambio) validar() error {
	if cambio.ValorOrigem > 9999.99 {
		return errors.New("o limite para a transfêrencia é de RS 9,999.99")
	}

	return nil
}

func (cambio *Cambio) formatar() error {
	cambio.ValorOrigem = math.Round(cambio.ValorOrigem*100) / 100
	cambio.ValorFinal = math.Round(cambio.ValorFinal*100) / 100

	return nil
}
