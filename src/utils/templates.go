package utils

import (
	"html/template"
	"net/http"
)

var templates *template.Template

// CarregarTemplates deve ser executada antes da função a baixo, para tal iremos chamar ela no pacote main = utils.Carregartemplates()

// CarregarTemplates insere os templates html na variável templates
func CarregarTemplates() {
	//Aqui pegaremos os arquivos que estão dentro da pasta views do nosso programa e pode pegar os arquivos que sejam qqr coisa *.html
	templates = template.Must(template.ParseGlob("views/*.html"))            //template. é o package
	templates = template.Must(templates.ParseGlob("views/templates/*.html")) //aqui esse templates.parse já é referenciado a variavel que agora já possui valores, é quase igual um append isso aqui
}

// ExecutarTemplate renderiza uma página html na tela
func ExecutarTemplate(w http.ResponseWriter, template string, dados any) { //template no caso e o nome do arquivo para ser renderizado, e dados são os dados que serão renderizados na tela
	templates.ExecuteTemplate(w, template, dados) //basicamente estamos embrulhando essa função em uma chamada da função que criamos

}
