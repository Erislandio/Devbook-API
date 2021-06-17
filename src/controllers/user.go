package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repositories"
	"api/src/responses"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	body, erro := ioutil.ReadAll(request.Body)

	if erro != nil {
		log.Fatal(erro)
	}

	var user models.User
	user.CreatedAt = time.Now()

	if erro = json.Unmarshal(body, &user); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare(); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	userRepository := repositories.NewUserRepo(db)
	result, erro := userRepository.Create(user)

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusCreated, result)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	nickname := strings.ToLower(request.URL.Query().Get("nickname"))

	db, erro := database.Connect()

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repository := repositories.NewUserRepo(db)

	results, erro := repository.GetByNickName(nickname)

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	users := ProcessResults(results)
	responses.JSON(response, http.StatusCreated, users)
}

func ProcessResults(results *sql.Rows) []models.User {
	var users []models.User

	for results.Next() {

		var id uint64
		var name string
		var nick string
		var email string
		var createdAt string

		err := results.Scan(&id, &name, &nick, &email, &createdAt)

		if err != nil {
			fmt.Println("Error on scan fields")
			log.Fatal(err)
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

	return users
}

func GetUserById(response http.ResponseWriter, request *http.Request) {
	userid := strings.ToLower(request.URL.Query().Get("id"))

	db, erro := database.Connect()

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
	}

	defer db.Close()

	repository := repositories.NewUserRepo(db)

	results, erro := repository.GetByID(userid)

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	users := ProcessResults(results)

	if len(users) == 0 {
		responses.Erro(response, http.StatusBadRequest, errors.New("User nor found"))
		return
	}

	responses.JSON(response, http.StatusCreated, users)

}

func GetUsers(response http.ResponseWriter, request *http.Request) {

	nickname := strings.ToLower(request.URL.Query().Get("nickname"))
	userid := strings.ToLower(request.URL.Query().Get("id"))

	if nickname != "" {
		GetUser(response, request)
		return
	}

	if userid != "" {
		GetUserById(response, request)
		return
	}

	db, erro := database.Connect()

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	userRepository := repositories.NewUserRepo(db)

	results, erro := userRepository.GetAll()

	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	users := ProcessResults(results)

	responses.JSON(response, http.StatusOK, users)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Criando usuário"))
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Criando usuário"))
}
