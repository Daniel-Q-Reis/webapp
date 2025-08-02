package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/requisicoes"
	"webapp/src/respostas"

	"github.com/gorilla/mux"
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
	url := fmt.Sprintf("%s/usuarios", config.APIURL)                               //%s aqui equivale a http://localhost:5000 /usuarios
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario)) //url pegou a variavel de ambiente config.APIURL = url = http://localhost:5000/usuarios
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

// PararDeSeguirUsuario chama a API para parar de seguir um usuário
func PararDeSeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/parar-de-seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// SeguirUsuario chama a API para seguir um usuário
func SeguirUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/usuarios/%d/seguir", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// EditarUsuario chama a API para editar um usuário
func EditarUsuario(w http.ResponseWriter, r *http.Request) {
	r.ParseForm() //Populamos o r.Form e R.PostForm e,
	//Aqui vamos criar um usuario
	usuario, erro := json.Marshal(map[string]string{
		"nome":  r.FormValue("nome"),
		"nick":  r.FormValue("nick"),
		"email": r.FormValue("email"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPut, url, bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// AtualizarSenha chama a API para atualizar a senha do usuário
func AtualizarSenha(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//Aqui montamos o objeto(map) que irá ser enviado para a API, tanto com a senha antiga quanto com a senha nova
	senhas, erro := json.Marshal(map[string]string{
		"atual": r.FormValue("atual"),
		"nova":  r.FormValue("nova"),
	})
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	//pegamos o usuario ID para montar a URL, assim a API indentifica o Usuario e compara as senhas
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d/atualizar-senha", config.APIURL, usuarioID)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodPost, url, bytes.NewBuffer(senhas)) //nesse caso aqui no corpo da requisição, vamos mandar as senhas
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}

// DeletarUsuario chama a API para deletar um usuário
func DeletarUsuario(w http.ResponseWriter, r *http.Request) {
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	url := fmt.Sprintf("%s/usuarios/%d", config.APIURL, usuarioID)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodDelete, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	respostas.JSON(w, response.StatusCode, nil)
}
