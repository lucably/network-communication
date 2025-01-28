package modelos

import "time"

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"` //omitempty => caso não tenha o ID ele não montara um JSON, sendo que sem o omitempty montaria um com id = 0
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`
}
