package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

/*
// init roda antes da func main
func init() {
	chave := make([]byte, 64)

	//poupulando 64 valores aleatórios na var chave
	if _, erro := rand.Read(chave); erro != nil {
		log.Fatal(erro)
	}

	//convertendo o slice de byte em base64
	stringBase64 := base64.StdEncoding.EncodeToString(chave)

	// valor aleatório para ser a chave do token = SECRET_KEY no .env
	fmt.Println(stringBase64)
}

*/

func main() {
	config.Carregar()
	r := router.Gerar()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
