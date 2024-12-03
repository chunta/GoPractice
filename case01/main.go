package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var items []Item
var mu sync.Mutex

func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem Item
	if err := json.NewDecoder(r.Body).Decode(&newItem); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	println(len(newItem.ID))

	mu.Lock()
	if newItem.ID == "" {
		newItem.ID = fmt.Sprintf("%d", len(items)+1) // Auto-generate string ID
	}
	items = append(items, newItem)
	mu.Unlock()

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

func itemsHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getItems(w, r)
	case http.MethodPost:
		createItem(w, r)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/items", itemsHandler)
	http.ListenAndServe(":8080", nil)
}
