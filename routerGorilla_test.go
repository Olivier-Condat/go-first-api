package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func returnGorrillaServerCall(action string, path string, body io.Reader) (string, error) {
    req, err := http.NewRequest(action, path, body)

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
    mesg, err  := returnGorrillaServerCall("GET", "/", nil)
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "get called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_GET_V1(t *testing.T){
    mesg, err  := returnGorrillaServerCall("GET", "/api/v1", nil)
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message from GetFunction": "v1 get called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_POST(t *testing.T){
    mesg, err  := returnGorrillaServerCall("POST", "/", nil)
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "post called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_POST_V1(t *testing.T){
    
    var jsonStr = []byte(``)
    mesg, err  := returnGorrillaServerCall("POST", "/api/v1", bytes.NewBuffer(jsonStr))
    
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message from GetFunction v1": "wrong request format"}`
    if( mesg != expectedBody ){
       t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }

    jsonStr = []byte(`{"temperature":29.09972319328132,"humidity":79.66377998881264}`)

    mesg, err = returnGorrillaServerCall("POST", "/api/v1", bytes.NewBuffer(jsonStr))
    
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody = `{"temperature: "29.099723", humidity: "79.663780"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_DELETE(t *testing.T){
    mesg, err  := returnGorrillaServerCall("DELETE", "/", nil)
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "delete called"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}

func TestGetRequest_DEFAULT(t *testing.T){
    mesg, err  := returnGorrillaServerCall("humhum","/", nil)
    if( err != nil ){
        t.Fatal(err)
    }

    expectedBody := `{"message": "not found"}`
    if mesg != expectedBody {
        t.Fatalf("\nExpected \n\t Body: '%s' but got '%s'", expectedBody, mesg)
    }
}