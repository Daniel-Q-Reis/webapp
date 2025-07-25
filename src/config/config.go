package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// APIURL representa a URL  para comunicação com a API
	APIURL = ""
	// Porta onde a aplicação web está rodando
	Porta = 0
	// HashKey é utilizada para autenticar o cookie
	HashKey []byte
	// BlockKey é utilizada para criptografar os dados do cookie gerando um nivel a mais de segurança
	BlockKey []byte
)

//SOMENTE APOS CRIAR O ARQUIVO .env na raiz do programa eu posso popular essas 4 variaveis acima

// Carregar inicializa as variáveis de ambiente
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	//***SOMENTE APOS CRIAR O ARQUIVO .env na raiz do programa eu posso popular essas 4 variaveis acima***
	//Começando pela porta, pois é a unica que pode dar erro
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))
} //agora podemos chamar essa função la dentro do main
