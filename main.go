package main

import (
	"log"
	"net/http"
)

func main() {
    server := &myServer{}
    router := server.InitRouter()
    log.Fatal(http.ListenAndServe(":8080", router))
}
