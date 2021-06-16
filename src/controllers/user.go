package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func CreateUser(respose http.ResponseWriter, request *http.Request) {
	body, erro := ioutil.ReadAll(request.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	user.CreatedAt = time.Now()

	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Erro(respose, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		responses.Erro(respose, http.StatusInternalServerError, erro)
		return
	}

	userRepository := repositories.NewUserRepo(db)
	result, erro := userRepository.Create(user)

	if erro != nil {
		responses.Erro(respose, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(respose, http.StatusCreated, result)
}

func GetUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}

func GetUsers(respose http.ResponseWriter, request *http.Request) {

	var users []models.User

	db, erro := database.Connect()

	if erro != nil {
		log.Fatal(erro)
	}

	userRepository := repositories.NewUserRepo(db)

	results, erro := userRepository.GetAll()

	if erro != nil {
		log.Fatal(erro)
	}

	for results.Next() {

		var id uint64
		var name string
		var nick string
		var email string
		var createdAt string

		err := results.Scan(&id, &name, &nick, &email, &createdAt)

		if err != nil {
			fmt.Println("Error on scan fields")
			log.Fatal(erro)
		}

		user := models.User{
			ID:        id,
			Name:      name,
			Nick:      nick,
			CreatedAt: time.Now(),
			Email:     email,
		}

		users = append(users, user)
	}

	defer results.Close()

	result, err := json.Marshal(users)

	if err != nil {
		log.Fatal(err)
	}

	respose.Write([]byte(result))
}

func UpdateUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}

func DeleteUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
}
