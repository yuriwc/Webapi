package controllers

import (
	"database/sql"
	"net/http"
	"webApi/models"

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

func CadastrarUsuario(usuario UserController, c *gin.Context) {
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

	CadastrarUsuario(user, c)
}

func GetUsuario(c *gin.Context) {
	var login LoginController
	err := c.BindJSON(&login)
  if err!= nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	result, err := models.GetUser(login.NumeroCelular, login.Senha)

	if err!= nil {
		if err.Error() == "sql: no rows in result set" {
			c.JSON(http.StatusOK, gin.H{"success": "True", "message": "Nenhum usuário foi encontrado" })
			return 
		} else {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
		}
	}
	c.JSON(http.StatusOK, gin.H{"success": "True", "data": result })

}