package requisicoes

import (
	"io"
	"net/http"
	"webapp/src/cookies"
)

// FazerRequisicaoComAutenticacao é utilizada para colocar o token na requisição
func FazerRequisicaoComAutenticacao(r *http.Request, metodo, url string, dados io.Reader) (*http.Response, error) {
	//Prametros = criar uma requisição -> depois passar o token no cabeçalho dessa requisição -> depois teremos um client Http para fazer essa requisição de fato para gente

	//criando a requisição
	request, erro := http.NewRequest(metodo, url, dados) //esse newRequest e diferente do parametro r da nossa função, nosso r vai ser usado somente para ler o cookie e nao fazer a requisição //3 parametros da nossa função
	if erro != nil {                                     //aqui cria a requisição, mas não vai chamar a API, vamos precisar de um CLient http para fazer essa requisição e, ela somente vai poder ser feita depois que o cookie ja estiver no header
		return nil, erro
	}

	//proximo passo = ler o cookie, lendo ignorando o erro
	//se chegamos até aqui, já passei pelo middleware e, ele já indentificou se houve ou não erro, logo não preciso de novo tratamento de erro
	cookie, _ := cookies.Ler(r)
	request.Header.Add("Authorization", "Bearer "+cookie["token"]) //o Bearer e onde a gente insere nosso token de segurança (lembrar do postman)

	//Criando o client do HTTP, que finalmente vai fazer essa requisição de fato
	client := &http.Client{}
	response, erro := client.Do(request) //Do = faça uma requisição e retorna um response
	if erro != nil {
		return nil, erro
	}
	//essa função FazerRequisição, conforme o nome sugere, vai apenas fazer a requisição, ela não vai tratar a resposta, quem irá fazer isso é quem chamar essa função
	//um exemplo de quem irá usar essa função é em paginas.go a função CarregarPaginaPrincipal
	return response, nil
}
