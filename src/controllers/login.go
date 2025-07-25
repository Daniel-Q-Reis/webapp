package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

// Fazer login utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) { //a rota que vai fazer o processamento é em go, porem quem vai chamar essa rota em go é nosso front end em java script, pelo ajax, para chamar a rota de login, assim como a criação de usuario
	//vamos em views login.html (assim como fizemos com a criação de usuarios), lá vamos linkar com arquivo em js, que ira fazer o envio do formulario do login e também o envio da rota em go (login.js na pasta js)
	//assim como na função criar usuários temos:
	r.ParseForm() //vai pegar o corpo da requisição e vai colocar em algumas propriedades que vamos conseguir acessar para pegar seus valores

	//agora ja posso começar a pegar os valores
	// com os dados fazemos um map deles e convertemos para json, usando o json.Marshal
	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//chegando nesse ponto podemos fazer a requisição na api
	url := fmt.Sprintf("%s/login", config.APIURL)                                  //%s aqui equivale a http://localhost:5000
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario)) //era: ("http://localhost:5000/login", ""application/json, bytes...) será para uma variavel de ambiente mais para frente .env, pois se precisar trocar o endereço no futuro e mais facil -> agora : url := fmt.Sprintf("%s/login", config.APIURL)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response) //que veio ali do http.Post()
		return
	}

	var DadosAutenticacao modelos.DadosAutenticacao //struct criada para retornar o ID e o Token separados
	if erro = json.NewDecoder(response.Body).Decode(&DadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// Nesse momento eu tenhos as informacoes de autenticação desse usuario salvo na variavel DadosAutenticação
	// Então preciso salvar esses dados em algum lugar para não perder eles (pois não queremos que o usuário ao mudar de pagina perca seu login)
	// então vou pegar esses dados e salvar dentro de um Browser do usuário através de um cookie, existe um pacote de go especifico para isso
	// o cookie então será utilizado em todas as rotas para que assim eu veja de forma segura se o usuário está autenticado
	// Após implementar-mos isso tudo vamos dar um resposta simples la pro nosso ajax, não vamos passar nada nela, vai ser somente um: respostas.Json(w, http.StatusOK, nil)
	// Logo não iremos devolver esses dados da autenticação... eles vao ser usados somente pelo COOKIE e vai ficar salvo no browser do usuário
	// Essa forma e muito comum quando usamos JWT (Jason web token), cria um token e devolve para o usuá rio
	// Para isso vamos criar agora nosso package config que vai conter as variaveis de ambiente dentro de nossa aplicação web
	// Vamos fazer isso antes de mexer com os dados de autenticação porque vamos precisar de duas informações para criação do cookie
	// eles vao estar armazenados em variaveis de ambiente, são bem parecidas com o secret key que geramos na api para a geração do token
	// ai ja vamos aproveitar para colocar a URL da API em uma variável de ambiene
	// Vamos criar entao dentro de SRC uma pasta CONFIG/config. com package config

	respostas.JSON(w, http.StatusOK, nil)
}
