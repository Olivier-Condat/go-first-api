package main

import (
	"net/http"

	//sw "github.com/Olivier-Condat/openapi-go-first-api-test"
	"github.com/gorilla/mux"
)

func (server *myServer) InitRouter() *mux.Router{
    router := mux.NewRouter()
    router.HandleFunc("/", get).Methods(http.MethodGet)
    router.HandleFunc("/", post).Methods(http.MethodPost)
    router.HandleFunc("/", put).Methods(http.MethodPut)
    router.HandleFunc("/", delete).Methods(http.MethodDelete)
    router.HandleFunc("/", notFound)

    // configuration := sw.NewConfiguration()
    // api_client := sw.NewAPIClient(configuration)
    
    // resp, r, err := api_client.DefaultApi.DataGet(context.Background()).Execute()
    // if err != nil {
    //     fmt.Fprintf(os.Stderr, "Error when calling `DefaultApi.DataGet``: %v\n", err)
    //     fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    // }
    // // response from `DataGet`: string
    // fmt.Fprintf(os.Stdout, "Response from `DefaultApi.DataGet`: %v\n", resp)

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