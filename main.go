package main

import (
	"log"
	"net/http"

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
	s := http.Server{
		Addr: addr,
	}

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/shorten/", shortenHandler)

	log.Println("Listening on: ", addr)
	log.Fatal(s.ListenAndServe())
}
