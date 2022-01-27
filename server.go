package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	name      string
	age       int
	gender    string
	isMarried bool
}

var Users []User

func main() {
	Users = []User{
		{name: "chinedu", age: 24, gender: "male", isMarried: true},
		{name: "chinedu", age: 24, gender: "male", isMarried: true},
	}
	fmt.Println(Users)
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go programing api")
	fmt.Println("home endpoint hit")

}
func getAllUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all users endpoint hit")
	json.NewEncoder(w).Encode(Users)
}
func handleRequest() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/users", getAllUsers)
	log.Fatal(http.ListenAndServe(":5001", nil))
}
