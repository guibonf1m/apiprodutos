package controller // Define o pacote responsável pelos controllers (camada web da API)

import (
	"encoding/json"                           // Permite converter dados entre JSON e structs nativos de Go
	"github.com/gin-gonic/gin"                // Framework utilizado para facilitar o roteamento e a resposta HTTP
	"github.com/guibonf1m/apiprodutos/entity" // Importa a definição da struct Produto
	"net/http"                                // Fornece constantes e funções do protocolo HTTP    // Converte strings para outros tipos, útil ao lidar com query params etc
	"strconv"
)

var todosprodutos []entity.Produto // Slice global que armazena produtos (simula um “banco em memória”)
var id int = 1                     // Inicializa um ID fictício para usos de exemplo

type ResponseInfo struct { // Struct que define o modelo padrão de resposta (com campo de erro e resultado)
	Error  bool `json:"error"`
	Result any  `json:"result"`
}

func ListarProdutos(c *gin.Context) { // Handler para listar todos os produtos
	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,         // Indica ausência de erro
		Result: todosprodutos, // Dados de resposta são todos os produtos salvos em memória
	})
}

func AddProduto(c *gin.Context) { // Handler para adicionar um novo produto via POST
	var produto entity.Produto                              // Struct para receber e armazenar os dados do body
	err := json.NewDecoder(c.Request.Body).Decode(&produto) // Decodifica o JSON do request para o struct
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,        // Indica erro na entrada do usuário
			Result: err.Error(), // Mensagem descritiva do erro ocorrido
		})
		return // Interrompe a execução caso houve erro de parsing
	}

	if produto.Nome == "" || produto.Categoria == "" || produto.Preco == 0 ||
		produto.Quantidade == 0 {
		// Retorna erro ao cliente se algum campo obrigatório está faltando
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: "Todos os campos sao obrigatorios",
		})
		return // Interrompe o cadastro se houver campos inválidos
	}

	// Se passar pela validação, atribui ID e adiciona ao slice de produtos
	produto.ID = id
	todosprodutos = append(todosprodutos, produto)

	id++

	c.JSON(http.StatusCreated, ResponseInfo{
		Error:  false,
		Result: "Criado com sucesso",
	})
}

func BuscarProdutoPorId(c *gin.Context) {

}

func UpadateProductId(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	var produto entity.Produto
	err = json.NewDecoder(c.Request.Body).Decode(&produto)
	if err != nil {
		c.JSON(http.StatusBadRequest, ResponseInfo{
			Error:  true,
			Result: err.Error(),
		})
		return
	}

	for i, v := range todosprodutos {
		if v.ID == id {
			todosprodutos[i] = produto
		}
	}

	c.JSON(http.StatusOK, ResponseInfo{
		Error:  false,
		Result: todosprodutos,
	})

}
func DeletarProdutoPorId(c *gin.Context) {

}
