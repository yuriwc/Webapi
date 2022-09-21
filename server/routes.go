package server

import (
	"webApi/controllers"
	"webApi/server/middlewares"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		user := main.Group("user")
		{
			user.POST("/", controllers.CriarUsuario)
			//r.POST("/createPessoa", controllers.CreatePessoa)
			//r.POST("/createUser", controllers.CriarUsuario)
			user.POST("/getUser", controllers.GetUsuario)
			//r.GET("/getPessoa/:numero", controllers.GetPessoaByNumberController)
		}
		pessoa := main.Group("pessoa", middlewares.Auth())
		{
			pessoa.POST("/createPessoa", controllers.CreatePessoa)
			pessoa.POST("/createAddress", controllers.CriarEndereco)
		}
	}

	return router
}