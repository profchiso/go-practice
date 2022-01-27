package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	handleRequest()
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to go programing api")
	fmt.Println("home endpont hit")

}
func handleRequest() {
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":6000", nil))
}
