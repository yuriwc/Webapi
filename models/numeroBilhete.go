package models

import (
	"fmt"
	"webApi/services"
)

type NumeroBilhete struct {
	IdNumeroBilhete uint `db:"idPessoaBilhete"`
	IdBilhete       uint `db:"idBilhete"`
	Numero          uint `db:"numero"`
}

func CriarNumeroBilhetes(numeroBilhete NumeroBilhete) (NumeroBilhete, error) {
	db := services.ConnectToDB()

	error := db.QueryRow("INSERT INTO pessoaBilhete (numero, idBilhete) OUTPUT INSERTED.idPessoaBilhete VALUES (?, ?);", numeroBilhete.Numero, numeroBilhete.IdBilhete).Scan(&numeroBilhete.IdNumeroBilhete)

	if error != nil {
		println("deu erro")
	}

	db.Commit()
	return numeroBilhete, error
}

func GetAllNumbersInsertedByIdBicho(idBicho uint) []NumeroBilhete {
	db := services.ConnectToDB()
	var numeroBilhete []NumeroBilhete
	db.Select(&numeroBilhete, "SELECT pessoaBilhete.numero FROM pessoaBilhete JOIN Bilhete ON pessoaBilhete.idBilhete = Bilhete.idBilhete JOIN Bicho ON Bicho.idBicho = Bilhete.idBicho WHERE Bicho.idBicho = ?", idBicho)

	fmt.Println(numeroBilhete)
	return numeroBilhete
}
