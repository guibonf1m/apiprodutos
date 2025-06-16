package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/guibonf1m/apiprodutos/handler"
	"github.com/guibonf1m/apiprodutos/repository"
	"github.com/guibonf1m/apiprodutos/service"

	"net/http"
)

func main() {
	r := gin.Default()

	ProdutoService := &service.ProdutoService{
		Repo: &repository.ProdutoRepository{},
	}

	produtoHandler := &handler.ProdutoHandler{
		Service: ProdutoService,
		Repo:    &repository.ProdutoRepository{},
	}

	r.GET("", index)
	r.GET("/produtos", produtoHandler.GetProdutos)
	r.POST("/produtos", produtoHandler.AddProduto)
	r.GET("/produtos/:id", produtoHandler.GetProduto)
	r.PUT("/produtos/:id", produtoHandler.UpdateProduto)
	r.DELETE("/produtos/:id", produtoHandler.DeleteProduto)

	r.Run(":8080")
}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo a minha primeira API!")
}
