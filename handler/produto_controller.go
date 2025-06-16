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
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: h.Repo.GetProdutos(),
	})
}

func (h *ProdutoHandler) GetProduto(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "o parametro nao e um numero",
		})
		return
	}

	produto := h.Repo.GetProduto(id)
	if produto.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "produto não existe",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produto,
	})
}

func (h *ProdutoHandler) GetProdutoPeloNome(c *gin.Context) {
	nome := c.Param("name")

	produto := h.Repo.GetProdutoPeloNome(nome)
	if produto.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "Produto não existe",
		})
		return
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produto,
	})
}

func (h *ProdutoHandler) GetProdutoPelaCategoria(c *gin.Context) {
	categoria := c.Param("category")

	produto := h.Repo.GetProdutoPelaCategoria(categoria)
	if produto.ID == 0 {
		c.JSON(http.StatusNotFound, ResponseInfo{
			Error:  true,
			Result: "Produto não existe",
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

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: produtoCriado,
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

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: produtoAtualizado,
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
