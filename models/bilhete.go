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

type BilhetePerson struct {
	Numero    int    `db:"numero"`
	IdBilhete uint   `db:"idBilhete"`
	Concurso  int    `db:"concurso"`
	DataBicho string `db:"dataBicho"`
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

func GetAllBilhetesFromAPerson(idPessoa int) ([]BilhetePerson, error) {
	db := services.ConnectToDB()

	rows, error := db.Query("SELECT pessoaBilhete.numero, Bilhete.idBilhete, Bicho.Concurso, Bicho.DataBicho FROM Bilhete JOIN Bicho ON Bicho.idBicho = Bilhete.idBicho JOIN Pessoa ON Bilhete.idPessoa = Pessoa.idPessoa JOIN pessoaBilhete ON Bilhete.idBilhete = pessoaBilhete.idBilhete WHERE Pessoa.idPessoa = ?", idPessoa)

	if error != nil {
		println("deu erro")
	}

	var bilhetes []BilhetePerson

	for rows.Next() {
		var bilhete BilhetePerson
		rows.Scan(&bilhete.Numero, &bilhete.IdBilhete, &bilhete.Concurso, &bilhete.DataBicho)
		bilhetes = append(bilhetes, bilhete)
	}

	return bilhetes, error
}
