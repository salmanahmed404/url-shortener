package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

func main() {

	//Database initialization
	db, err := gorm.Open("sqlite3", "gorl.db")
	defer db.Close()

	if err != nil {
		log.Fatal("DB Connection Error: ", err)
	}

	errors := db.AutoMigrate(&URL{}).Error
	if errors != nil {
		log.Fatal(errors.Error())
	}

	//HttpServer
	addr := "127.0.0.1:8080"
	r := mux.NewRouter()

	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/shorten/", shortenHandler).Methods("POST")
	r.HandleFunc("/gorl/{shortURL}", redirectHandler).Methods("GET")

	log.Println("Listening on: ", addr)
	log.Fatal(http.ListenAndServe(addr, r))
}
