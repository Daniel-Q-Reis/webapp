package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webapp/src/config"
	"webapp/src/cookies"
	"webapp/src/modelos"
	"webapp/src/requisicoes"
	"webapp/src/respostas"
	"webapp/src/utils"

	"github.com/gorilla/mux"
)

//Neste arquivo, ficam todas as funções que irão renderizar as paginas

// Carregar tela de login renderiza a tela de login
func CarregarTelaDeLogin(w http.ResponseWriter, r *http.Request) {
	//Aqui vamos verificar se o usuário já está autenticado, pois se ele estiver, vamos redireciona-lo para a pagina /home, ao invés de executar utils.ExecutarTemplate(w, "login.html", nil)
	cookie, _ := cookies.Ler(r) //aqui teremos o Map e vamos verificar se a propriedade cookie não está em branco, se não estiver vamos dar o redirect para o home
	if cookie["token"] != "" {
		http.Redirect(w, r, "/home", http.StatusFound) //302
		return
	}
	//login.html é o nome do , terceiro argumento é nil, pois não iremos jogar nenhum dado variavel na tela de login, será sempre um conteúdo fixo
	utils.ExecutarTemplate(w, "login.html", nil) //esses arquivos vão se encontrar na pasta views da aplicação
}

// CarregarPaginaDeCadastroDeUsuario carrega a página de cadastro de usuário
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

// CarregarPaginaDeAtualizacaoDePublicacao carrega a pagina de edição de publicação
func CarregarPaginaDeAtualizacaoDePublicacao(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	publicacaoID, erro := strconv.ParseUint(parametros["publicacaoId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/publicacoes/%d", config.APIURL, publicacaoID) //vai pegar uma unica publicação especifica para atualizar/atualizar
	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}

	var publicacao modelos.Publicacao
	if erro = json.NewDecoder(response.Body).Decode(&publicacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	//Chegando aqui sem erros, vamos jogar essa publicação na tela
	utils.ExecutarTemplate(w, "atualizar-publicacao.html", publicacao)
}

// CarregarPaginaDeUsuarios carrega a página com os usuários que atendem o filtro passado
func CarregarPaginaDeUsuarios(w http.ResponseWriter, r *http.Request) {
	//Aqui vamos pegar o dado da URL que está sendo buscado pelo usuario, usuario=...
	nomeOuNick := strings.ToLower(r.URL.Query().Get("usuario"))             //nome que está vindo antes do = na URL ex: usuario=usuario3 (chave, valor no mapa)
	url := fmt.Sprintf("%s/usuarios?usuario=%s", config.APIURL, nomeOuNick) //ai ja mandamos para a API qual a URL completa e qual a Query, para podermos fazer nossa requisição (GET)

	response, erro := requisicoes.FazerRequisicaoComAutenticacao(r, http.MethodGet, url, nil)
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		return
	}
	//Vamos criar um slice de Usuario e vamos popular ele com a resposta da API, ele vai ser um pouco diferente do []modelos.Usuario da API
	var usuarios []modelos.Usuario //aqui para trazer para pagina web, ele vai ter algumas informações que não estavamos trazendo na API por padrão, motivo está relacionado quando estivermos carregando o perfil de cada um dos usuários
	if erro = json.NewDecoder(response.Body).Decode(&usuarios); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuarios.html", usuarios)
}

// CarregarPerfilDoUsuario carrega a página do perfil do usuário
func CarregarPerfilDoUsuario(w http.ResponseWriter, r *http.Request) {
	parametros := mux.Vars(r)
	usuarioID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	cookie, _ := cookies.Ler(r)
	usuarioLogadoID, _ := strconv.ParseUint(cookie["id"], 10, 64)

	if usuarioID == usuarioLogadoID {
		http.Redirect(w, r, "/perfil", 302)
		return
	}

	usuario, erro := modelos.BuscarUsuarioCompleto(usuarioID, r) //Precisamos do request, pois vamos precisar do token para autenticar na API, depois vai distribuir esse token por todas as chamadas
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	utils.ExecutarTemplate(w, "usuario.html", struct {
		Usuario         modelos.Usuario
		UsuarioLogadoID uint64
	}{
		Usuario:         usuario,
		UsuarioLogadoID: usuarioLogadoID,
	})
}
