package repositorios

import (
	"api/src/modelos"
	"database/sql"
)

type Usuarios struct {
	db *sql.DB
}

// Recebe um banco que é aberto pelo controller e o controller que irá chamar essa função
// NovoRepositorioDeUsuario joga o banco em um struct de usuario criando um repositório de usuários
func NovoRepositorioDeUsuario(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

/*
Parametros:
(repositorio Usuario) => Criação do repositorio Usuario.
Criar(usuario  modelos.Usuario) => Metodo que recebe um modelo Usuario como parametro
(uint64, error) => Retorna um int e um erro, retorna um int pois irá retornar o ID do usuario inserido.

obs: uint NUNCA terá valor negativo.
*/
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, erro := repositorio.db.Prepare("insert into usuarios  (nome, nick, email, senha) values(?, ?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	// somente convertando para uint64 o ID e retornando nil pq n tem erro.
	return uint64(ultimoIDInserido), nil
}
