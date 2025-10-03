package router

import "github.com/gin-gonic/gin"

func Initialize() *gin.Engine {

	// Criação do Gin Router padrão
	router := gin.Default()
	// Inicia as rotas
	initializeRoutes(router)
	// Inicia o servidor na porta 8080
	router.Run(":8080")
	// Retorna o router
	return router
}
