package database

import (
	"api/src/configs"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

// Conectar abre a conexão com o banco de dados e a retorna.
func ConectarBD() (*sql.DB, error) {
	db, err := sql.Open("mysql", configs.ConexaoBD)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
