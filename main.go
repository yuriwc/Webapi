package main

import (
	"net/http"
	"webApi/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
  r := gin.Default()

  r.GET("/", func(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"data": "hello world"})    
  })

	r.POST("/createPessoa", controllers.CreatePessoa)
	r.POST("/createUser", controllers.CriarUsuario)
	r.POST("/getUser", controllers.GetUsuario)
	r.GET("/getPessoa/:numero", controllers.GetPessoaByNumberController)

  r.Run()
}