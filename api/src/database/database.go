package database

import (
	"api/src/config"
	"database/sql"
	"fmt"

	//Driver
	_ "github.com/lib/pq"
)

func InitializePostgres() (*sql.DB, error) {
	//fmt.Println(config.StringConnetionPostgres)
	db, erro := sql.Open(config.DriverPostgres, config.StringConnetionPostgres)
	if erro != nil {
		return nil, erro
	}

	if erro := db.Ping(); erro != nil {
		db.Close()
		fmt.Println("Initializing db postgres test", erro.Error())
		return nil, erro
	}

	return db, nil
}
