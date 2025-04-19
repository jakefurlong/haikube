package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/openai/openai-go"
)

type HaikuResponse struct {
	Text string `json:"text"`
}

func GenerateHaiku(ctx context.Context) (*HaikuResponse, error) {
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
	haiku, err := GenerateHaiku(r.Context())
	if err != nil {
		http.Error(w, "Failed to generate haiku", http.StatusInternalServerError)
		log.Println("OpenAI error:", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*") // 🔥 Allow frontend requests (dev only)
	json.NewEncoder(w).Encode(haiku)
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
