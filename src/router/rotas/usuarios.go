package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

//Aqui temos todas as rotas relativas aos usuários, tanto as rotas que vão chamar a api, quanto as rotas que vão chamar as paginas.

var rotasUsuarios = []Rota{
	{
		URI:                "/criar-usuario",
		Metodo:             http.MethodGet,
		Funcao:             controllers.CarregarPaginaDeCadastroDeUsuario,
		RequerAutenticacao: false,
	},
	{
		URI:                "/usuarios",
		Metodo:             http.MethodPost,
		Funcao:             controllers.CriarUsuario,
		RequerAutenticacao: false,
	},
}
