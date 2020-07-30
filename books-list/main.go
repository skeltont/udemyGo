package main

import (
	// "fmt"
	// "reflect"
	// "encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"

	"books-list/controllers"
	"books-list/driver"
	"books-list/models"
	"database/sql"
	"github.com/subosito/gotenv"
)

func init() {
	gotenv.Load()
}

var books []models.Book
var db *sql.DB

func logFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	db = driver.ConnectDB()
	controller := controllers.Controllers{}
	router := mux.NewRouter()

	router.HandleFunc("/books", controller.GetBooks(db)).Methods("GET")
	router.HandleFunc("/books/{id}", controller.GetBook(db)).Methods("GET")
	// router.HandleFunc("/books", controller.AddBook).Methods("POST")
	// router.HandleFunc("/books", controller.UpdateBook).Methods("PUT")
	// router.HandleFunc("/books/{id}", controller.RemoveBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
