package controllers

import (
	"net/http"
	"webapp/src/utils"
)

// Carregar tela de login irá carregar a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//login.html é o nome do , terceiro argumento é nil, pois não iremos jogar nenhum dado variavel na tela de login, será sempre um conteúdo fixo
	utils.ExecutarTemplate(w, "login.html", nil)

}
