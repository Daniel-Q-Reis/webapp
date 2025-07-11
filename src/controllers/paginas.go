package controllers

import (
	"net/http"
	"webapp/src/utils"
)

//Neste arquivo, ficam todas as funções que irão renderizar as paginas

// Carregar tela de login irá carregar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//login.html é o nome do , terceiro argumento é nil, pois não iremos jogar nenhum dado variavel na tela de login, será sempre um conteúdo fixo
	utils.ExecutarTemplate(w, "login.html", nil) //esses arquivos vão se encontrar na pasta views da aplicação

}

// CarregarPaginaDeCadastroDeUsuario vai carregar a página de cadastro de usuário
func CarregarPaginaDeCadastroDeUsuario(w http.ResponseWriter, r *http.Request) {
	utils.ExecutarTemplate(w, "cadastro.html", nil)
}
