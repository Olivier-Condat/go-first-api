package main

import (
	"net/http"

	//sw "github.com/Olivier-Condat/openapi-go-first-api-test"
	"github.com/gorilla/mux"
)

func (server *myServer) InitRouter() *mux.Router{
    router := mux.NewRouter()
    router.HandleFunc("/api/v1", GetFunction).Methods(http.MethodGet)
    
    router.HandleFunc("/", get).Methods(http.MethodGet)
    router.HandleFunc("/", post).Methods(http.MethodPost)
    router.HandleFunc("/", put).Methods(http.MethodPut)
    router.HandleFunc("/", delete).Methods(http.MethodDelete)
    router.HandleFunc("/", notFound)
    return router
}

func get(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    res.Write([]byte(`{"message": "get called"}`))
}

func post(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    w.Write([]byte(`{"message": "post called"}`))
}

func put(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusAccepted)
    w.Write([]byte(`{"message": "put called"}`))
}

func delete(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "delete called"}`))
}

func notFound(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte(`{"message": "not found"}`))
}