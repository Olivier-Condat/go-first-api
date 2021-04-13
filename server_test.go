package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func returnServerCall(action string) (string, error) {
    req, err := http.NewRequest(action, "http://localhost:8080", nil)

    if err != nil {
        return "", err
    }

    w := httptest.NewRecorder()

    s := &Server{}
    handler := http.HandlerFunc(s.ServeHTTP)

    handler.ServeHTTP(w, req)

    return  w.Body.String(), err
}

func TestServeHTTP_GET(t *testing.T){

    mesg, err  := returnServerCall("GET")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "get called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
 }

 func TestServeHTTP_POST(t *testing.T){

    mesg, err  := returnServerCall("POST")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "post called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
 }