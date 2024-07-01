package controllers

import (
	"fmt"
	"net/http"
)

func GetUser(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	name := queryParams["name"]
	surname := queryParams["surname"]
	passSer := queryParams["passSer"]
	passNum := queryParams["passNum"]
	patronymic := queryParams["patronymic"]
	address := queryParams["address"]

	fmt.Println(name, surname, passNum, passSer, patronymic, address)

	w.Write([]byte("Get User"))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create User"))
}
