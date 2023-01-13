package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
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

	if erro = user.Prepare("cadastro"); erro != nil {
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
	parametros := mux.Vars(request)

	userID, erro := strconv.ParseUint(parametros["usuarioId"], 10, 64)
	if erro != nil {
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

	user, erro := repo.SearchByID(userID)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusOK, user)
}

func UpdateUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := auth.ExtractUserId(request)
	if erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	if userIDToken != userID {
		responses.Erro(response, http.StatusForbidden, erro)
		return
	}

	resquestBody, erro := ioutil.ReadAll(request.Body)
	if erro != nil {
		responses.Erro(response, http.StatusUnprocessableEntity, erro)
		return
	}

	var user models.User
	if erro = json.Unmarshal(resquestBody, &user); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	if erro = user.Prepare("edicao"); erro != nil {
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
	if erro = repo.Update(userID, user); erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusNoContent, nil)
}

func DeleteUser(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	userIDToken, erro := auth.ExtractUserId(request)
	if erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	if userIDToken != userID {
		responses.Erro(response, http.StatusForbidden, erro)
		return
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if erro = repo.Delete(userID); erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusNoContent, nil)

}
