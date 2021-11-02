package repositories

import (
	"api/api/src/models"
	"database/sql"
	"fmt"
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

func (repositorio Usuario) Buscar(nomeOuNick string)([]models.Usuario, error){
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick)

	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome like ? or nick like ?",
		nomeOuNick, nomeOuNick,
	);

	if erro != nil{
		return nil, erro
	}

	defer linhas.Close()

	var usuarios []models.Usuario

	for linhas.Next(){
		var usuario models.Usuario

		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return nil, erro
		}

		usuarios = append(usuarios, usuario)
	}

	return usuarios, nil
}

func (repositorio Usuario) BuscarPorID(ID uint64) (models.Usuario, error){
	linhas, erro := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where id  = ?",
		ID,
	)

	if erro != nil{
		return models.Usuario{}, erro
	}

	defer linhas.Close()

	var usuario models.Usuario

	if linhas.Next(){
		if erro = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); erro != nil{
			return models.Usuario{}, erro
		}
	}

	return usuario, nil
}


func (repositorio Usuario) Atualizar(ID uint64, Usuario models.Usuario) error{
	statement, erro := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)

	if erro != nil {
		return  erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(Usuario.Nome, Usuario.Nick, Usuario.Email, ID); erro != nil{
		return erro
	}

	return nil
}

func (repositorio Usuario) Deletar(ID uint64) error{
	statement, erro := repositorio.db.Prepare(
		"delete from usuarios where id = ?",
	)

	if erro != nil {
		return  erro
	}

	defer statement.Close()

	if _, erro := statement.Exec(ID); erro != nil{
		return erro
	}

	return nil
}