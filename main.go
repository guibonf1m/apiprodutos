package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"github.com/guibonf1m/apiprodutos/nomedoseumodulo/controller"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("", index)
	r.GET("Produtos", controller.GetProduto)
	r.POST("Produtos", controller.AddProduto)

	r.Run(":8080")

}

func index(c *gin.Context) {
	c.JSON(http.StatusOK, "Bem vindo a minha primeira API!")
}
