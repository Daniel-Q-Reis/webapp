package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// Erro representa a resposta de erro da API
type ErroAPI struct {
	Erro string `json:"erro"`
}

// JSON retorna uma resposta em formato JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if erro := json.NewEncoder(w).Encode(dados); erro != nil { //da um json enconde nos dados que estão vindo na requisição
		log.Fatal(erro) // em produção posso querer usar um log.Panic ou um mecanismo de recuperação
	}
}

// TratarStatusCodeDeErro trata as requisições com status code 400 ou superior
func TratarStatusCodeDeErro(w http.ResponseWriter, r *http.Response) { //essa função faz sentido pois nem sempre retorna um erro o programa, e sim somente um statuscode, onde 400+ é porque algo foi negado
	var erro ErroAPI
	json.NewDecoder(r.Body).Decode(&erro)
	JSON(w, r.StatusCode, erro)
} //com essa função agora podemos ir no controllers e fazer o tratamento do statuscode >=400
