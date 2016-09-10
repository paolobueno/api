package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	//create a test server based from our router
	server := httptest.NewServer(router())
	defer server.Close() //notice we use defer here to ensure our server is closed
	//Make a new request using the test server url
	res, err := http.NewRequest("POST", server.URL+"/api/echo", strings.NewReader(`{"message":"test"}`))
	if err != nil {
		log.Fatal(err)
	}
	//readAll of the body back as as bytes
	echo, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	message := &Message{} //allocate a message to store our data in
	//Unmarshal takes a slice / array of bytes and converts it into the passed type
	if err := json.Unmarshal(echo, message); err != nil {
		log.Fatal(err)
	}
	if "test2" != message.Message {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}
