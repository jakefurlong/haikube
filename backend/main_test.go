package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGenerateHaiku(t *testing.T) {
	t.Skip("Live test – remove Skip to hit OpenAI API")

	resp, err := haikuGenerator(context.Background())
	if err != nil {
		t.Fatalf("haikuGenerator returned error: %v", err)
	}

	if resp == nil || resp.Text == "" {
		t.Error("Expected non-empty haiku text")
	}
}

func TestHandleHaiku(t *testing.T) {
	// ✅ This ensures you're overriding the correct function
	original := haikuGenerator
	haikuGenerator = func(ctx context.Context) (*HaikuResponse, error) {
		return &HaikuResponse{
			Text: "Mocked haiku\nfrom the backend\nfor testing only",
		}, nil
	}
	defer func() { haikuGenerator = original }()

	req := httptest.NewRequest("GET", "/haiku", nil)
	rec := httptest.NewRecorder()

	handleHaiku(rec, req)

	res := rec.Result()
	defer func() {
		if err := res.Body.Close(); err != nil {
			t.Errorf("Failed to close response body: %v", err)
		}
	}()

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200, got %d", res.StatusCode)
	}

	var parsed HaikuResponse
	err := json.NewDecoder(res.Body).Decode(&parsed)
	if err != nil {
		t.Fatalf("Failed to decode response body: %v", err)
	}

	if !strings.Contains(parsed.Text, "Mocked haiku") {
		t.Errorf("Expected mocked haiku in response, got: %s", parsed.Text)
	}
}
