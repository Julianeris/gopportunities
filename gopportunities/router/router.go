package router

import "github.com/gin-gonic/gin"

func Initialize() *gin.Engine {

	// Criação do Gin Router padrão
	router := gin.Default()
	// Criação de uma rota GET para a rota /ping
	router.GET("/ping", func(c *gin.Context) {
		//Criação de uma resposta json dentro do context Gin
		c.JSON(200, gin.H{
			"message": " Pong",
		})
	})
	// Inicia o servidor na porta 8080
	router.Run(":8080")
	// Retorna o router
	return router
}
