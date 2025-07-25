package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"webapp/src/respostas"
)

//Aqui ficarão as rotas relativas a manipulação de usuários

// CriarUsuario chama a API para cadastrar um usuário no banco de dados
func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //vai pegar o corpo da requisição e vai colocar em algumas propriedades que vamos conseguir acessar para pegar seus valores

	//agora ja posso começar a pegar os valores
	// com os dados fazemos um map deles e convertemos para json, usando o json.Marshal
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"email": r.FormValue("email"),
		"nick":  r.FormValue("nick"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//bytes.NewBuffer faz com que a gente consiga ler os slices do JSON em string
	//Abrindo comunicação bastante simples com a nossa API, já que por enquanto até aqui não precisamos de autenticação

	response, erro := http.Post("http://localhost:5000/usuarios", "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close() //é obrigatório, mesmo que o corpo do response.body esteja vazio

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return //o return é para nao chegar no respostas.JSON abaixo e me dar a resposta duas vezes
	}

	respostas.JSON(w, response.StatusCode, nil)
}
