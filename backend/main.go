package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/openai/openai-go"
)

var (
	haikuGenerator = GenerateHaiku
	db            *sql.DB // This var looks weird because it is a pointer to the database that can be used by all functions in this file.
)

type HaikuResponse struct {
	Text string `json:"text"`
}

// Define the database table
type StoredHaiku struct { 
	ID        int64     `json:"id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

// Initialize the database
func initDB() error { 
	var err error //declares a variable err of type error
	db, err = sql.Open("sqlite3", "./haikus.db") //opens a connection to the database, accepting the driver and the path to the database file.
	if err != nil { // if there is an error, return it
		return err
	}

	// Create haikus table if it doesn't exist
	//   The _ is a "blank identifier", it tells Go that you know there is a return value (the table itself) but you are ignoring it.
	_, err = db.Exec(`  
		CREATE TABLE IF NOT EXISTS haikus (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			text TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	return err
}

// Store the haiku in the database
//The ? is a placeholder for the text parameter and prevents SQL injection by escaping the input. The _ ignores the result of sql.Result because we don't need it. This is called in GenerateHaiku.
func storeHaiku(text string) error { 
	_, err := db.Exec("INSERT INTO haikus (text) VALUES (?)", text)
	return err
}

// Generate a haiku
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

	haiku := &HaikuResponse{ 
		Text: chatCompletion.Choices[0].Message.Content,
	}

	// Store the haiku in the database
	// The if statement sets the error to nil if the haiku is stored successfully. 
	if err := storeHaiku(haiku.Text); err != nil {
		log.Printf("Failed to store haiku: %v", err)
		// Don't return the error to the user, just log it
	}

	return haiku, nil
}

// This GETs the haiku from the OpenAI API, returns it to the user and stores it in the database.
func handleHaiku(w http.ResponseWriter, r *http.Request) {
	// Generate a haiku
	haiku, err := haikuGenerator(r.Context()) 
	if err != nil {
		http.Error(w, "Failed to generate haiku", http.StatusInternalServerError)
		log.Println("OpenAI error:", err)
		return
	}

	// Set the content type to json and allow cross-origin requests
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	// Encode the haiku as json and send it to the user
	if err := json.NewEncoder(w).Encode(haiku); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		log.Println("JSON encode error:", err)
	}
}

// This RETRIEVES the haikus from the database and returns them to the user.
func handleGetHaikus(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, text, created_at FROM haikus ORDER BY RANDOM() LIMIT 3") //get 3 random haikus
	if err != nil {
		http.Error(w, "Failed to fetch haikus", http.StatusInternalServerError)
		log.Println("Database error:", err)
		return
	}
	defer rows.Close() // cleans up DB resources, prevents memory leaks

	// Create a slice to store the haikus, row.Scan() adds the haiku to the slice
	var haikus []StoredHaiku
	for rows.Next() {
		var h StoredHaiku
		if err := rows.Scan(&h.ID, &h.Text, &h.CreatedAt); err != nil {
			http.Error(w, "Failed to read haiku", http.StatusInternalServerError)
			log.Println("Row scan error:", err)
			return
		}
		haikus = append(haikus, h)
	}

	// Converts the haikus to json and then sends it to the user
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(haikus)
}

func main() {
	// Initialize the database, if it fails, log the error and exit without impacting user experience.
	if err := initDB(); err != nil { // init the DB
		log.Fatal("Failed to initialize database:", err)
	}
	defer db.Close() // closes the DB properly when the program exits

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/haiku", handleHaiku) // gets the haiku from OpenAI, displays it, and stores it in the DB.
	http.HandleFunc("/haikus", handleGetHaikus) // gets the haikus from the DB and displays them.
	log.Printf("🌐 Backend running on http://localhost:%s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

// To test:
// In one terminal, cd into the backend directory and run `go run main.go`.
// In another terminal, run `curl http://localhost:8080/haiku` to get a haiku.
// Then run `curl http://localhost:8080/haikus` to get a set of 3 previously stored haiku.
