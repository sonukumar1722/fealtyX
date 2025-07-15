package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Request payload structure for Ollama API
type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

// Expected response from Ollama API
type OllamaResponse struct {
	Response string `json:"response"`
}

// Generate a professional summary of a student using AI (Ollama)
func GenerateStudentSummary(w http.ResponseWriter, id int) {
	val, ok := students.Load(id)
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	s := val.(Student)

	// Prompt to be sent to the AI model
	prompt := fmt.Sprintf("Generate a brief professional summary for this student:\nName: %s\nAge: %d\nEmail: %s", s.Name, s.Age, s.Email)

	// New request structure for /api/chat
	requestData := map[string]interface{}{
		"model": "llama3",
		"messages": []map[string]string{
			{
				"role":    "user",
				"content": prompt,
			},
		},
		"stream": false,
	}

	jsonPayload, _ := json.Marshal(requestData)

	// POST to /api/chat (correct endpoint and structure)
	resp, err := http.Post("http://localhost:11434/api/chat", "application/json", bytes.NewBuffer(jsonPayload))
	if err != nil {
		http.Error(w, "Failed to contact Ollama", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	// Parse the nested "message.content" field from the response
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		http.Error(w, "Invalid response from Ollama", http.StatusInternalServerError)
		return
	}

	message, ok := result["message"].(map[string]interface{})
	if !ok {
		http.Error(w, "Malformed Ollama response", http.StatusInternalServerError)
		return
	}

	content, ok := message["content"].(string)
	if !ok {
		http.Error(w, "Missing summary content", http.StatusInternalServerError)
		return
	}

	// Send the summary back to the frontend/client
	json.NewEncoder(w).Encode(map[string]string{"summary": content})
}
