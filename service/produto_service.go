package service

import (
	"errors"
	"fmt"
	"github.com/guibonf1m/apiprodutos/entity"
	"github.com/guibonf1m/apiprodutos/repository"
)

var categoriasPermitidas = map[string]bool{
	"Alimentos": true,
	"Bebidas":   true,
	"Higiene":   true,
	"Limpeza":   true,
	"Outros":    true,
}

type ProdutoService struct {
	Repo *repository.ProdutoRepository
}

type ProdutoResponse struct {
	ID         int     `json:"id"`
	Nome       string  `json:"nome"`
	Categoria  string  `json:"categoria"`
	Preco      float64 `json:"preco"`
	EmEstoque  bool    `json:"em_estoque"`
	Quantidade int     `json:"quantidade"`
	Desconto   float64 `json:"desconto"`
	PrecoFinal float64 `json:"preco_final"`
}

type CategoriaFiltro struct {
	Nome         *string
	Categoria    *string
	EmEstoque    *bool
	MostrarTodos bool
}

func (p *ProdutoService) ValidarECriarProduto(produto entity.Produto) (entity.Produto, error) {

	if !categoriasPermitidas[produto.Categoria] {
		er := errors.New("Erro: Categoria inválida. Categorias permitidas: Alimentos, Bebidas, Higiene, Limpeza, Outros.")
		return entity.Produto{}, er
	}

	if produto.Preco <= 0 {
		er := errors.New("O produto tem preço inválido.")
		return entity.Produto{}, er
	}

	if produto.Desconto >= 50 {
		er := errors.New("Esse desconto é inválido.")
		return entity.Produto{}, er
	}

	if produto.EmEstoque && produto.Quantidade <= 0 {
		er := errors.New("A quantidade deve ser positiva, se tiver em estoque.")
		return entity.Produto{}, er
	}

	novoProduto := p.Repo.AddProduto(produto)
	return novoProduto, nil
}

func (p *ProdutoService) BuscarPorCategoria(filtroDoUsuario CategoriaFiltro) ([]entity.Produto, error) {

	produtos := p.Repo.GetProdutos()

	// se o usuário não mandou nenhuma categoria (nil) ou mandou uma string vazia ("")
	if filtroDoUsuario.Categoria == nil || *filtroDoUsuario.Categoria == "" {
		// Então: não tem filtro → devolve todos os produtos.
		return produtos, nil
	}

	var filtrados []entity.Produto
	for _, produto := range produtos {

		// Verifica se o usuário passou um filtro de categoria. Se passou (categoria != nil)
		//e o produto atual NÃO pertence à categoria filtrada,

		if filtroDoUsuario.Categoria != nil && produto.Categoria != *filtroDoUsuario.Categoria {
			continue // ignora esse produto e segue pro próximo
		}

		// Verifica se o usuário passou um filtro de estoque.
		// Se passou (emEstoque != nil) e o valor do produto for diferente do filtro,

		if filtroDoUsuario.EmEstoque != nil && produto.EmEstoque != *filtroDoUsuario.EmEstoque {
			continue // ignora esse produto e segue pro próximo
		}

		// Filtro por preço > 10 (somente se não quiser mostrar todos)
		if !filtroDoUsuario.MostrarTodos && produto.Preco <= 10 {
			continue
		}

		// Se nenhum dos continue foi acionado ou seja a categoria está OK e o Estoque OK também.
		filtrados = append(filtrados, produto)
	}

	if len(filtrados) == 0 {
		return nil, errors.New("Nenhum produto encontrado para essa categoria.")
	}
	return filtrados, nil
}

func (p *ProdutoService) AtualizarProdutoPorId(produto entity.Produto) (entity.Produto, error) {

	produtoExistente := p.Repo.GetProduto(produto.ID)

	if !categoriasPermitidas[produto.Categoria] {
		er := errors.New("Erro: Categoria inválida. Categorias permitidas: Alimentos, Bebidas, Higiene, Limpeza, Outros.")
		return entity.Produto{}, er
	}

	if produtoExistente.ID == 0 {
		er := errors.New("Produto não encontrado para atualização.")
		return entity.Produto{}, er
	}

	if produto.Preco <= 0 {
		er := errors.New("O produto tem preço inválido.")
		return entity.Produto{}, er
	}

	if produto.Desconto >= 50 {
		er := errors.New("Esse desconto é inválido.")
		return entity.Produto{}, er
	}

	if produto.EmEstoque && produto.Quantidade <= 0 {
		er := errors.New("A quantidade deve ser positiva, se tiver em estoque.")
		return entity.Produto{}, er
	}

	produtoAtualizado := p.Repo.UpdateProduto(produto)
	return produtoAtualizado, nil
}

func NovoProdutoResponse(produto entity.Produto) ProdutoResponse {

	valorTotal := produto.Preco * float64(produto.Quantidade)
	desconto := valorTotal * produto.Desconto / 100
	precoFinal := valorTotal - desconto

	fmt.Println(precoFinal)

	respostaProduto := ProdutoResponse{
		ID:         produto.ID,
		Nome:       produto.Nome,
		Categoria:  produto.Categoria,
		Preco:      produto.Preco,
		EmEstoque:  produto.EmEstoque,
		Quantidade: produto.Quantidade,
		Desconto:   produto.Desconto,
		PrecoFinal: precoFinal,
	}
	return respostaProduto
}
