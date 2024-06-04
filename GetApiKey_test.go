package main

import (
	"net/http"
	"testing"

	"github.com/bootdotdev/learn-cicd-starter/internal/auth"
)

// TestGetApiKey test the GetApiKey method
func TestGetApiKey(t *testing.T) {
	headers := http.Header{}
	headers.Add("Authorization", "ApiKey 123456")

	apiKey, err := auth.GetAPIKey(headers)

	if err != nil {
		t.Errorf("Expected no error, got %s", err)
	}

	if apiKey != "123456" {
		t.Errorf("Expected api key 123456, got %s", apiKey)
	}
}
