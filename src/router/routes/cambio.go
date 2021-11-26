package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasCambios = []Rota{
	{
		URI:                "/api/exchanges",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarCambio,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/list-exchanges",
		Metodo:             http.MethodPost,
		Funcao:             controllers.BuscarTodosCambiosPorUsuarioEMoeda,
		RequerAutenticacao: true,
	},
}
