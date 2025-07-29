package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
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

// CarregarPaginaPrincipal carrega a página principal com as publicações
func CarregarPaginaPrincipal(w http.ResponseWriter, r *http.Request) {
	//vai buscar as publicações la no banco de dados da API, e depois exibir isso na tela
	url := fmt.Sprintf("%s/publicacoes", config.APIURL)
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil) //dados é nil, pois não vamos passar nada ja que se trata de um get
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close() //fechar o corpo da requisição quando a função terminar

	// Aqui fazemos a avaliação de status code
	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	//Pegando as publicações que estão vindo na resposta (response), para tal usaremos o decode
	var publicacoes []modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacoes); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	// Implementando se uma publicação la na pagina inicial pertence ou não ao usuário que está logado
	cookie, _ := cookies.Ler(r)
	usuarioID, _ := strconv.ParseUint(cookie["id"], 10, 64) //tem que converter para uint pois vamos comparar com o autorID da publicação, e os tipos tem que ser iguais, essa comparação será feita dentro do arquivo html, usando os templates do go
	// se chegar até aqui não precisa de tratamento de erros,pois sabemos que não irá dar erro, pois já passou pelo middleware que havalia o erro, e se der erro, irá retornar ao login, e não vai renderizar a pagina home

	//depois que eu tiver os valores do usuarioID, podemos fazer a struct no executarTemplate
	utils.ExecutarTemplate(w, "home.html", struct {
		Publicacoes []modelos.Publicacao
		UsuarioID   uint64
	}{
		Publicacoes: publicacoes,
		UsuarioID:   usuarioID,
	})
	//como agora estamos passando mais de um campo (Publicacoes e UsuarioID) pra pagina, devemos ir no home.html e referenciar esses campo (antes estava só publicacoes como range ., agora {{range .Publicacoes}})
}
