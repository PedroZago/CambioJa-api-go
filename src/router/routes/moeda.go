package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasMoedas = []Rota{
	{
		URI:                "/api/currencys",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarMoeda,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/currencys/{moedaID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarMoeda,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/currencys/{moedaID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarMoeda,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/currencys",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodasMoedas,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/currencys/{moedaID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarMoedaPorID,
		RequerAutenticacao: true,
	},
}
