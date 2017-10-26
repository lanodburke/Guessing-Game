package main

import (
	"net/http"
	"html/template"
)

type msg struct {
	Message string
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {

	message := &msg{Message: "Guess a number between 1 and 20"}
	
	s, _ := template.ParseFiles("guess.tmpl")

	s.Execute(w, message)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}