package controllers

import (
	"api/api/src/database"
	"api/api/src/models"
	"api/api/src/repositories"
	"api/api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	bodyRequest, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		responses.Erro(w, http.StatusUnprocessableEntity, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(bodyRequest, &usuario); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/// Campos vazios
	if erro = usuario.Preparar(); erro != nil {
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	/// Abrir conexão com banco de dados
	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	// Finalizar conexão com banco de dados
	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuario(db)
	usuario.ID, erro = repositorio.Criar(usuario)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	// w.Write([]byte(fmt.Sprintf("Id inserid: %d", usuarioId)))
	responses.JSON(w, http.StatusCreated, usuario)
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
