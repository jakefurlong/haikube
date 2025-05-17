package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/openai/openai-go"
)

var haikuGenerator = GenerateHaiku

type HaikuResponse struct {
	Text string `json:"text"`
}

func GenerateHaiku(ctx context.Context) (*HaikuResponse, error) {
	key, err := getAPIKeyFromSecretsManager(ctx, "openai/api-key")
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve OpenAI API key: %w", err)
	}

	// Set the key as an environment variable, so openai-go can use it
	if err := os.Setenv("OPENAI_API_KEY", key); err != nil {
		return nil, fmt.Errorf("failed to set environment variable: %w", err)
	}

	client := openai.NewClient()

	chatCompletion, err := client.Chat.Completions.New(ctx, openai.ChatCompletionNewParams{
		Messages: []openai.ChatCompletionMessageParamUnion{
			openai.UserMessage("Write a humorous devops haiku."),
		},
		Model: openai.ChatModelGPT4o,
	})
	if err != nil {
		return nil, err
	}

	return &HaikuResponse{
		Text: chatCompletion.Choices[0].Message.Content,
	}, nil
}

func handleHaiku(w http.ResponseWriter, r *http.Request) {
	haiku, err := haikuGenerator(r.Context())
	if err != nil {
		http.Error(w, "Failed to generate haiku", http.StatusInternalServerError)
		log.Println("OpenAI error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if err := json.NewEncoder(w).Encode(haiku); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("JSON encode error:", err)
	}
}

func getAPIKeyFromSecretsManager(ctx context.Context, secretName string) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-west-1"))
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	svc := secretsmanager.NewFromConfig(cfg)

	resp, err := svc.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	if err != nil {
		return "", fmt.Errorf("failed to get secret: %w", err)
	}

	if resp.SecretString == nil {
		return "", fmt.Errorf("secret string is nil")
	}

	return *resp.SecretString, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/haiku", handleHaiku)
	log.Printf("🌐 Backend running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
