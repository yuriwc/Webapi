package models

import (
	"database/sql"
	"webApi/services"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type Pessoa struct {
	Nome       string         `db:"Nome"`
	Telefone   string         `db:"Telefone"`
	CPF        string         `db:"CPF"`
	idEndereco sql.NullString `db:"idEndereco"`
}

func CreatePessoa(pessoa Pessoa) {
	tx := services.ConnectToDB()

	tx.NamedExec("INSERT INTO Pessoa (Nome, Telefone, CPF) VALUES (:Nome, :Telefone, :CPF)", &Pessoa{Nome: pessoa.Nome, Telefone: pessoa.Telefone, CPF: pessoa.CPF})
	tx.Commit()
}

func GetPessoaByNumber(number string) Pessoa {
	db := services.ConnectToDB()
	pessoa := Pessoa{}

	err := db.QueryRow("SELECT Nome, Telefone, CPF, idEndereco FROM Pessoa WHERE Telefone = ?", number).Scan(&pessoa.Nome, &pessoa.Telefone, &pessoa.CPF, &pessoa.idEndereco)

	if err != nil {
		return Pessoa{}
	}
	return pessoa
}

func UpdateEnderecoPessoa(idPessoa int, idEndereco uint) {
	db := services.ConnectToDB()
	_, err := db.Exec("UPDATE Pessoa SET idEndereco = ? WHERE idPessoa = ?", idEndereco, idPessoa)
	db.Commit()

	if err != nil {
		return
	}
}
