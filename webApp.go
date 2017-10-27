package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
)

type msg struct {
	Title string
	Guess int
	Message string
	HasWon bool
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

	cookie, err := r.Cookie("target")
	if err != nil {	
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(myrand),
			Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w, cookie)
	}		

	guess, _ := strconv.Atoi(r.FormValue("guess"))

	message_t := &msg{Title: "Guess a number between 1 and 20", Guess: guess, HasWon: false}

	target, _ := strconv.Atoi(cookie.Value)

	if target == guess {
		// correct
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(myrand),
			Expires: time.Now().Add(72 * time.Hour),
		}

		http.SetCookie(w, cookie)

		message_t.Message = "Your guess was Correct, Congratulations!"
		message_t.HasWon = true
	} else if guess < target {
		message_t.Message = "Your guess was too low, try again!"
	} else if guess > target {
		message_t.Message = "Your guess was too high, try again!"
	}
	
	s, _ := template.ParseFiles("guess.tmpl")

	s.Execute(w, message_t)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}