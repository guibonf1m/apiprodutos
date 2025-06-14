package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/guibonf1m/apiprodutos/controller"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("", index)
	r.GET("Produtos", controller.ListarProdutos)
	r.POST("Produtos", controller.AddProduto)
	r.PATCH("Produtos", controller.UpadateProductId)
	r.DELETE("Produtos", controller.DeletarProdutoPorId)

	r.Run(":8080")

}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo a minha primeira API!")
}
