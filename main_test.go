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
	server := httptest.NewServer(router())
	defer server.Close() //notice we use defer here to ensure our server is closed
	res, err := http.NewRequest("POST", server.URL+"/api/echo", strings.NewReader(`{"message":"test"}`))
	if err != nil {
		log.Fatal(err)
	}
	echo, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	message := &Message{}
	if err := json.Unmarshal(echo, message); err != nil {
		log.Fatal(err)
	}
        log.Println(message)
	if "test2" != message.Message {
		t.Fail()
		log.Println("expected the message to equal test")
	}
}

