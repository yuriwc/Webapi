package models

import (
	"database/sql"
	"fmt"
	"webApi/services"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type Pessoa struct {
	IdPessoa   int            `db:"idPessoa"`
	Nome       string         `db:"Nome"`
	Telefone   string         `db:"Telefone"`
	CPF        string         `db:"CPF"`
	idEndereco sql.NullString `db:"idEndereco"`
}

type Ganhador struct {
	Nome      string `db:"Nome"`
	Concurso  int    `db:"Concurso"`
	numero    int    `db:"numero"`
	Descricao string `db:"Descricao"`
}

func CreatePessoa(pessoa Pessoa) (int, error) {
	tx := services.ConnectToDB()

	err := tx.QueryRow("INSERT INTO Pessoa (Nome, Telefone, CPF) OUTPUT INSERTED.idPessoa VALUES (?,?,?)", pessoa.Nome, pessoa.Telefone, pessoa.CPF).Scan(&pessoa.IdPessoa)
	tx.Commit()

	if err != nil {
		return 0, err
	}

	return pessoa.IdPessoa, nil
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

func GetGanhandor(numeroBilhete int, concurso int) (Ganhador, error) {
	db := services.ConnectToDB()
	ganhador := Ganhador{}

	err := db.QueryRow("SELECT Pessoa.Nome, Bicho.Concurso, pessoaBilhete.numero, StatusBilhete.Descricao FROM Bicho JOIN Bilhete ON Bicho.idBicho = Bilhete.idBicho JOIN pessoaBilhete ON pessoaBilhete.idBilhete = Bilhete.idBilhete JOIN Pessoa ON Pessoa.idPessoa = Bicho.idPessoaBicho JOIN StatusBilhete ON Bilhete.idStatusBilhete = StatusBilhete.idStatus WHERE pessoaBilhete.numero = ? AND Concurso = ?", numeroBilhete, concurso).Scan(&ganhador.Nome, &ganhador.Concurso, &ganhador.numero, &ganhador.Descricao)

	if err != nil {
		fmt.Println(err.Error())
		if err.Error() == "sql: no rows in result set" {
			return Ganhador{}, nil
		}
		return Ganhador{}, err
	}

	return ganhador, err
}
