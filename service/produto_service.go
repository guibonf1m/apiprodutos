package service

import (
	"errors"
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

type CategoriaFiltro struct {
	Nome      *string
	Categoria *string
	EmEstoque *bool
}

func (p *ProdutoService) ValidarECriarProduto(produto entity.Produto) (entity.Produto, error) {

	if !categoriasPermitidas[produto.Categoria] {
		er := errors.New("Categoria inválida.")
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

func (p *ProdutoService) BuscarPorCategoria(filtro CategoriaFiltro) ([]entity.Produto, error) {

	produtos := p.Repo.GetProdutos()

	if filtro.Categoria == nil || *filtro.Categoria == "" {
		return produtos, nil
	}

	var filtrados []entity.Produto
	for _, produto := range produtos {
		if filtro.Categoria != nil && produto.Categoria == *filtro.Categoria {
			filtrados = append(filtrados, produto)
		}
	}
	if len(filtrados) == 0 {
		return nil, errors.New("Nenhum produto encontrado para essa categoria.")
	}
	return filtrados, nil
}

func (p *ProdutoService) AtualizarProdutoPorId(produto entity.Produto) (entity.Produto, error) {

	produtoExistente := p.Repo.GetProduto(produto.ID)

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
