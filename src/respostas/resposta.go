package respostas

import (
	"encoding/json"
	"log"
	"net/http"
)

// dados interface{} => uma interface generica
// JSON => retorna uma resposta em JSON para a requisição
func JSON(w http.ResponseWriter, statusCode int, dados interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// Essa validação refere-se ao Atualizar, pois não tem retorno no body "dados == nil", esse metodo só serve para os outros que possui retorno
	if dados != nil {
		if erro := json.NewEncoder(w).Encode(dados); erro != nil {
			//Aqui sim, caso de algum erro, deve-se matar a aplicação.
			log.Fatal(erro)
		}
	}
}

// Erro retorna um erro em formato JSON
func Erro(w http.ResponseWriter, statusCode int, erro error) {

	// Uma struct que tem o campo Erro e retorna o Erro
	// campo=> Erro string `json:"erro"` retorno=> Erro: erro.Error(), pois o erro tem um interface que tem um metodo que retorna um erro
	JSON(w, statusCode, struct {
		Erro string `json:"erro"`
	}{
		Erro: erro.Error(),
	})
}
