package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func CreateUser(respose http.ResponseWriter, request *http.Request) {
	body, erro := ioutil.ReadAll(request.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User

	if erro = json.Unmarshal(body, &user); erro != nil {
		log.Fatal(erro)
	}

	db, erro := database.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

	userRepository := repositories.NewUserRepo(db)
	userRepository.Create(user)
}

func GetUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}

func GetUsers(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}

func UpdateUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}

func DeleteUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}
