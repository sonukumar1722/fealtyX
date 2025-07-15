package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

// Student struct represents the student model used throughout the app
type Student struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

// In-memory data store using sync.Map for thread safety
var students = sync.Map{}

// Ensure student has all required valid fields
func validateStudent(s Student) error {
	if s.Name == "" || s.Email == "" || s.Age <= 0 {
		return errors.New("invalid student data")
	}
	return nil
}

// Handles both single and bulk student creation
func CreateStudent(w http.ResponseWriter, r *http.Request) {
	var multiple []Student

	// Try to decode into a slice of students for bulk insert
	if err := json.NewDecoder(r.Body).Decode(&multiple); err == nil {
		var created []Student
		for _, s := range multiple {
			if validateStudent(s) == nil {
				if _, exists := students.Load(s.ID); !exists {
					students.Store(s.ID, s)
					created = append(created, s)
				}
			}
		}
		json.NewEncoder(w).Encode(created)
		return
	}

	// Fallback to handling single student creation
	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validateStudent(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if _, exists := students.Load(s.ID); exists {
		http.Error(w, "ID already exists", http.StatusBadRequest)
		return
	}
	students.Store(s.ID, s)
	json.NewEncoder(w).Encode(s)
}

// Fetch and return all students in the system
func GetAllStudents(w http.ResponseWriter, r *http.Request) {
	var all []Student
	students.Range(func(_, value any) bool {
		all = append(all, value.(Student))
		return true
	})
	json.NewEncoder(w).Encode(all)
}

// Fetch a single student by ID
func GetStudentByID(w http.ResponseWriter, id int) {
	val, ok := students.Load(id)
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(val)
}

// Update a single student using PUT /students/{id}
func UpdateStudent(w http.ResponseWriter, r *http.Request, id int) {
	_, ok := students.Load(id)
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	var s Student
	if err := json.NewDecoder(r.Body).Decode(&s); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := validateStudent(s); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	s.ID = id // enforce path ID over body ID
	students.Store(id, s)
	json.NewEncoder(w).Encode(s)
}

// Delete a student by ID
func DeleteStudent(w http.ResponseWriter, id int) {
	_, ok := students.Load(id)
	if !ok {
		http.Error(w, "Student not found", http.StatusNotFound)
		return
	}
	students.Delete(id)
	w.Write([]byte(`{"message":"Student deleted"}`))
}

// Bulk update students. Only valid and existing entries are updated.
func UpdateMultipleStudents(w http.ResponseWriter, r *http.Request) {
	var list []Student
	if err := json.NewDecoder(r.Body).Decode(&list); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	var updated []Student
	for _, s := range list {
		if _, ok := students.Load(s.ID); ok && validateStudent(s) == nil {
			students.Store(s.ID, s)
			updated = append(updated, s)
		}
	}
	json.NewEncoder(w).Encode(updated)
}

// Bulk delete students using query param: /students?ids=1,2,3
func DeleteMultipleStudents(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("ids")
	if query == "" {
		http.Error(w, "No ids provided", http.StatusBadRequest)
		return
	}
	idList := strings.Split(query, ",")
	var deleted []int
	for _, idStr := range idList {
		id, err := strconv.Atoi(strings.TrimSpace(idStr))
		if err == nil {
			if _, ok := students.Load(id); ok {
				students.Delete(id)
				deleted = append(deleted, id)
			}
		}
	}
	json.NewEncoder(w).Encode(map[string]any{"deleted": deleted})
}
