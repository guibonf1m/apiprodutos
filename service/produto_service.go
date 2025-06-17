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

type FiltroProduto struct {
	Nome      *string
	Categoria *string
	EmEstoque *bool
}

func (p *ProdutoService) ValidarECriarProduto(produto entity.Produto) (entity.Produto, error) {

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

func (p *ProdutoService) BuscarPorLista(filtrosDoUsuario FiltroProduto) ([]entity.Produto, error) {

	produtos := p.Repo.GetProdutos()

	for _, v := range produtos {
		if filtrosDoUsuario.EmEstoque != nil {
			if v.EmEstoque != *filtrosDoUsuario.EmEstoque {
				er := errors.New("Categoria não encontrada, para esse produto.")
				return entity.Produto{}, er
			}
		}
	}
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
