package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Message wraps a message and stamps it
type Message struct {
	Message string `json:"message"` //tells the decoder what to decode into
	Stamp   int64  `json:"stamp,omitempty"`
}

//BuisnessLogic does awesome BuisnessLogic
func BuisnessLogic(text string) *Message {
	mess := &Message{}
	mess.Message = text
	mess.Stamp = time.Now().Unix()
	return mess
}

//Echo echoes what you send
func Echo(res http.ResponseWriter, req *http.Request) {
	var (
		jsonDecoder = json.NewDecoder(req.Body) //decoder reading from the post body
		jsonEncoder = json.NewEncoder(res)      //encoder writing to the response stream
		message     = &Message{}         // something to hold our data
	)
	res.Header().Add("Content-type", "application/json")
	if err := jsonDecoder.Decode(message); err != nil { //decode our data into our struct
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	pointless := BuisnessLogic(message.Message)
	if err := jsonEncoder.Encode(pointless); err != nil { //encode our data and write it back to the response stream
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Setup our simple router
func router() http.Handler {
	http.HandleFunc("/api/echo", Echo)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func main() {
	router := router()
	//start our server on port 3001
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
