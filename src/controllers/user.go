package controllers

import "net/http"

func CreateUser(respose http.ResponseWriter, request *http.Request) {
	respose.Write([]byte("Criando usuário"))
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
