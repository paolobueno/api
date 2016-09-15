package api

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestEchoCmd(t *testing.T) {
	mockPoster := func(api, contentType string, body io.Reader) (*http.Response, error) {
		u, err := url.Parse(api)
		if err != nil {
			t.Fatal(err.Error())
		}
		if u.Path != "/api/echo" {
			t.Fatal(err.Error())
		}
		resBody := `{"message":"test"}`
		bodyRC := ioutil.NopCloser(bytes.NewReader([]byte(resBody)))
		res := &http.Response{StatusCode: 200, Body: bodyRC}
		return res, nil
	}
	var buffer bytes.Buffer
	cmd := &echoCmd{poster: mockPoster, writer: &buffer}
	args := []string{"test"}
	if err := cmd.Echo(args); err != nil {
		t.Fatal(err.Error)
	}
	message := map[string]string{}
	if err := json.Unmarshal(buffer.Bytes(), &message); err != nil {
		t.Fatal(err.Error())
	}
	if _, ok := message["message"]; !ok {
		t.Fatal("expected their to be a message")
	}
	m := message["message"]
	if m != "test" {
		t.Fatal("expected the messag to be test")
	}
}
