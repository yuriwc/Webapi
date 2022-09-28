package controllers

import (
	"fmt"
	"net/http"
	"webApi/models"

	"github.com/gin-gonic/gin"
)

type NumeroBilheteController struct {
	IdBilhete uint `json:"idBilhete" binding:"required"`
	Numero    uint `json:"numero" binding:"required"`
}

type SearchNumeroBilhete struct {
	Numero uint `json:"idBicho" binding:"required"`
}

func saveNumeroBilheteToDB(numeroBilhete NumeroBilheteController, c *gin.Context) (models.NumeroBilhete, error) {
	var numeroBilheteDB models.NumeroBilhete
	numeroBilheteDB.IdBilhete = numeroBilhete.IdBilhete
	numeroBilheteDB.Numero = numeroBilhete.Numero

	result, erro := models.CriarNumeroBilhetes(numeroBilheteDB)
	c.JSON(200, gin.H{
		"numeroBilhete": numeroBilhete,
	})
	return result, erro
}

func CriarNumeroBilhete(c *gin.Context) {
	var numeroBilhete NumeroBilheteController
	err := c.BindJSON(&numeroBilhete)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	_, erro := saveNumeroBilheteToDB(numeroBilhete, c)
	if erro != nil {
		fmt.Println(erro.Error())
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    numeroBilhete,
	})
}

func GetAllNumbersInsertedByIDBichoFromDB(c *gin.Context) {
	var numero SearchNumeroBilhete
	err := c.BindJSON(&numero)
	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	result := models.GetAllNumbersInsertedByIdBicho(numero.Numero)
	c.JSON(http.StatusOK, gin.H{"success": "True", "data": result})

	return
}
