package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/MasoudHeydari/golang-crud-rest-api/database"
	"github.com/MasoudHeydari/golang-crud-rest-api/models"
	"github.com/MasoudHeydari/golang-crud-rest-api/repository/user"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type UserController struct {
	repo user.UserRepository
}

func NewUserHandler(db *database.DB) *UserController {
	return &UserController{
		repo: user.NewSQLUserRepository(db.Sql),
	}
}

func (usrController *UserController) Home(w http.ResponseWriter, r *http.Request) {
	msgStr := "welcome to home page, this is a simple Golang REST API CRUD app"
	respondWithJSON(w, http.StatusOK, map[string]string{"massage": msgStr})
}

func (usrController *UserController) CreateNewUser(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{}
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondWithError(w, 400, "error occurred while decoding new user!")
		return
	}

	fmt.Println("new request")
	fmt.Println(newUser)
	newUserId := usrController.repo.CreateNewUser(&newUser)
	newUserStr := fmt.Sprintf("%d", newUserId)

	respondWithJSON(w, http.StatusCreated, map[string]string{
		"massage": "new user created successfully",
		"user_id": newUserStr,
	})
}

func (usrController *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	allUsers := usrController.repo.GetAllUsers()
	respondWithJSON(w, http.StatusOK, allUsers)
}

func (usrController UserController) Update(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "id is not valid")
		return
	}
	userToUpdate := models.User{}
	userToUpdate.ID = uint(userId)
	json.NewDecoder(r.Body).Decode(&userToUpdate)
	userToUpdate = *usrController.repo.Update(userToUpdate.ID, &userToUpdate)
	respondWithJSON(w, http.StatusOK, userToUpdate)
}

func (usrController UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userId, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "id is not valid")
	}

	usrController.repo.Delete(uint(userId))
	respondWithJSON(w, http.StatusOK, map[string]string{"massage": "user deleted successfully"})
}

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, statusCode int, msg string) {
	respondWithJSON(w, statusCode, map[string]string{"massage": msg})
}
