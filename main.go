package main

import (
	"example/api-mysql/controller"
	"example/api-mysql/database"
	"example/api-mysql/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql" //Required for MySQL dialect
)

func main() {

	initDB()

	startServer()
}

func startServer() {
	log.Println("Starting the HTTP server on port 8090")
	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initDB() {
	config :=
		database.Config{
			ServerName: "localhost:3305",
			User:       "root",
			Password:   "password",
			DB:         "gomysql",
		}

	connectionString := database.GetConnectionString(config)
	err := database.Connect(connectionString)
	if err != nil {
		panic(err.Error())
	}
	database.Migrate(&entity.Person{})
}

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/create", controller.CreatePerson).Methods("POST")
	router.HandleFunc("/get", controller.GetAllPerson).Methods("GET")
	router.HandleFunc("/get/{id}", controller.GetPersonByID).Methods("GET")
	router.HandleFunc("/update/{id}", controller.UpdatePersonByID).Methods("PUT")
	router.HandleFunc("/delete/{id}", controller.DeletPersonByID).Methods("DELETE")
}
