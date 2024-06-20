package config

import (
	"api/src/util"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//String de conexao
	StringConnetionPostgres = ""
	//Porta de coonexao
	Port = 0
	//Driver de coonexao
	DriverPostgres = ""

	SecretKey []byte
)

// Carrega valores
func Initialize() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Port, erro = strconv.Atoi(os.Getenv("API_PORT"))
	if erro != nil {
		Port = 9000
	}

	DriverPostgres = os.Getenv("DB_DRIVER")

	SecretKey = []byte(os.Getenv("API_SECRET_KEY"))
	if len(SecretKey) == 0 {
		SecretKey = []byte(util.GeraKey())
	}
	StringConnetionPostgres = fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable search_path=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_SCHEMA"),
	)

	/*
	   StringConnetionMysql = fmt.Sprintf("%s:%s@%s?charset=utf8&parseTime=True&loc=Local",

	   	os.Getenv("DB_USER"),
	   	os.Getenv("DB_PASSWORD"),
	   	os.Getenv("DB_DATABASE"),

	   )
	*/
}
