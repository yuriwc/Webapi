package controllers

import (
	"fmt"
	"net/http"
	"webApi/models"

	"github.com/gin-gonic/gin"
)

type EnderecoController struct {
	Rua string `json:"Rua" binding:"required"`
	Bairro string `json:"Bairro" binding:"required"`
	Cidade string `json:"Cidade" binding:"required"`
	Estado string `json:"Estado" binding:"required"`
	IdUser uint `json:"IdUser" binding:"required"`
}

func cadastrarEndereco(endereco EnderecoController, c *gin.Context) (models.Endereco, error) {
	var enderecoModel models.Endereco
	enderecoModel.RUA = endereco.Rua
	enderecoModel.Bairro = endereco.Bairro
	enderecoModel.Cidade = endereco.Cidade
	enderecoModel.Estado = endereco.Estado

	result, erro :=  models.CriarEndereco(enderecoModel)
	return result, erro
}

func CriarEndereco(c *gin.Context) {
	var endereco EnderecoController
	err := c.BindJSON(&endereco)
	if err!= nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	result, erro := cadastrarEndereco(endereco, c)
	if erro!= nil {
		fmt.Print(erro.Error())
		return
	}
	
	models.UpdateEnderecoPessoa(int(endereco.IdUser), result.IdEndereco)
	
}