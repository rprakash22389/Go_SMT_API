package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)


func main() {
	http.HandleFunc("/software-details", handleSoftwareDetails)
	http.HandleFunc("/autocomplete", handleAutocomplete)
	http.HandleFunc("/post-user-details", handlePostUserDetails)
	http.HandleFunc("/servicenow-tickets", handleServiceNowTickets)

	fmt.Println("Server is running on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

// 1. Software Package Details API
func handleSoftwareDetails(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	email := r.URL.Query().Get("email")

	if username == "" || email == "" {
		http.Error(w, "Username and email are required", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success":   true,
		"username":  username,
		"email":     email,
		"software": []map[string]string{
			{
				"name":        "GoLang",
				"description": "A statically typed, compiled programming language designed for simplicity and performance.",
				"version":     "1.20.3",
				"link":        "https://golang.org",
			},
			{
				"name":        "Visual Studio Code",
				"description": "A lightweight code editor with support for various languages.",
				"version":     "1.76.2",
				"link":        "https://code.visualstudio.com/",
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 2. Autocomplete API
func handleAutocomplete(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q") // The search query
	usernames := []map[string]string{
		{"username": "johndoe", "email": "johndoe@example.com"},
		{"username": "janedoe", "email": "janedoe@example.com"},
		{"username": "johnsmith", "email": "johnsmith@example.com"},
		{"username": "janesmith", "email": "janesmith@example.com"},
	}

	var results []map[string]string
	for _, user := range usernames {
		if query != "" && (contains(user["username"], query) || contains(user["email"], query)) {
			results = append(results, user)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]interface{}{
		"success": true,
		"results": results,
	})
}

// Helper function to check if a string contains a substring
func contains(value, query string) bool {
	return len(value) >= len(query) && value[:len(query)] == query
}

// 3. Post User Details API
func handlePostUserDetails(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var requestBody map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success": true,
		"message": "User details received",
		"details": requestBody,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// 4. ServiceNow Ticket Details API
func handleServiceNowTickets(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username is required", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"success":  true,
		"username": username,
		"tickets": []map[string]interface{}{
			{"ticket_id": "INC0012345", "status": "Resolved", "description": "System issue resolved."},
			{"ticket_id": "INC0012346", "status": "Open", "description": "Pending network configuration."},
		},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
