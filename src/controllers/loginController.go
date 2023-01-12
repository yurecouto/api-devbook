package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Login(response http.ResponseWriter, request *http.Request) {
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

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	userInDatabase, erro := repo.FindUserByEmail(user.Email)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(user.Password, userInDatabase.Password); erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	token, erro := auth.GenerateToken(userInDatabase.ID)
	if erro != nil {

		fmt.Println(erro)
	}
	response.Write([]byte(token))
}
