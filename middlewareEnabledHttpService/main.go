package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

// middleware filter incoming HTTP requests.
// if the request pass the filter, it calls the next HTTP handler.
type middleware func(http.HandlerFunc) http.HandlerFunc

// Chain provides mechanism to concatenate middlewares
func Chain(f http.HandlerFunc, mids ...middleware) http.HandlerFunc {
	for _, m := range mids {
		f = m(f)
	}
	return f
}

var chain []middleware

func pass(args string) func(http.HandlerFunc) http.HandlerFunc {
	return func(f http.HandlerFunc) http.HandlerFunc {

		// Define the http.HandlerFunc
		return func(w http.ResponseWriter, r *http.Request) {

			//split args and check if the request as this method
			acceptedMethods := strings.Split(args, "|")
			for _, v := range acceptedMethods {
				if r.Method == v {
					// Call the next middleware in chain
					f(w, r)
					return
				}
			}
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
	}
}

func controller(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there!")
}

func main() {
	chain = append(chain, pass("GET")) //change method to thest that it work!
	http.HandleFunc("/", Chain(controller, chain...))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
