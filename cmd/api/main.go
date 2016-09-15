package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var hostFlag = flag.String("host", "http://localhost:3001", "host is used to change the default host to call")

const (
	usage = "available commands: \n echo \n map \n slice "
	ECHO  = "echo"
	MAP   = "map"
	SLICE = "slice"
)

func main() {
	var err error
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	args := flag.Args()
	cmd := args[0]
	switch cmd {
	case ECHO:
		err = echoCmd(args[1:])
	}
	log.Fatalf("error running command %s : %s", cmd, err.Error())
}

func echoCmd(args []string) error {
	if len(args) != 1 {
		printAndExit("echo expects a message. echo <yourmessage>")
	}
	url := fmt.Sprintf("%s/api/echo", *hostFlag)
	msg := map[string]string{
		"message": args[0],
	}
	data, err := json.Marshal(msg)
	if err != nil {
		return fmt.Errorf("failed to prepare data for posting %s ", err.Error())
	}
	res, err := http.Post(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to post to echo endpoint %s ", err.Error())
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to post to echo endpoint %s ", err.Error())
	}
	log.Println(string(resData))
	return nil
}

func printAndExit(msg string) {
	log.Println(msg)
	os.Exit(0)
}
