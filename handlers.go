package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head> 
		<title> URL Shortener </title>
	</head>

	<body>
		<h2> URL Shortener </h2>
		<form action="http://127.0.0.1:8080/shorten/" method="POST">
			<label for="url_input"> Enter url </label>
			<input type="text" name="url_input">
			<br><br>
			<input type="submit" value="Submit">
		</form>
	</body>

	</html>`

	w.Write([]byte(str))
}

func shortenHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	inputURL := r.PostForm.Get("url_input")
	db, err := gorm.Open("sqlite3", "gorl.db")
	defer db.Close()
	if err != nil {
		w.Write([]byte("DB Init Error"))
		log.Fatal(err)
	}
	url := URL{LongURL: inputURL}
	result := db.Create(&url)
	if result.Error != nil {
		w.Write([]byte("DB Create Error"))
		log.Fatal(result.Error)
	}
	shortURL := base62Encode(url.ID)
	w.Write([]byte(shortURL))
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	shortURL := mux.Vars(r)["shortURL"]
	db, err := gorm.Open("sqlite3", "gorl.db")
	defer db.Close()
	if err != nil {
		w.Write([]byte("DB Init Error"))
		log.Fatal(err)
	}
	id := base62Decode(shortURL)
	var url URL
	db.First(&url, id)
	redirectURL := "http://" + url.LongURL
	http.Redirect(w, r, redirectURL, 301)
}
