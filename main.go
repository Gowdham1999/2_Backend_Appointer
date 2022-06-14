package main

import (
	"2_Backend_Appointer/database"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/entity"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	fmt.Println("Hi!! Welcome to Go")

	router.HandleFunc("/customer", customer)
}

func init() {
	database.Migrate(&entity.Login{})
}

func customer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Success")
}
