package repositories

import (
	"api/api/src/models"
	"database/sql"
)

type Usuario struct {
	db *sql.DB
}

func NovoRepositorioDeUsuario(db *sql.DB) *Usuario {
	return &Usuario{db}
}

func (repositorio Usuario) Criar(Usuario models.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?,?,?,?)",
	)

	if erro != nil {
		return 0, erro
	}

	defer statement.Close()

	resultado, erro := statement.Exec(Usuario.Nome, Usuario.Nick, Usuario.Email, Usuario.Senha)

	if erro != nil {
		return 0, erro
	}

	ultimoIdInserido, erro := resultado.LastInsertId()

	if erro != nil {
		return 0, erro
	}

	return uint64(ultimoIdInserido), nil
}
