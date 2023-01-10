package controllers

import (
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
)

func CreateUser(response http.ResponseWriter, request *http.Request) {
	requestBody, erro := ioutil.ReadAll(request.Body)
	if erro != nil {
		responses.Erro(response, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(requestBody, &user); erro != nil {
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
	defer db.Close()

	repository := repository.NewUserRepository(db)
	user.ID, erro = repository.Create(user)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusCreated, user)
}

func GetUsers(response http.ResponseWriter, request *http.Request) {
	nameOrNick := strings.ToLower(request.URL.Query().Get("user"))

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repository := repository.NewUserRepository(db)
	users, erro := repository.Search(nameOrNick)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusOK, users)
}

func GetUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Buscando um usuario"))
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Atualizando um usuario"))
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Deletando um usuario"))
}
