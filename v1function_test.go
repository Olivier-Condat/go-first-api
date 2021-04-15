package main

import (
	"math"
	"net/url"
	"testing"

	"github.com/gorilla/schema"
)

func TestMessageFromJson(t *testing.T){

    var expectedTemp float64 = 23.475861
    var expectedHumidity float64 =  9.674

    message := Message{expectedTemp,expectedHumidity}
    form := url.Values{}

    encoder := schema.NewEncoder()
    err := encoder.Encode(message, form)

    if(err != nil){
        t.Fatalf("\nError: %v" , err)
    }

    messageDecoded := MessageFromJson(form)

    diff :=  math.Abs( float64 (messageDecoded.Temperature) -  23.475861)
    if( diff > 0.0001  ){
        t.Fatalf("\nExpected \n\t temperature: '%f' but got '%f'", 23.47586, messageDecoded.Temperature)
    }

    diff = math.Abs(float64 (messageDecoded.Humidity) -   9.674)
    if(  diff > 0.0001  ){
        t.Fatalf("\nExpected \n\t humidity: '%f' but got '%f'", 9.674, message.Humidity)
    }
 }

 func TestFloatToString(t *testing.T){
    var value = 2034923904.0298409238409280498
    expectedValue := "2034923904.029841"

    testValue := FloatToString(value)

    if(expectedValue != testValue){
        t.Fatalf("\nExpected \n\t humidity: '%s' but got '%s'", expectedValue, testValue)
    }
 }