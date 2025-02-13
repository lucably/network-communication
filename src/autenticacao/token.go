package autenticacao

//Criamos um ALIAS para a chamada do "github.com/dgrijalva/jwt-go" como jwt

import (
	"api/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// CriarToken cria o token com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	//MapClaims => Cria um conjunto de regras para o jwt
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix() //time.Hour * 6 => expira em 6 horas; Unix() => Passa para milisegundos
	permissoes["usuarioId"] = usuarioID

	//Criação da assinatura, 1° Metodo da assinatura e 2° as permissoes
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte(config.SecretKey))
}
