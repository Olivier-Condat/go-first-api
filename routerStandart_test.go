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

    s := &myServer{}
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

 func TestServeHTTP_PUT(t *testing.T){

    mesg, err  := returnServerCall("PUT")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "put called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
 }

 func TestServeHTTP_DELETE(t *testing.T){

    mesg, err  := returnServerCall("DELETE")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "delete called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
 }

 func TestServeHTTP_DEFAULT(t *testing.T){

    mesg, err  := returnServerCall("this-is-not-covered")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "not found"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
 }