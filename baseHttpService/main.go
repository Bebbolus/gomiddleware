package main

import (
    "fmt"
    "log"
    "net/http"
)

func controller(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there!")
}

func main() {
    http.HandleFunc("/", controller)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
