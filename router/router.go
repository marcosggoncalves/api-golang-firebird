package router

import (
	"firebird-golang/middleware"

	"github.com/gorilla/mux"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/api/contrato/{id}", middleware.GetContrato).Methods("GET", "OPTIONS")

	return router
}
