package rotas

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

// Rota representa todas as rotas da Aplicação Web
type Rota struct {
	URI                string
	Metodo             string
	Funcao             func(http.ResponseWriter, *http.Request)
	RequerAutenticacao bool
}

// Configurar coloca todas as rotas dentro do router
func Configurar(router *mux.Router) *mux.Router {
	rotas := rotasLogin
	rotas = append(rotas, rotasUsuarios...)
	rotas = append(rotas, rotaPaginaPrincipal)

	for _, rota := range rotas {
		//O que vamos fazer aqui é, se precisar autenticar, eu vou chamar também a função autenticar do middleware
		if rota.RequerAutenticacao {
			router.HandleFunc(rota.URI,
				middlewares.Logger(middlewares.Autenticar(rota.Funcao)),
			).Methods(rota.Metodo)
		} else { //se caso não precise de autenticação, é so fazer o logger, sem o autenticar
			router.HandleFunc(rota.URI,
				middlewares.Logger(rota.Funcao),
			).Methods(rota.Metodo)
		}

	}

	//fileserver irá apontar para o Go onde estão os arquivos que iremos usar para estilo quanto os arquivos java script
	fileServer := http.FileServer(http.Dir("./assets/"))
	//para evitar que o html tenha que voltar pastas pelo .. (tipo cd ..) usaremos PathPrefix
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
