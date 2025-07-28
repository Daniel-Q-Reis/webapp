package modelos

import "time"

// Publicacao representa uma publicação feita por um usuário
type Publicacao struct {
	ID        uint64    `json:"id,omitempty"`
	Titulo    string    `json:"titulo,omitempty"`
	Conteudo  string    `json:"conteudo,omitempty"`
	AutorID   uint64    `json:"autorId,omitempty"`
	AutorNick string    `json:"autorNick,omitempty"`
	Curtidas  uint64    `json:"curtidas"`
	CriadaEm  time.Time `json:"criadaEm,omitempty"`
}

//Aqui todos os campos estão sendo mapeados para json, pois vamos trabalhar com json
//Em publicação não teremos nenhum metodo, pois as questões de formatação serão feitas quando formos cadastrar uma publicação a gente vai usar um map, assim como na criação de usuarios
