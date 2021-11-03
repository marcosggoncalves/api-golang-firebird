package controllers

import (
	"api/api/src/database"
	"api/api/src/repositories"
	"api/api/src/responses"
	"net/http"
	"github.com/gorilla/mux"
	"strconv"
)

func BuscarContrato(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)

	contratoID, erro := strconv.ParseInt(parametros["contratoId"], 0, 64)

	db, erro := database.Conectar();
	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}

	defer db.Close()

	repositorio := repositories.NovoRepositorioContrato(db)
	contratos, erro := repositorio.BuscarContrato(contratoID)

	if erro != nil {
		responses.Erro(w, http.StatusInternalServerError, erro)
		return
	}
	responses.JSON(w, http.StatusOK, contratos)
}