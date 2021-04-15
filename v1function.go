package main

import "net/http"

func GetFunction(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    res.Write([]byte(`{"message from GetFunction": "v1 get called"}`))
}