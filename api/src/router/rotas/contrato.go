package rotas

import (
	"api/api/src/controllers"
	"net/http"
)

var rotasContrato = []Rota{
	{
		URI:                "/contrato/{contratoId}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarContrato,
		RequerAutenticacao: false,
	},
}
