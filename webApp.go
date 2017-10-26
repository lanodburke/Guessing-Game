package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
)

type msg struct {
	Message string
}

func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func requestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func guessHandler(w http.ResponseWriter, r *http.Request) {
	myrand := xrand(1, 20)

	mycookie, err := r.Cookie("target")
	if(err == nil) {
		target, _ := strconv.Atoi(mycookie.Value)

		if target == 0 {
			target = myrand
		}
	}

	http.SetCookie(w, mycookie)

	message := &msg{Message: "Guess a number between 1 and 20"}
	
	s, _ := template.ParseFiles("guess.tmpl")

	s.Execute(w, message)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}