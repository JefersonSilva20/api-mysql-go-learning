package controller

import (
	"encoding/json"
	"example/api-mysql/database"
	"example/api-mysql/entity"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllPerson get all person data
func GetAllPerson(response http.ResponseWriter, request *http.Request) {
	log.Println("Listando todos as pessoas")
	var persons []entity.Person
	database.Connector.Find(&persons)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(persons)
}

//GetPersonByID returns person with specific ID
func GetPersonByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	log.Println("Recuperando pessoa com id", key)

	var person entity.Person
	database.Connector.First(&person, key)
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(person)
}

//CreatePerson creates person
func CreatePerson(response http.ResponseWriter, request *http.Request) {
	requestBody, _ := ioutil.ReadAll(request.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusCreated)
	json.NewEncoder(response).Encode(person)
}

//UpdatePersonByID updates person with respective ID
func UpdatePersonByID(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	log.Println("Atualizando dados da pessoa com id", key)
	requestBody, _ := ioutil.ReadAll(request.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	json.NewEncoder(response).Encode(person)
}

//DeletPersonByID delete's person with specific ID
func DeletPersonByID(response http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	log.Println("Deletando dados da pessoa com id", key)
	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	response.WriteHeader(http.StatusNoContent)
}
