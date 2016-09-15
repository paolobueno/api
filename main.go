package main

//import our dependencies

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

//Message wraps a message and stamps it
//Define a custom type to hold our data.
type Message struct {
	Message string `json:"message"` //tells the decoder what to decode into this is needed if the property is lowercase and the property is uppercase
	Stamp   int64  `json:"stamp,omitempty"`
}

//BusinessLogic does awesome BusinessLogic taking a string and returning a pointer to a Message
func BusinessLogic(text string) *Message {
	mess := &Message{}
	mess.Message = text
	mess.Stamp = time.Now().Unix()
	return mess
}

//Echo echoes what you send
//http.ResponseWriter is responsible for writing things back to the response stream
//http.Request represents the incoming request
func Echo(res http.ResponseWriter, req *http.Request) {
	var (
		jsonDecoder = json.NewDecoder(req.Body) //decoder reading from the post body
		jsonEncoder = json.NewEncoder(res)      //encoder writing to the response stream
		message     = &Message{}                // something to hold our data
	)
	//Add a content type header
	res.Header().Add("Content-type", "application/json")
	//decode our data into our struct. Notice the assignment and the err check can be done is a single line as long as you use the ;
	if err := jsonDecoder.Decode(message); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
	//call our BusinessLogic function and assign the return value
	pointless := BusinessLogic(message.Message)
	//Encode the Message contained in pointless and write it back to the response.
	if err := jsonEncoder.Encode(pointless); err != nil { //encode our data and write it back to the response stream
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

type TimeHolder struct {
	Time int64 `json:"time"`
}

func Time(res http.ResponseWriter, req *http.Request) {
	var (
		timeHolder = &TimeHolder{}
		encoder    = json.NewEncoder(res)
	)
	res.Header().Add("Content-type", "application/json")
	timeHolder.Time = time.Now().Unix()

	if err := encoder.Encode(timeHolder); err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}

//Setup our simple router
func router() http.Handler {
	//http.HandleFunc expects a func that takes a http.ResponseWriter and http.Request
	http.HandleFunc("/api/echo", Echo)
	http.HandleFunc("/api/time", Time)
	return http.DefaultServeMux //this is a stdlib http.Handler
}

func main() {
	router := router()
	//start our server on port 3001
	if err := http.ListenAndServe(":3001", router); err != nil {
		log.Fatal(err)
	}
}
