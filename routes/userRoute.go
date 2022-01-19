package routes

import (
	userController "github.com/MasoudHeydari/golang-crud-rest-api/controllers"
	"github.com/gorilla/mux"
)

func RegisterUserRoutes(router *mux.Router, userController *userController.UserController) {
	router.HandleFunc("/", userController.Home).Methods("GET")
	router.HandleFunc("/user/", userController.CreateNewUser).Methods("POST")
}