package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type OpenAIChoice struct {
	Text string `json:"text"`
}

type OpenAIResponse struct {
	Choices []OpenAIChoice `json:"choices"`
}

type OpenAIRequest struct {
	Model string `json:"model"`
	Input string `json:"prompt"`
}

func main() {
	// Retrieve the API key from the environment variable
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatal("API key not found. Please set the OPENAI_API_KEY environment variable.")
	}

	// Define the OpenAI API endpoint
	url := "https://api.openai.com/v1/completions"

	// Create the request payload
	requestPayload := OpenAIRequest{
		Model: "gpt-4o",
		Input: "Write a humorous devops haiku.",
	}

	// Marshal the request payload into JSON
	requestBody, err := json.Marshal(requestPayload)
	if err != nil {
		log.Fatalf("Error marshalling request payload: %v", err)
	}

	// Create a new HTTP request
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		log.Fatalf("Error creating HTTP request: %v", err)
	}

	// Add the necessary headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	// Make the API request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error making request: %v", err)
	}
	defer resp.Body.Close()

	// Check for a successful response
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Request failed with status: %s", resp.Status)
	}

	// Decode the response body into a struct
	var response OpenAIResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Fatalf("Error decoding response: %v", err)
	}

	// Debugging: print the raw response to see the structure
	fmt.Println("Raw response:", response)

	// Print the generated text if available
	if len(response.Choices) > 0 {
		fmt.Println(response.Choices[0].Text)
	} else {
		log.Println("No choices received in the response.")
	}
}
