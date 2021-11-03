package database

import (
	"api/api/src/config"
	"database/sql"

	_"github.com/nakagami/firebirdsql"
)

func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("firebirdsql", config.StringConexaoBanco)

	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}

	return db, nil
}
