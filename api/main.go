package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Initialize()
	//fmt.Println(config.StringConnetionPostgres)
	//fmt.Println(config.Port)

	fmt.Println("Api Radando...")
	r := router.Initialize()

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))
}
