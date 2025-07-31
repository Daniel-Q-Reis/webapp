package cookies

import (
	"net/http"
	"time"
	"webapp/src/config"

	"github.com/gorilla/securecookie"
)

// Esse tipo será utilizado para codificar e decodificar os valores colocados no browser
var s *securecookie.SecureCookie // = securecookie.New() como estão em variaveis de ambiente, não poderá fazer essa chamada diretamente aqui, por isso vamos criar a função abaixo

// Configurar utiliza as variáveis de ambiente para a criação do SecureCookie
func Configurar() { // vai ser chamado la na função main -> será utilizado para codificar e decodificar os valores colocados no browser
	s = securecookie.New(config.HashKey, config.BlockKey) // Se olharmos na descrição da func New(), iremos ver que para HashKey é recomendado que usamos uma key de 32 ou 64 bytes de tamanho
} // Para blockKey recomendado tamanho de 16, 24 ou 32bytes

// Salavar regitra as informações de autenticação
func Salvar(w http.ResponseWriter, ID, token string) error {
	//Criação dos dados através do map
	dados := map[string]string{ //map de chaves string e valores string
		"id":    ID,
		"token": token,
	}

	//chamando metodo s.Encode (que retorna dois valores)
	dadosCodificados, erro := s.Encode("dados", dados) //Aqui os dados ficam no formato que eu quero, pois estou usando o s, atravez do securecookie
	if erro != nil {                                   // Que conforme vimos em Configurar() acima, eles já possuem a Hashkey e a blockKey
		return erro // Com isso já vai estar todo mundo codificado e já criptografado (dados)
	}

	//Colocando o cookie no browser
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    dadosCodificados,
		Path:     "/",  //aqui é barra pois queremos ele funcionando em todos os lugares
		HttpOnly: true, // Esse parametro aqui, serve para mitigarmos o risco de esse cookie ser acessado pelo lado do cliente
	})
	return nil
}

//Proximo passo agora é criar uma função para ler e converter esses valores salvos no browser, para que a gente consiga trabalhar com elas
//A rota de login e a rota de criação de usuário, são as unicas duas rotas da nossa aplicação que não vão precisar de autenticação, e por isso estão prontas: Na API também era assim
//Todas as demais rotas vão precisar de autenticação então Agora vamos criar novas rotas além do login, que vao utilizar essa forma de validação, juntamente com o pacote middlewares que será implementado
//vamos criar agora o arquivo home.html, que vai ser a pagina home, vamos criar uma rota para tratar disso e vai ser uma rota que vai precisar de autenticação
//rotas/home.go
//views/home.html

// Ler retorna os valores armazenados no cookie
func Ler(r *http.Request) (map[string]string, error) { // esse primeiro valor são os valores que estão no cookie, criado na func salvar acima, usando o map, ID token
	cookie, erro := r.Cookie("dados") //conforme na func salvar() acima atravez do http.SetCookie, teremos os "dados" o nome é importante estar igual
	if erro != nil {                  //"dados" é bem importante, usamos ele para codificar, salvar, descodificar, para puxar do browser, etc, como podemos ter varios cookies no browser e importante indentificar ele com o nome correto
		return nil, erro
	}
	//Aqui eu já tenho o cookie, mas ele ainda está com os dados codificados
	//Para descodificar vou usar o pacote secure cookie, e vou criar a variavel valores

	valores := make(map[string]string)                                 //aqui estou alocando um map vazio e jogando na memoria da variavel valores
	if erro = s.Decode("dados", cookie.Value, &valores); erro != nil { //&valores é o endereco de memoria de valores, é onde vou colocar os valores
		return nil, erro //novamente "dados" é para o navegador saber de qual cookie estamos falando
	}

	return valores, nil //se finalmente chegar até aqui retornamos os cookies que agora podem ser lidos
}

// Deletar remove os valores armazenados no cookie
func Deletar(w http.ResponseWriter) { //aqui não precisa de request, somente do response
	//Aqui basicamente estaremos sobrescrevendo o que foi setado no salvar, e o tornando um valor em branco, assim o usuário irá perder o acesso ao token, podendo assim efetuar o logout
	http.SetCookie(w, &http.Cookie{
		Name:     "dados",
		Value:    "", //agora o valor é em branco
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Unix(0, 0), //tempo de expiração de cookie, não foi passado na criação (assim ele fica com valor padrão se não passar nada), aqui o tempo como zero, 0,0, significa que ele já está expirado
		//consequentimente limpando o cookie
	})
}
