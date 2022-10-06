package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"webApi/models"

	"github.com/gin-gonic/gin"
)

type BilheteController struct {
	IdBicho         uint `json:"idBicho" binding:"required"`
	IdPessoa        uint `json:"idPessoa" binding:"required"`
	IdStatusBilhete uint `json:"idStatusBilhete" binding:"required"`
	IdBilhete       uint `json:"idBilhete"`
	Numero          uint `json:"numero" binding:"required"`
}

func saveBilheteToDB(bilhete BilheteController, c *gin.Context) (models.Bilhete, uint, error) {
	var bilheteDB models.Bilhete
	bilheteDB.IdBicho = bilhete.IdBicho
	bilheteDB.IdPessoa = bilhete.IdPessoa
	bilheteDB.IdStatusBilhete = bilhete.IdStatusBilhete

	result, idBilhete, erro := models.CriarBilhete(bilheteDB)
	return result, idBilhete, erro
}

func CriarBilhete(c *gin.Context) {
	var bilhete BilheteController
	err := c.BindJSON(&bilhete)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	//Get All numbers unavaillaibles
	errorNumbers := models.GetAllNumbersInsertedByIdBicho(bilhete.IdBicho)
	for _, element := range errorNumbers {
		if element.Numero == bilhete.Numero {
			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Número já selecionado"})
			return
		}
	}

	_, numeroBilhete, erro := saveBilheteToDB(bilhete, c)
	if erro != nil {
		fmt.Println(erro.Error())
		return
	}

	bilhete.IdBilhete = numeroBilhete
	var numeroBilheteSorte NumeroBilheteController
	numeroBilheteSorte.IdBilhete = bilhete.IdBilhete
	numeroBilheteSorte.Numero = bilhete.Numero
	saveNumeroBilheteToDB(numeroBilheteSorte, c)

	c.JSON(200, gin.H{
		"success": true,
		"data":    bilhete,
	})
}

func UpdateBilhete(c *gin.Context) {
	var bilhete BilheteController
	err := c.BindJSON(&bilhete)

	if err != nil {
		c.JSON(http.StatusNotAcceptable, gin.H{"error": err.Error()})
		return
	}

	if bilhete.IdBilhete == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Favor informar o id do bilhete"})
		return
	}

	var bilheteDB models.Bilhete
	bilheteDB.IdBicho = bilhete.IdBicho
	bilheteDB.IdPessoa = bilhete.IdPessoa
	bilheteDB.IdStatusBilhete = bilhete.IdStatusBilhete
	bilheteDB.IdBilhete = bilhete.IdBilhete

	result, err := models.UpdateBilheteStatus(bilheteDB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    result,
	})
}

func GetAllBilhetesFromAPerson(c *gin.Context) {
	idPessoa, erro := c.Request.URL.Query()["idPessoa"]

	if erro == false {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Favor informar o id da pessoa"})
		return
	}
	novoItem := strings.Join(idPessoa, " ")
	intVar, err := strconv.Atoi(novoItem)

	bilhetes, err := models.GetAllBilhetesFromAPerson(intVar)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"success": true,
		"data":    bilhetes,
	})
}
