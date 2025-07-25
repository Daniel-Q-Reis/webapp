package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
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
	response, erro := http.Post("http://localhost:5000/login", "application/json", bytes.NewBuffer(usuario)) //será para uma variavel de ambiente mais para frente .env, pois se precisar trocar o endereço no futuro e mais facil
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	token, _ := io.ReadAll(response.Body)
	//para ler esse response.Body que é do tipo io.ReadCloser, precisamos da func acima io.ReadAll(response.Body) (ele passa um slice de byte, vamos converter para string, string(token))
	fmt.Println(response.Status, string(token)) //por hora somente para saber se deu certo ->string(token) -> equivale a response.Body
}
