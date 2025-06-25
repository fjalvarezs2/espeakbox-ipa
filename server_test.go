package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestVoicesHandler(t *testing.T) {
	// Avoid external command execution by setting the cached response
	cachedVoicesJSON = []byte(`{"names": ["en"]}`)

	mux := http.NewServeMux()
	mux.HandleFunc("/voices", voicesHandler)
	ts := httptest.NewServer(mux)
	defer ts.Close()

	resp, err := http.Get(ts.URL + "/voices")
	if err != nil {
		t.Fatalf("failed to GET /voices: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	var data map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		t.Fatalf("failed to decode JSON: %v", err)
	}

	if _, ok := data["names"]; !ok {
		t.Fatalf("response JSON missing 'names' key")
	}
}
