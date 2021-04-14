package main

import (
	"net/http"

	openapiclient "./openapi"

	"github.com/gorilla/mux"
)

func (server *myServer) InitRouter() *mux.Router{
    router := mux.NewRouter()
    router.HandleFunc("/", get).Methods(http.MethodGet)
    router.HandleFunc("/", post).Methods(http.MethodPost)
    router.HandleFunc("/", put).Methods(http.MethodPut)
    router.HandleFunc("/", delete).Methods(http.MethodDelete)
    router.HandleFunc("/", notFound)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    router.HandleFunc("/data", api_client.DefaultApi.ApiDataGetRequest).Methods(http.MethodGet)

    return router
}

func get(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    w.Write([]byte(`{"message": "get called"}`))
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