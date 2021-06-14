package main

import (
	"fmt"
	"log"
	"net/http"

	"api/src/config"
	"api/src/router"
)

func main() {
	fmt.Println("Carregando envs")
	config.Config()
	fmt.Println(config.Path)
	fmt.Println("inciando: API - DevBook -> http:localhost:5000")
	router := router.Init()

	port := fmt.Sprintf(":%d", config.Port)
	log.Fatal(http.ListenAndServe(port, router))

}
