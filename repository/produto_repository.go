package repository

import "github.com/guibonf1m/apiprodutos/entity"

var todosprodutos []entity.Produto // Slice global que armazena produtos (simula um “banco em memória”)
var id int = 1                     // Inicializa um ID fictício para usos de exemplo

type ProdutoRepository struct {
}

func (r *ProdutoRepository) GetProdutos() []entity.Produto {
	return todosprodutos
}

func (r *ProdutoRepository) GetProduto(id int) entity.Produto {
	for _, v := range todosprodutos {
		if v.ID == id {
			return v
		}
	}

	return entity.Produto{}
}

func (r *ProdutoRepository) GetProdutoPeloNome(name string) entity.Produto {
	for _, v := range todosprodutos {
		if v.Nome == name {
			return v
		}
	}

	return entity.Produto{}
}

func (r *ProdutoRepository) GetProdutoPelaCategoria(categoria string) entity.Produto {
	for _, v := range todosprodutos {
		if v.Nome == categoria {
			return v
		}
	}

	return entity.Produto{}
}

func (r *ProdutoRepository) AddProduto(produto entity.Produto) entity.Produto {
	produto.ID = id
	todosprodutos = append(todosprodutos, produto)

	id++

	return produto
}

func (r *ProdutoRepository) UpdateProduto(produto entity.Produto) entity.Produto {
	for i, v := range todosprodutos {
		if v.ID == produto.ID {
			todosprodutos[i] = produto
		}
	}

	return produto
}

func (r *ProdutoRepository) DeleteProduto(id int) entity.Produto {
	for i, v := range todosprodutos {
		if v.ID == id {
			todosprodutos = append(todosprodutos[:i], todosprodutos[i+1:]...)
		}
	}

	return entity.Produto{}
}
