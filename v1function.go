package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"
	"strconv"

	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

type Message struct {
    Temperature float64 `json:"temperature"`
    Humidity float64 `json:"humidity"`
}

func GetFunction(res http.ResponseWriter, req *http.Request) {
    res.Header().Set("Content-Type", "application/json")
    res.WriteHeader(http.StatusOK)
    res.Write([]byte(`{"message from GetFunction": "v1 get called"}`))
}

type Test struct {
    Temperature float64 `json:"temperature"`
    Humidity float64 `json:"humidity"`
}

func PostFunction(res http.ResponseWriter, req *http.Request){
    err := req.ParseForm()

    if err != nil {
        log.Printf("Error in function PostFunction: %v" , err)
        return
    }

    res.Header().Set("Content-Type", "application/json")

    var message Message
 
    err = json.NewDecoder(req.Body).Decode(&message)
    if err != nil {
        res.WriteHeader(http.StatusBadRequest)
        res.Write([]byte(`{"message from GetFunction v1": "wrong request format"}`))
        return
    }
 
    res.WriteHeader(http.StatusOK)

    var result string
    result = ` get a temperature value of: ` + FloatToString(message.Temperature)
    result += ` and a humidity value of: ` +  FloatToString(message.Humidity)

    res.Write([]byte(result))
}

// not used anymore but could be interesting in the future
func MessageFromJson(urls url.Values) Message{
    message := Message{}
    _ = decoder.Decode(&message, urls)
    return message
}

func FloatToString(value float64) string {
    return strconv.FormatFloat(value, 'f', 6, 64)
}
