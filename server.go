package main

import (
	"fmt"
	"log"
	"net/http"
)

type user struct {
	name      string
	age       int
	gender    string
	isMarried bool
}

var users []user

func main() {
	users = []user{
		{name: "chinedu", age: 24, gender: "male", isMarried: true},
		{name: "chinedu", age: 24, gender: "male", isMarried: true},
	}
	fmt.Println(users)
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go programing api")
	fmt.Println("home endpoint hit")

}
func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":5001", nil))
}
