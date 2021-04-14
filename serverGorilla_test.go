package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func returnGorrillaServerCall(action string) (string, error) {
    req, err := http.NewRequest(action, "/", nil)

    if err != nil {
        return "", err
    }

    w := httptest.NewRecorder()
    server := &myServer{}
    router := server.InitRouter()
    router.ServeHTTP(w, req)

    return  w.Body.String(), err
}


func TestGetRequest_GET(t *testing.T){
    mesg, err  := returnGorrillaServerCall("GET")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "get called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_POST(t *testing.T){
    mesg, err  := returnGorrillaServerCall("POST")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "post called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_DELETE(t *testing.T){
    mesg, err  := returnGorrillaServerCall("DELETE")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "delete called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_DEFAULT(t *testing.T){
    mesg, err  := returnGorrillaServerCall("humhum")
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "not found"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}