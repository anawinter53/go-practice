package main

import (
	"fmt"
	"net/http"
)

func server() {
	// setting handlers
	http.HandleFunc("/", handler)
	http.HandleFunc("/Hello", handler2)
	http.ListenAndServe(":8080", nil)


}

// creating handlers
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to my homepage\n")
}

// w is the response
func handler2(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}
