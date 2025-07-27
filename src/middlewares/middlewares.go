package middlewares

import (
	"log"
	"net/http"
	"webapp/src/cookies"
)

// Logger escreve informações da requisição no terminal
func Logger(proximaFuncao http.HandlerFunc) http.HandlerFunc { //alinhamento de funções, uma na outra
	return func(w http.ResponseWriter, r *http.Request) {
		log.Printf("\n %s %s %s", r.Method, r.RequestURI, r.Host) //informações relativas a requisição
		proximaFuncao(w, r)                                       //executa a proxima função passando o responsewriter e o request
	}
}

// Autenticar verifica a existência de cookies
func Autenticar(proximaFuncao http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//Termos que criar a função ler la no cookies
		//antes de executar a ProximaFuncao() vindo no parametro, ela vai verificar se os dados da autenticação estão salvos no cookie
		if _, erro := cookies.Ler(r); erro != nil { //é importante frizar que essa função somente verifica se estão la, mas não se reponsabiliza para verificar se eles são validos e estão corretos
			http.Redirect(w, r, "/login", http.StatusFound) //por exemplo verificar se o ID e o token, que estão lá são os corretos e estão validos, isso é responsabilidade da API
			return                                          //http.StatusFound = 302 //por exemplo se eu fizer uma requisição na API, e o token for invalido, a requisição nao será completada, pois o JWT não irá permitir
		} //basicamente esse redirect significa, se der erro na autenticação, volta para pagina /login e loga novamente e da o return, para nao executar a proxima função
		proximaFuncao(w, r)
	}
}
