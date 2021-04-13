package main

import (
	"log"
	"net/http"
)

func main() {
    s := &Server{}
    http.Handle("/", s)
    log.Fatal(http.ListenAndServe(":8080", nil))
}