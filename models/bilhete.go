package models

import (
	"fmt"
	"webApi/services"
)

type Bilhete struct {
	IdBilhete       uint `db:"idBilhete"`
	IdBicho         uint `db:"idBicho"`
	IdPessoa        uint `db:"idPessoa"`
	IdStatusBilhete uint `db:"idStatusBilhete"`
}

func CriarBilhete(bilhete Bilhete) (Bilhete, uint, error) {
	db := services.ConnectToDB()

	error := db.QueryRow("INSERT INTO Bilhete (idBicho, idPessoa, idStatusBilhete) OUTPUT INSERTED.idBilhete VALUES (?, ?, ?);", bilhete.IdBicho, bilhete.IdPessoa, bilhete.IdStatusBilhete).Scan(&bilhete.IdBilhete)

	fmt.Println(bilhete.IdBilhete)

	if error != nil {
		println("deu erro")
	}

	db.Commit()
	return bilhete, bilhete.IdBilhete, error
}

func UpdateBilheteStatus(bilhete Bilhete) (Bilhete, error) {
	db := services.ConnectToDB()

	_, err := db.Exec("UPDATE Bilhete SET idStatusBilhete = ? WHERE idBilhete = ?;", bilhete.IdStatusBilhete, bilhete.IdBilhete)

	db.Commit()
	return bilhete, err
}
