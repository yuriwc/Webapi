package models

import (
	"webApi/services"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type User struct {
	IdUser        int    `db:"idUser"`
	IdPessoa      int    `db:"idPessoa"`
	IdNivel       int    `db:"idNivel"`
	NumeroCelular string `db:"NumeroCelular"`
	Senha         string `db:"Senha"`
}

type Senha struct {
	Senha  string `db:"Senha"`
	IdUser uint   `db:"IdUser"`
}

func CreateUser(user User) (int, error) {
	tx := services.ConnectToDB()
	var userNew User

	if user.IdPessoa != 0 {
		error := tx.QueryRow(
			"INSERT INTO UserLogin (idNivel, idPessoa, NumeroCelular, Senha) OUTPUT INSERTED.idUser VALUES (:idNivel, :idPessoa, :NumeroCelular, :Senha)", &User{IdNivel: 3, IdPessoa: user.IdPessoa, NumeroCelular: user.NumeroCelular, Senha: user.Senha},
		).Scan(&userNew.IdUser)

		tx.Commit()

		if error != nil {
			return 0, error
		}

		return userNew.IdUser, nil
	} else {
		err := tx.QueryRow(
			"INSERT INTO UserLogin (idNivel, NumeroCelular, Senha) OUTPUT INSERTED.idUser VALUES (?,?,?);", 3, user.NumeroCelular, user.Senha).Scan(&userNew.IdUser)
		tx.Commit()

		if err != nil {
			return 0, err
		}

		return userNew.IdUser, nil
	}
}

func GetUser(login string, password string) (User, error) {
	db := services.ConnectToDB()
	var user User

	err := db.QueryRow("SELECT idNivel, NumeroCelular FROM UserLogin WHERE NumeroCelular = ? AND Senha = ?", login, password).Scan(&user.IdNivel, &user.NumeroCelular)

	if err != nil {
		return User{}, err
	}
	return user, err
}

func GetUserByUsername(login string) (Senha, error) {
	db := services.ConnectToDB()
	var senha Senha

	err := db.QueryRow("SELECT Senha, idUser FROM UserLogin WHERE NumeroCelular = ?", login).Scan(&senha.Senha, &senha.IdUser)

	if err != nil {
		return Senha{}, err
	}
	return senha, err
}

func UpdateIdPessoaFromUser(idPessoa int, idUser int) error {
	db := services.ConnectToDB()
	_, err := db.Exec("UPDATE UserLogin SET idPessoa = ? WHERE idUser = ?", idPessoa, idUser)
	db.Commit()

	if err != nil {
		return err
	}
	return nil
}
