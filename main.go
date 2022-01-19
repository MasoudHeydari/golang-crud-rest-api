package main

import (
	"fmt"
	userController "github.com/MasoudHeydari/golang-crud-rest-api/controllers"
	"github.com/MasoudHeydari/golang-crud-rest-api/database"
	"github.com/MasoudHeydari/golang-crud-rest-api/routes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("A Simple REST API Project with MySQL Database")

	dbConn, err := database.ConnectToDB()
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
	router := mux.NewRouter()
	userHandler := userController.NewUserHandler(dbConn)
	routes.RegisterUserRoutes(router, userHandler)
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
