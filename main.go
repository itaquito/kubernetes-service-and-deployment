package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type Message struct {
	Message  string `json:"message"`
	Status   string `json:"status"`
	Instance string `json:"instance"`
	Port     string `json:"port"`
}

func messageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}

	response := Message{
		Message:  "Hello from Go server!",
		Status:   "success",
		Instance: hostname,
		Port:     "8080",
	}

	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/api", messageHandler)

	log.Println("Server starting on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
