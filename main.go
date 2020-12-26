package main

import (
	"log"
	"net/http"
)

func main() {

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
