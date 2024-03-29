package webserver

import (
	"net/http"
	"net/url"
	"os"
	"testing"
)

func TestSendGETRequestToLocalhost(t *testing.T) {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		t.Fatal("SERVER_URL environment variable is not set")
		return
	}

	resp, err := http.Get(serverURL)
	if err != nil {
		t.Errorf("Failed to send GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}
}

func TestSendGETRequestToLocalhostOnApi(t *testing.T) {
	serverURL := os.Getenv("SERVER_URL")
	if serverURL == "" {
		t.Fatal("SERVER_URL environment variable is not set")
		return
	}

	getUrl, _ := url.JoinPath(serverURL, "/api")
	resp, err := http.Get(getUrl)
	if err != nil {
		t.Errorf("Failed to send GET request: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("Expected status code %d, but got %d", http.StatusNotFound, resp.StatusCode)
	}
}
