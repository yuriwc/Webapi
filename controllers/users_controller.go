package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"webApi/models"
	"webApi/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	IdPessoa int `json:"IdPessoa"`
	IdNivel sql.NullInt16
	NumeroCelular string `json:"NumeroCelular" binding:"required"`
	Senha string `json:"Senha" binding:"required"`
}

type LoginController struct {
	NumeroCelular string `json:"NumeroCelular" binding:"required"`
	Senha string `json:"Senha" binding:"required"`
}

func cadastrarUsuario(usuario UserController, c *gin.Context) {
	var userModel models.User

	userModel.NumeroCelular = usuario.NumeroCelular
	userModel.Senha = usuario.Senha
	userModel.IdPessoa = usuario.IdPessoa
	_, err := models.CreateUser(userModel)
	if err!= nil {
    c.JSON(http.StatusOK, gin.H{"success": "False", "message": err.Error() })
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": "True", "message": "Usuário criado com sucesso" })
}

func CriarUsuario(c *gin.Context) {
	var user UserController
	err := c.BindJSON(&user)
	if err!= nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	user.Senha = services.SHA256Enconder(user.Senha)
	fmt.Print(user.Senha)	
 
	cadastrarUsuario(user, c)
}

func GetUsuario(c *gin.Context) {
	var login LoginController
	err := c.BindJSON(&login)
  if err!= nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	senha, err := models.GetUserByUsername(login.NumeroCelular)

	if senha.Senha != services.SHA256Enconder(login.Senha) {
		c.JSON(400, gin.H{
			"ERROR": "Senha inválida",
		})
		return
	}

	token, err := services.NewJWTService("secret-key", "web-api").GenerateToken(1)

	if err != nil {
		c.JSON(500, gin.H {
			"ERROR": err.Error(),
		})
		return 
	}

	c.JSON(200, gin.H{
		"token": token,
	})

	/* result, err := models.GetUser(login.NumeroCelular, login.Senha)

	if err!= nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusOK, gin.H{"success": "True", "message": "Nenhum usuário foi encontrado" })
			return 
		} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": "True", "data": result }) */
}