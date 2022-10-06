package models

import "webApi/services"

type Bicho struct {
	IdBicho       uint   `db:"idBicho"`
	Concurso      int    `db:"concurso"`
	DataBicho     string `db:"dataBicho"`
	IdPessoaBicho uint   `db:"idPessoaBicho"`
	IdStatusBicho uint   `db:"idStatusBicho"`
}

func CriarBicho(bicho Bicho) (Bicho, error) {
	db := services.ConnectToDB()

	error := db.QueryRow("INSERT INTO Bicho OUTPUT INSERTED.idBicho VALUES (?, ?, ?, 1);", bicho.DataBicho, bicho.Concurso, bicho.IdPessoaBicho).Scan(&bicho.IdBicho)

	if error != nil {
		println("deu erro")
	}
	db.Commit()
	return bicho, error
}

func GetAllBichos() ([]Bicho, error) {
	db := services.ConnectToDB()

	rows, error := db.Query("SELECT * FROM Bicho")

	if error != nil {
		println("deu erro")
	}

	var bichos []Bicho

	for rows.Next() {
		var bicho Bicho
		rows.Scan(&bicho.IdBicho, &bicho.DataBicho, &bicho.Concurso, &bicho.IdPessoaBicho, &bicho.IdStatusBicho)
		bichos = append(bichos, bicho)
	}

	return bichos, error
}
