package controllers

import (
	"net/http"
	"webapp/src/cookies"
)

// FazerLogout remove os dados de autenticação salvos no browser do usuário
func FazerLogout(w http.ResponseWriter, r *http.Request) {
	//Essa função em especifico não precisa nem se comunicar com API, pois o token está sendo salvo do lado do cliente e, não do lado do servidor
	//logo o logout será feito apenas no Front-END
	cookies.Deletar(w)                              //Ela não retorna erro, então após executar ela, vamos redirecionar para a tela de login
	http.Redirect(w, r, "/login", http.StatusFound) //=statusFound = 302
}
