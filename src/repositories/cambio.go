package repositories

import (
	"api/src/models"
	"database/sql"
)

type Cambios struct {
	db *sql.DB
}

func NovoRepositorioCambio(db *sql.DB) *Cambios {
	return &Cambios{db}
}

func (repositorio Cambios) CriarCambio(cambio models.Cambio) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO cambio (valorTransferido, resultadoConversao, usuarioID, moedaID) values (?, ?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(cambio.ValorOrigem, cambio.ValorFinal, cambio.UsuarioID, cambio.MoedaID)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Cambios) BuscarTodosCambiosPorUsuarioEMoeda(cambio models.Cambio) ([]models.Cambio, error) {
	linhas, err := repositorio.db.Query(
		`SELECT cb.cambioID, cb.valorTransferido, cb.resultadoConversao, mo.nome, us.nome, cb.dataCambio
			FROM cambio AS cb, usuario AS us, moeda AS mo
			WHERE mo.moedaID = ? AND
			cb.moedaID = mo.moedaID AND
			cb.usuarioID = ? AND
			cb.usuarioID = us.usuarioID`, cambio.MoedaID, cambio.UsuarioID,
	)

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var cambios []models.Cambio

	for linhas.Next() {
		var cambio models.Cambio

		if err = linhas.Scan(
			&cambio.CambioID,
			&cambio.ValorOrigem,
			&cambio.ValorFinal,
			&cambio.MoedaNome,
			&cambio.UsuarioNome,
			&cambio.DataCambio,
		); err != nil {
			return nil, err
		}

		cambios = append(cambios, cambio)
	}

	return cambios, nil
}
