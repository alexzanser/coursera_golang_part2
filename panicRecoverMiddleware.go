package main

import (
	"fmt"
	"log"
	"net/http"
)

func greeting(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlind greeting task at %s", r.URL.Path)
	fmt.Fprintf(w, "Hello, World!")
}

func doPanic(w http.ResponseWriter, r *http.Request) {
	log.Printf("handlind panic at %s", r.URL.Path)
	panic("some kind of error")
}

func panicRecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("error recovered: %s", err)
				http.Error(w, "Internal server error", http.StatusInternalServerError)
			}
		}() 
		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello/", greeting)
	mux.HandleFunc("/panic/", doPanic)
	pr := panicRecoverMiddleware(mux)
	http.ListenAndServe(":8080", pr)
}