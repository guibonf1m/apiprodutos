package entity

type Produto struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Categoria  string  `json:"categoria"`
	Preco      float64 `json:"preco"`
	EmEstoque  bool    `json:"em_estoque"`
	Quantidade int     `json:"quantidade"`
	Desconto   float64 `json:"desconto"` // percentual de 0 a 100
}
