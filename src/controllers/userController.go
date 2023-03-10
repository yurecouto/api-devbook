package controllers

import (
	"api/src/auth"
	"api/src/database"
	"api/src/models"
	"api/src/repository"
	"api/src/responses"
	"api/src/security"
	"encoding/json"
	"errors"
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

func FollowUser(response http.ResponseWriter, request *http.Request) {
	followerID, erro := auth.ExtractUserId(request)
	if erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	if userID == followerID {
		responses.Erro(response, http.StatusForbidden, errors.New("Cant follow yourself"))
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if erro = repo.Follow(userID, followerID); erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusNoContent, nil)
}

func UnFollowUser(response http.ResponseWriter, request *http.Request) {
	followerID, erro := auth.ExtractUserId(request)
	if erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	if userID == followerID {
		responses.Erro(response, http.StatusForbidden, errors.New("Cant follow yourself"))
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	if erro = repo.UnFollow(userID, followerID); erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusNoContent, nil)
}

func FindFollowers(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	followers, erro := repo.FindFollowers(userID)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusOK, followers)
}

func FindFollowing(response http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)
	followers, erro := repo.FindFollowing(userID)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusOK, followers)
}

func UpdatePassword(response http.ResponseWriter, request *http.Request) {
	userIDToken, erro := auth.ExtractUserId(request)
	if erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	params := mux.Vars(request)
	userID, erro := strconv.ParseUint(params["usuarioId"], 10, 64)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	if userID != userIDToken {
		responses.Erro(response, http.StatusUnauthorized, errors.New("Nao se pode atualizar outro usuario"))
	}

	requestBody, erro := ioutil.ReadAll(request.Body)

	var password models.Password
	if erro = json.Unmarshal(requestBody, &password); erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
	}

	db, erro := database.Connect()
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}
	defer db.Close()

	repo := repository.NewUserRepository(db)

	passwordDB, erro := repo.FindPasswordById(userID)
	if erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	if erro = security.VerifyPassword(passwordDB, password.Old); erro != nil {
		responses.Erro(response, http.StatusUnauthorized, erro)
		return
	}

	passwordHash, erro := security.Hash(password.New)
	if erro != nil {
		responses.Erro(response, http.StatusBadRequest, erro)
		return
	}

	if erro = repo.UpdatePassword(userID, string(passwordHash)); erro != nil {
		responses.Erro(response, http.StatusInternalServerError, erro)
		return
	}

	responses.JSON(response, http.StatusNoContent, nil)
}
