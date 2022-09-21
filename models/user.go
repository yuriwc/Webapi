package models

import (
	"database/sql"
	"webApi/services"

	_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/lib/pq"
)

type User struct {
	IdPessoa int`db:"idPessoa"`
	IdNivel int `db:"idNivel"`
	NumeroCelular string `db:"NumeroCelular"`
	Senha string `db:"Senha"`
}

type Senha struct {
	Senha string `db:"Senha"`
}

func CreateUser(user User) (sql.Result, error){
	tx := services.ConnectToDB()

	if (user.IdPessoa != 0) {
		result, err := tx.NamedExec(
			"INSERT INTO UserLogin (idNivel, idPessoa, NumeroCelular, Senha) VALUES (:idNivel, :idPessoa, :NumeroCelular, :Senha)", &User{IdNivel: 3, IdPessoa: user.IdPessoa, NumeroCelular: user.NumeroCelular, Senha: user.Senha},
		)
		tx.Commit()
		return result, err
	} else {
		result, err := tx.NamedExec(
			"INSERT INTO UserLogin (idNivel, NumeroCelular, Senha) VALUES (:idNivel, :NumeroCelular, :Senha)", &User{IdNivel: 3, NumeroCelular: user.NumeroCelular, Senha: user.Senha},
		)
		tx.Commit()
		return result, err
	}
}

func GetUser(login string, password string) (User, error){
	db := services.ConnectToDB()
	var user User

	err := db.QueryRow("SELECT idNivel, NumeroCelular FROM UserLogin WHERE NumeroCelular = ? AND Senha = ?", login, password).Scan(&user.IdNivel, &user.NumeroCelular)

	if err!= nil {
		return User{}, err
	}
	return user, err
}

func GetUserByUsername(login string) (Senha, error){
	db := services.ConnectToDB()
	var senha Senha

	err := db.QueryRow("SELECT Senha FROM UserLogin WHERE NumeroCelular = ?", login).Scan(&senha.Senha)

	if err!= nil {
		return Senha{}, err
	}
	return senha, err
}