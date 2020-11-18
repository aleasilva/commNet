package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopSecretSplit(test *testing.T) {
	testServer := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	req, err := http.NewRequest("POST", testServer.URL+"/topsecret_split/luisa",
		bytes.NewBuffer(creatTopSecretSplitMessage()))

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 400 {
		test.Fatalf("Expected status code 400, got %v", resp.StatusCode)
	}

	//Test error 200
	req, err = http.NewRequest("POST", testServer.URL+"/topsecret_split/skywalker",
		bytes.NewBuffer(creatTopSecretSplitMessage()))

	req.Header.Set("Content-Type", "application/json")

	client = &http.Client{}
	resp, err = client.Do(req)

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		test.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

//creatTopSecretMessage Create an message to test /topsecret call.
func creatTopSecretSplitMessage() []byte {
	var jsonStr = []byte(`
		{
			"distance": 100.0,
			"message": ["este", "", "", "mensaje", ""]  
		}  
	`)

	return jsonStr
}

func TestTopSecretCall(test *testing.T) {
	testServer := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	req, err := http.NewRequest("POST", testServer.URL+"/topsecret", bytes.NewBuffer(creatTopSecretMessage()))
	req.Header.Set("X-Custom-Header", "myvalue")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		test.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}

//creatTopSecretMessage Create an message to test /topsecret call.
func creatTopSecretMessage() []byte {
	var jsonStr = []byte(`
		{
		  "satellites":[ 
							{
								"name": "kenobi",
								"distance": 100.0,
								"message": ["este", "", "", "mensaje", ""]                       
							},
							{
								"name": "skywalker",  
								"distance": 115.5,
								"message": ["", "es", "", "", "secreto"]
							},
							{
								"name": "sato",  
								"distance": 142.7,
								"message": ["este", "", "un", "", ""]
							}
						]
		}
	`)

	return jsonStr
}

//Test if the server is responde
func TestPingEndpoint(test *testing.T) {
	testServer := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))

	if err != nil {
		test.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		test.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
