package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

// User represents a simple user with a Name field.
type User struct {
	Name string `json:"name"`
}

// userCache is an in-memory store for user data.
// It maps an integer ID to a User object.
var userCache = make(map[int]User)

// cacheMutex is used to safely access userCache in concurrent environments.
var cacheMutex sync.RWMutex

func main() {
	// Create a new multiplexer (router).
	mux := http.NewServeMux()

	// Define HTTP handlers.
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("POST /user", createUser)  // Custom Go 1.22+ pattern
	mux.HandleFunc("GET /user/{id}", getUser) // Go 1.22+ dynamic path segment
	mux.HandleFunc("DELETE /user/{id}", deleteUser)

	// Start the server on port 8080.
	fmt.Println("Server listening on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

// rootHandler returns a simple greeting for root endpoint.
func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Client!")
}

// createUser handles POST /user to create a new user.
func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	// Decode JSON request body into User struct.
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validate user name is not empty.
	if user.Name == "" {
		http.Error(w, "User name can't be empty!", http.StatusBadRequest)
		return
	}

	// Lock cache for writing and assign a new ID.
	cacheMutex.Lock()
	userID := len(userCache) + 1
	userCache[userID] = user
	cacheMutex.Unlock()

	// Respond with status 202 Accepted.
	w.WriteHeader(http.StatusAccepted)
}

// getUser handles GET /user/{id} to retrieve a user by ID.
func getUser(w http.ResponseWriter, r *http.Request) {
	// Extract and parse the path variable `{id}`.
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Lock cache for reading.
	cacheMutex.RLock()
	user, ok := userCache[id]
	cacheMutex.RUnlock()

	if !ok {
		http.Error(w, "User ID not found", http.StatusNotFound)
		return
	}

	// Return user as JSON.
	w.Header().Set("Content-Type", "application/json")
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error encoding response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// deleteUser handles DELETE /user/{id} to remove a user by ID.
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	// Check if user exists before deletion.
	cacheMutex.Lock()
	defer cacheMutex.Unlock()

	if _, ok := userCache[id]; !ok {
		http.Error(w, "User ID not found", http.StatusNotFound)
		return
	}

	// Delete user.
	delete(userCache, id)
	w.WriteHeader(http.StatusNoContent) // 204 No Content
}
