package main

import (
	"fmt"
	"net/http"
)
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Awesome Go App1")
}

func setupRoutes() {
	http.HandleFunc("/", homePage)
}


func main() {
fmt.Println("Hello Welcome to Golng World!")
	setupRoutes()
	http.ListenAndServe(":8000", nil)
}
