package repositories

import (
	"api/src/models"
	"database/sql"
)

type Moedas struct {
	db *sql.DB
}

func NovoRepositorioMoeda(db *sql.DB) *Moedas {
	return &Moedas{db}
}

func (repositorio Moedas) CriarMoeda(moeda models.Moeda) (uint64, error) {
	statement, err := repositorio.db.Prepare("INSERT INTO moeda (nome, codISO, cotacao) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(moeda.Nome, moeda.CodISO, moeda.Cotacao)
	if err != nil {
		return 0, err
	}

	ultimoIDInserido, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(ultimoIDInserido), nil
}

func (repositorio Moedas) AtualizarMoeda(moedaID uint64, moeda models.Moeda) error {
	statement, err := repositorio.db.Prepare("UPDATE moeda SET nome = ?, codISO = ?, cotacao = ? WHERE moedaID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(moeda.Nome, moeda.CodISO, moeda.Cotacao, moedaID); err != nil {
		return err
	}

	return nil
}

func (repositorio Moedas) DeletarMoeda(moedaID uint64) error {
	statement, err := repositorio.db.Prepare("DELETE FROM moeda WHERE moedaID = ?")
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err := statement.Exec(moedaID); err != nil {
		return err
	}

	return nil
}

func (repositorio Moedas) BuscarMoedaPorID(moedaID uint64) (models.Moeda, error) {
	linha, err := repositorio.db.Query("SELECT moedaID, nome, codISO, cotacao FROM moeda WHERE moedaID = ?", moedaID)

	if err != nil {
		return models.Moeda{}, err
	}
	defer linha.Close()

	var moeda models.Moeda

	if linha.Next() {
		if err = linha.Scan(
			&moeda.MoedaID,
			&moeda.Nome,
			&moeda.CodISO,
			&moeda.Cotacao,
		); err != nil {
			return models.Moeda{}, err
		}
	}

	return moeda, nil
}

func (repositorio Moedas) BuscarTodasMoedas() ([]models.Moeda, error) {
	linhas, err := repositorio.db.Query("SELECT moedaID, nome, codISO, cotacao FROM moeda")

	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var moedas []models.Moeda

	for linhas.Next() {
		var moeda models.Moeda

		if err = linhas.Scan(
			&moeda.MoedaID,
			&moeda.Nome,
			&moeda.CodISO,
			&moeda.Cotacao,
		); err != nil {
			return nil, err
		}

		moedas = append(moedas, moeda)
	}

	return moedas, nil
}

func (repository Moedas) PegarCotacao(moedaID uint64) (float64, error) {
	queryValue, err := repository.db.Query("SELECT cotacao FROM moeda WHERE moedaID = ?", moedaID)
	if err != nil {
		return 0, err
	}
	defer queryValue.Close()

	var cotacaoMoeda float64

	if queryValue.Next() {
		if err := queryValue.Scan(&cotacaoMoeda); err != nil {
			return 0, err
		}
	}

	return cotacaoMoeda, nil
}
