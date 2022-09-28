package models

import "webApi/services"

type Bicho struct {
	IdBicho       uint   `db:"idBicho"`
	Concurso      int    `db:"concurso"`
	DataBicho     string `db:"dataBicho"`
	IdPessoaBicho uint   `db:"idPessoaBicho"`
}

func CriarBicho(bicho Bicho) (Bicho, error) {
	db := services.ConnectToDB()

	error := db.QueryRow("INSERT INTO Bicho OUTPUT INSERTED.idBicho VALUES (?, ?, ?);", bicho.DataBicho, bicho.Concurso, bicho.IdPessoaBicho).Scan(&bicho.IdBicho)

	if error != nil {
		println("deu erro")
	}
	db.Commit()
	return bicho, error
}
