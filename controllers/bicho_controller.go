package controllers

import (
	"fmt"
	"net/http"
	"webApi/models"

	"github.com/gin-gonic/gin"
)

type BichoController struct {
	DataBicho     string `json:"DataBicho" binding:"required"`
	Concurso      int    `json:"Concurso" binding:"required"`
	IdPessoaBicho uint   `json:"idPessoaBicho" binding:"required"`
}

func saveToDB(bicho BichoController, c *gin.Context) (models.Bicho, error) {
	var bichoDB models.Bicho

	bichoDB.DataBicho = bicho.DataBicho
	bichoDB.Concurso = bicho.Concurso
	bichoDB.IdPessoaBicho = bicho.IdPessoaBicho

	result, erro := models.CriarBicho(bichoDB)
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
	}
	return result, erro
}

func CriarBicho(c *gin.Context) {
	var bicho BichoController
	err := c.BindJSON(&bicho)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	_, erro := saveToDB(bicho, c)
	if erro != nil {
		fmt.Println(erro.Error())
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    bicho,
	})
}

func GetAllBichos(c *gin.Context) {
	bichos, erro := models.GetAllBichos()
	if erro != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": erro.Error()})
	}
	c.JSON(200, gin.H{
		"success": true,
		"data":    bichos,
	})
}
