package main

import (
	"net/http"
	"html/template"
	"math/rand"
	"time"
	"strconv"
)

// message struct
type msg struct {
	Title string
	Guess int
	Message string
	HasWon bool
}

// func to generate a random number in a given range
func xrand(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

// serve index.html when app is run
func requestHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

// /guess handler 
func guessHandler(w http.ResponseWriter, r *http.Request) {
	// if cookie isnt name target set a new cookie called target and change its value to xrand(range)
	cookie, err := r.Cookie("target")
	if err != nil {	
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(xrand(1,20)), // set value to xrand(range)
		}

		http.SetCookie(w, cookie)
	}		

	// get the value from the input field in guess.tmpl
	guess, _ := strconv.Atoi(r.FormValue("guess"))

	// set title to string and set Guess value to guess, set HasWon boolean to false
	message_t := &msg{Title: "Guess a number between 1 and 20", Guess: guess, HasWon: false}

	// set target to cookie value(myrand)
	target, _ := strconv.Atoi(cookie.Value)

	// if the cookie is equal to the guess value set the cookies value to a new random number
	if target == guess {
		// correct
		cookie = &http.Cookie{
			Name: "target",
			Value: strconv.Itoa(xrand(1,20)), // set value to new xrand(range)
		}

		http.SetCookie(w, cookie)

		// set message to correct message string
		message_t.Message = "Your guess was Correct, Congratulations!"
		message_t.HasWon = true // set HasWon bool to true
	} else if guess < target {
		message_t.Message = "Your guess was too low, try again!"
	} else if guess > target {
		message_t.Message = "Your guess was too high, try again!"
	}
	
	// set template to parse guess.tmpl
	t, _ := template.ParseFiles("guess.tmpl")

	// execute template
	t.Execute(w, message_t)
}

func main() {
	http.HandleFunc("/", requestHandler)
	http.HandleFunc("/guess", guessHandler)
	http.ListenAndServe(":8080", nil)
}