package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestTopSecretCall(t *testing.T) {

}

//Test if the server is responde
func TestPingEndpoint(t *testing.T) {
	testServer := httptest.NewServer(SetupServer())

	// Shut down the server and block until all requests have gone through
	defer testServer.Close()

	// Make a request to our server with the {base url}/ping
	resp, err := http.Get(fmt.Sprintf("%s/ping", testServer.URL))

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if resp.StatusCode != 200 {
		t.Fatalf("Expected status code 200, got %v", resp.StatusCode)
	}
}
