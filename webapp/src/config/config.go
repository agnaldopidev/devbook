package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	//Url da api
	APIURL = ""
	//porta da api
	Porta = 0
	//utilizado para autenticar o cookie
	Hashkey []byte
	//utilizado para criptografar o cookie
	Blockkey []byte
)

// Inicializa as variaveis de ambiente
func Initialize() {
	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}
	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}

	APIURL = os.Getenv("API_URL")
	Hashkey = []byte(os.Getenv("HASH_KEY"))
	Blockkey = []byte(os.Getenv("BLOCK_KEY"))

}

/*
func geraKey() string {
	hashkeyGerado := hex.EncodeToString(securecookie.GenerateRandomKey(16))
	fmt.Println(hashkeyGerado)
	return hashkeyGerado
}
*/
