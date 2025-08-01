package modelos

import (
	"time"
)

// Usuario representa uma pessoa utilizando a rede social
type Usuario struct {
	ID          uint64       `json:"id"`
	Nome        string       `json:"nome"`
	Email       string       `json:"email"`
	Nick        string       `json:"nick"`
	CriadoEm    time.Time    `json:"criadoEm"`
	Seguidores  []Usuario    `json:"seguidores"`  //quando eu carregar o perfil de um usuario, vou querer ver quem segue ele, logo já vou trazer essa informação
	Seguindo    []Usuario    `json:"seguindo"`    // na tela dentro aqui da propria struct de usuário, vamos ter uma pagina que vai renderizar essas informações,
	Publicacoes []Publicacao `json:"publicacoes"` //logo faz sentido deixar-mos todas essas informações dentro de uma mesma struct
}
