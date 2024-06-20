package main

import (
	"fmt"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/cookie"
	"webapp/src/router"
	"webapp/src/utils"
)

func main() {
	config.Initialize()
	cookie.Inicializa()
	utils.CarregarTemplate()
	r := router.Init()
	fmt.Println("App rodando na porta...", config.Porta)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Porta), r))
}
