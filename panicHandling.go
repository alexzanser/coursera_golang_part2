package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!")
}

func doPanic(w http.ResponseWriter, r *http.Request) {
	panic("oops")
}


func main() {
	http.HandleFunc("/hello/", hello)
	http.HandleFunc("/panic/", doPanic)
	http.ListenAndServe(":8080", nil)
}