package controllers

import (
	"api/api/src/database"
	"api/api/src/models"
	"api/api/src/repositories"
	"api/api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"github.com/gorilla/mux"
	"strconv"
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
	if erro = usuario.Preparar("cadastro"); erro != nil {
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
	nomeOuNick :=  strings.ToLower(r.URL.Query().Get("usuario"))
	db, erro := database.Conectar();

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuario(db)
	usuarios, erro := repositorio.Buscar(nomeOuNick)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, usuarios)
}

func BuscarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	usuarioID, erro := strconv.ParseUint(parametros["usuariosId"], 10, 64)

	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuario(db)
	usuario, erro := repositorio.BuscarPorID(usuarioID)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusOK, usuario)
}

func AtualizarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	
	usuarioID, erro := strconv.ParseUint(parametros["usuariosId"], 10, 64)

	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	corpoRequisicao, erro := ioutil.ReadAll(r.Body)

	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	var usuario models.Usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	if erro = usuario.Preparar("edicao"); erro != nil{
		responses.Erro(w, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuario(db)

	if erro = repositorio.Atualizar(usuarioID, usuario); erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	
	usuarioID, erro := strconv.ParseUint(parametros["usuariosId"], 10, 64)

	if erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	db, erro := database.Conectar()
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioDeUsuario(db)

	if erro = repositorio.Deletar(usuarioID); erro != nil{
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
