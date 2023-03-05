package routes

import (
	"goapp/controllers"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()

	acai := r.Group("/acai")
	{
		acai.POST("/", controllers.CreateAcai)
		acai.GET("/", controllers.ListAcais)
		acai.GET("/findById/:id", controllers.ListAcaisById)
		acai.PATCH("/:id", controllers.UpdateAcai)
		acai.DELETE("/:id", controllers.DeleteAcai)
	}

	adicionais := r.Group("/adicionais")
	{
		adicionais.POST("/", controllers.CreateAdicional)
		adicionais.GET("/", controllers.ListAdicionais)
		adicionais.PATCH("/:id", controllers.UpdateAdicional)
		adicionais.DELETE("/:id", controllers.DeleteAdicional)
	}

	r.Run()
}
