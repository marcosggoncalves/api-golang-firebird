package controllers

import (
	"api/api/src/database"
	"api/api/src/modelos"
	"api/api/src/repositorios"
	"api/api/src/respostas"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		respostas.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario modelos.Usuario
	if erro = json.Unmarshal(bodyRequest, &usuario); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/// Campos vazios
	if erro = usuario.Preparar(); erro != nil {
		respostas.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/// Abrir conexão com banco de dados
	db, erro := database.Conectar()
	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	// Finalizar conexão com banco de dados
	defer db.Close()

	repositorio := repositorios.NovoRepositorioDeUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		respostas.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	// w.Write([]byte(fmt.Sprintf("Id inserid: %d", usuarioId)))
	respostas.JSON(w, http.StatusCreated, usuario)
}

func BuscarUsuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando todos usuários"))
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Buscando um usuário"))
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Atualizando usuário"))
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deletando usuário"))
}
