package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
    //Azure manages himself the available port
    customHandlerPort, exists := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT")
    if !exists {
        customHandlerPort = "8080"
    }

    server := &myServer{}
    router := server.InitRouter()
    log.Fatal(http.ListenAndServe(":"+customHandlerPort, router))
}
