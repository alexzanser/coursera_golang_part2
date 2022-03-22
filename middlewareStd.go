package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func Greeting(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome! Thanks for visiting!\n")
}	

func LoggindMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %s\n", r.Method, r.RequestURI, time.Since(start))
	})
}


func main () {
	handler := http.TimeoutHandler(http.HandlerFunc(Greeting), 15 * time.Microsecond, "timed out")	
	lm := LoggindMiddleware(handler)
	log.Fatal(http.ListenAndServe("localhost:8080", lm))
}