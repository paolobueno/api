//Package api is the root of the cli for our rest api and lives in cmd/api it is a convention in Golang to have "sub" binaries
//live inside a cmd package
package api

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

//configure our flags using the builtin flags lib
var hostFlag = flag.String("host", "http://localhost:3001", "host is used to change the default host to call")

//declare some constants to use
const (
	usage = "available commands: \n echo \n map \n slice "
	ECHO  = "echo"
	MAP   = "map"
	SLICE = "slice"
)

func main() {
	var err error
	//set up some usage info
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Println(usage)
		flag.PrintDefaults()
	}
	//ensure we parse our flags
	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}
	args := flag.Args() //cuts off things like ./api and any flags that are passed
	cmd := args[0]
	switch cmd {
	case ECHO:
		//configure the poster as the default http implementation and the writer as stdout
		echo := echoCmd{poster: http.Post, writer: os.Stdout}
		err = echo.Echo(args[1:]) //pass in all after echo
	}
	log.Fatalf("error running command %s : %s", cmd, err.Error())
}

//define a custom type that takes a poster and a writer this allows for cleaner simpler testing
type echoCmd struct {
	poster func(string, string, io.Reader) (*http.Response, error)
	writer io.Writer
}

//Echo calls the echo api in the web server writing the respose to the writer
func (cmd echoCmd) Echo(args []string) error {
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
	res, err := cmd.poster(url, "application/json", bytes.NewReader(data))
	if err != nil {
		return fmt.Errorf("failed to post to echo endpoint %s ", err.Error())
	}
	defer res.Body.Close()
	resData, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("failed to post to echo endpoint %s ", err.Error())
	}
	if _, err := cmd.writer.Write(resData); err != nil {
		return fmt.Errorf("failed to write to out %s", err.Error())
	}
	return nil
}

func printAndExit(msg string) {
	log.Println(msg)
	os.Exit(0)
}
