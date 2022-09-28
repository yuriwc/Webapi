package controllers

import (
	"net/http"
	"webApi/models"

	"github.com/gin-gonic/gin"
	"github.com/paemuri/brdoc"
)

type PessoaReq struct {
	Nome       string `json:"Nome" binding:"required"`
	Telefone   string `json:"Telefone" binding:"required"`
	CPF        string `json:"CPF" binding:"required"`
	IdEndereco int    `json:"idEndereco"`
}

func cadastrarPessoa(pessoa PessoaReq, c *gin.Context) {
	var pessoaModel models.Pessoa
	pessoaModel.Nome = pessoa.Nome
	pessoaModel.Telefone = pessoa.Telefone
	pessoaModel.CPF = pessoa.CPF
	models.CreatePessoa(pessoaModel)
	c.JSON(http.StatusOK, gin.H{"success": "True", "message": "Usuário criado com sucesso"})
}

func CreatePessoa(c *gin.Context) {
	var pessoa PessoaReq
	err := c.ShouldBindJSON(&pessoa)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}
	if brdoc.IsCPF(pessoa.CPF) == false {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": "CPF inválido"})
		return
	}

	cadastrarPessoa(pessoa, c)
}

func GetPessoaByNumberController(c *gin.Context) {
	result := models.GetPessoaByNumber(c.Param("numero"))
	if len(result.Nome) > 0 {
		c.JSON(http.StatusOK, gin.H{"success": "True", "data": result})
	} else {
		c.JSON(http.StatusOK, gin.H{"success": "True", "message": "Não foi encontrado nenhuma pessoa com esse número"})
	}
}
