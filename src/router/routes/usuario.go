package routes

import (
	"api/src/controllers"
	"net/http"
)

var rotasUsuario = []Rota{
	{
		URI:                "/api/users",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},

	{
		URI:                "/api/users/{usuarioID}",
		Metodo:             http.MethodPut,
		Funcao:             controllers.AtualizarUsuario,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/users/{usuarioID}",
		Metodo:             http.MethodDelete,
		Funcao:             controllers.DeletarUsuario,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/users",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarTodosUsuarios,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/users/{usuarioID}",
		Metodo:             http.MethodGet,
		Funcao:             controllers.BuscarUsuarioPorID,
		RequerAutenticacao: true,
	},

	{
		URI:                "/api/users/{usuarioID}/atualizar-senha",
		Metodo:             http.MethodPost,
		Funcao:             controllers.AtualizarSenha,
		RequerAutenticacao: true,
	},
}
