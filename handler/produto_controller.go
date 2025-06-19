package handler

import (
	"encoding/json"
	"github.com/guibonf1m/apiprodutos/repository"
	"github.com/guibonf1m/apiprodutos/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/guibonf1m/apiprodutos/entity"
)

type ResponseInfo struct {
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

type ProdutoHandler struct {
	Repo    *repository.ProdutoRepository
	Service *service.ProdutoService
}

func (h *ProdutoHandler) GetProdutos(c *gin.Context) {
	nome := c.Param("nome")
	emEstoqueRecebido := c.Query("em_estoque")
	categoria := c.Query("categoria")
	mostrarTodosRecebido := c.Query("mostrar_todos")

	var emEstoqueBool bool
	var aplicarFiltrosEstoque bool

	if emEstoqueRecebido != "" {
		var err error
		emEstoqueBool, err = strconv.ParseBool(emEstoqueRecebido)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseInfo{
				Error:  true,
				Result: "valor inválido para o filtro em_estoque, use (true ou false)",
			})
			return
		}
		aplicarFiltrosEstoque = true
	}

	mostrarTodos := false

	if mostrarTodosRecebido != "" {
		var err error
		mostrarTodos, err = strconv.ParseBool(mostrarTodosRecebido)
		if err != nil {
			c.JSON(http.StatusBadRequest, ResponseInfo{
				Error:  true,
				Result: "valor inválido para o filtro mostrar_todos, use (true ou false)",
			})
			return
		}

	}

	filtro := service.CategoriaFiltro{
		MostrarTodos: mostrarTodos,
	}

	if nome != "" {
		filtro.Nome = &nome
	}
	if categoria != "" {
		filtro.Categoria = &categoria
	}
	if aplicarFiltrosEstoque {
		filtro.EmEstoque = &emEstoqueBool
	}

	produtos, err := h.Service.BuscarPorCategoria(filtro)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produtos,
	})
}

func (h *ProdutoHandler) GetProduto(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro não é um número, tente novamente.",
		})
		return
	}

	produto := h.Repo.GetProduto(id)
	if produto.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "produto não existe, tente novamente.",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produto,
	})
}

func (h *ProdutoHandler) AddProduto(c *gin.Context) {
	var produto entity.Produto

	err := json.NewDecoder(c.Request.Body).Decode(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	produtoCriado, err := h.Service.ValidarECriarProduto(produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}
	produtoResponse := service.NovoProdutoResponse(produtoCriado)

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: produtoResponse,
	})
}

func (h *ProdutoHandler) UpdateProduto(c *gin.Context) {
	idParam := c.Param("id")
	var produto entity.Produto

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	err = json.NewDecoder(c.Request.Body).Decode(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	produto.ID = id
	produtoAtualizado, err := h.Service.AtualizarProdutoPorId(produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	produtoResponse := service.NovoProdutoResponse(produtoAtualizado)
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produtoResponse,
	})
}

func (h *ProdutoHandler) DeleteProduto(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	h.Repo.DeleteProduto(id)

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: "deletado com sucesso",
	})
}
