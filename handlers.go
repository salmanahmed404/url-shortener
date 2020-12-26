package main

import (
	"fmt"
	"net/http"
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
	formData := r.PostForm

	fmt.Fprintln(w, formData)
}
