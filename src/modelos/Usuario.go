package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json: "id,omitempty"` //omitempty => caso não tenha o ID ele não montara um JSON, sendo que sem o omitempty montaria um com id = 0
	Nome     string    `json: "nome,omitempty"`
	Nick     string    `json: "nick,omitempty"`
	Email    string    `json: "email,omitempty"`
	Senha    string    `json: "senha,omitempty"`
	CriadoEm time.Time `json: "CriadoEm,omitempty"`
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O Nickname é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email é obrigatório e não pode estar em branco")
	}

	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O e-mail inserido é inválido")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A Senha é obrigatório e não pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}

		usuario.Senha = string(senhaComHash)
	}

	return nil
}

// Preparar vai chamar os métodos validar e formatar do struct Usuario do usuario recebido. etapa string => Serve para saber se vai atualizar ou inserir um usuario novo.
func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}

	return nil
}
