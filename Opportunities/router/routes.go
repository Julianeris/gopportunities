package router

import (
	"gopportunities/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {

	//Criação do grupo principal de rotas do projeto
	v1 := router.Group("/api/v1")
	{
		//Criação do CRUD
		v1.GET("/opening", handler.ShowOpeningHandler)

		v1.GET("/openings", handler.ListOpeningsHandler)

		v1.POST("/opening", handler.CreateOpeningHandler)

		v1.DELETE("opening", handler.DeleteOpeningHandler)

		v1.PUT("opening", handler.UpdateOpeningHandler)
	}

}
