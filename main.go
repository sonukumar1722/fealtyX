package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"os"
)

func main() {
	// Route for operations involving multiple students
	http.HandleFunc("/students", studentsHandler)
	
	// Route for single student operations and student summary
	http.HandleFunc("/students/", studentHandler)

	// Start the server and log if it crashes
	fmt.Println("🚀 Server running at http://localhost:8080")
	// log.Fatal(http.ListenAndServe(":8080", nil))
	port := os.Getenv("PORT")
	if port == "" {
	    port = "8080"
	}
	log.Fatal(http.ListenAndServe(":"+port, nil))
	
	}

// Handle requests related to all or multiple students
func studentsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		GetAllStudents(w, r)          // List all students
	case http.MethodPost:
		CreateStudent(w, r)           // Add one or many students
	case http.MethodPut:
		UpdateMultipleStudents(w, r)  // Bulk update
	case http.MethodDelete:
		DeleteMultipleStudents(w, r)  // Bulk delete by IDs
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// Handle operations for a specific student (by ID or for summary)
func studentHandler(w http.ResponseWriter, r *http.Request) {
	parts := strings.Split(strings.Trim(r.URL.Path, "/"), "/")
	if len(parts) < 2 {
		http.Error(w, "Student ID missing", http.StatusBadRequest)
		return
	}

	// Extract student ID from URL
	id, err := strconv.Atoi(parts[1])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	// Check if user is requesting the AI-generated summary
	if len(parts) == 3 && parts[2] == "summary" && r.Method == http.MethodGet {
		GenerateStudentSummary(w, id)
		return
	}

	// Handle regular operations on a single student
	switch r.Method {
	case http.MethodGet:
		GetStudentByID(w, id)
	case http.MethodPut:
		UpdateStudent(w, r, id)
	case http.MethodDelete:
		DeleteStudent(w, id)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
