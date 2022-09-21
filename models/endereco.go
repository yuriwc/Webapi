package models

import (
	"webApi/services"
)

type Endereco struct {
	IdEndereco uint `db:"idEndereco"`
	RUA string `db:"RUA"`
	Bairro string `db:"Bairro"`
	Cidade string `db:"Cidade"`
	Estado string `db:"Estado"`
}

func CriarEndereco(endereco Endereco) (Endereco, error){
	db := services.ConnectToDB()

	error := db.QueryRow("INSERT INTO Endereco (RUA, Bairro, Cidade, Estado) OUTPUT INSERTED.idEndereco VALUES (?, ?, ?, ?);", endereco.RUA, endereco.Bairro, endereco.Cidade, endereco.Estado).Scan(&endereco.IdEndereco)

	if error != nil {
		println("deu erro")
	}
	db.Commit()
	
	return endereco, error
}