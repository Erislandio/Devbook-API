package config

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	Path                string = ""
	Port                int    = 5000
	SqlConnectionString        = ""
)

func Config() {

	var erro error

	if erro = godotenv.Load(); erro != nil {
		log.Fatal(erro)
	}

	Path = os.Getenv("PATH_SQLITE")

	Port, erro = strconv.Atoi(os.Getenv("PORT"))

	if erro != nil {
		message := fmt.Sprintf("Port not have value, apply default value: %d", Port)
		fmt.Println(message)
	}

	SqlConnectionString = fmt.Sprintf(
		"%s:%s@/%s?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

}
