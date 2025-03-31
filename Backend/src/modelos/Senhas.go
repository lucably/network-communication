package modelos

// Senha represta o formato da requisição de alteração de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual`
}
