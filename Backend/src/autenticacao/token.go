package autenticacao

//Criamos um ALIAS para a chamada do "github.com/dgrijalva/jwt-go" como jwt

import (
	"api/src/config"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
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

// extrairToken funcao privada para validar se o token veio no header.
func extrairToken(r *http.Request) string {
	token := r.Header.Get("Authorization")

	//Verificando se o token possui 2 tamanhos(Bearer Token)
	if len(strings.Split(token, " ")) == 2 {
		if strings.Split(token, " ")[0] == "Bearer" {
			return strings.Split(token, " ")[1]
		}
	}
	return ""
}

// retornarChaveDeVerificacao valida se o tipo da assinatura do token é valido.
func retornarChaveDeVerificacao(token *jwt.Token) (interface{}, error) {
	//Verificando se o método de assinatura do jwt é válido.
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Método de assinatura insesperado! %v", token.Header["alg"])
	}
	return config.SecretKey, nil
}

func ExtrairUsuarioID(r *http.Request) (uint64, error) {
	tokenString := extrairToken(r)

	//Passando a função deste jeito => retornarChaveDeVerificacao, ele ja passa a variavel "token" direto, sem precisar fazer retornarChaveDeVerificacao(token)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return 0, erro
	}

	if permissoes, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		usuarioID, erro := strconv.ParseUint(fmt.Sprintf("%.0f", permissoes["usuarioId"]), 10, 64)
		if erro != nil {
			return 0, erro
		}
		return usuarioID, nil
	}
	return 0, errors.New("Token inválido")
}

// ValidarToken verifica se o token passado e valido.
func ValidarToken(r *http.Request) error {
	tokenString := extrairToken(r)

	//Passando a função deste jeito => retornarChaveDeVerificacao, ele ja passa a variavel "token" direto, sem precisar fazer retornarChaveDeVerificacao(token)
	token, erro := jwt.Parse(tokenString, retornarChaveDeVerificacao)
	if erro != nil {
		return erro
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return errors.New("Token inválido!")
}
