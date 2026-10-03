package main

import (
	"fmt"
	"net/http"
)

var accessCounter int = 0

func messageResponse(w http.ResponseWriter, r *http.Request){

	fmt.Println(r.URL)
	accessCounter++
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"message": "hello", "amount": %d}`, accessCounter)
	fmt.Printf("Accessed %d times.\n", accessCounter)
}

func main() {
	http.HandleFunc("/", messageResponse)

	http.ListenAndServe(":8080", nil)
}