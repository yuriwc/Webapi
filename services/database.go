package services

import (
	"log"

	"github.com/jmoiron/sqlx"
)

func ConnectToDB() *sqlx.Tx {
	db, err := sqlx.Connect(
		"mssql",
		"sqlserver://sa:TwoMonkeys4Yuri@127.0.0.1:1433?database=JogoDaSorte&connection+timeout=30",
	)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	tx := db.MustBegin()

	return tx
}
