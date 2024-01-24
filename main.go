package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
)

type User struct {
	UserID      int    `json:"user_id"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	NocTransfer int    `json:"noc_transfer"`
}

// DataStore represents an in-memory data store for users
type DataStore struct {
	mu     sync.RWMutex
	users  map[int]User
	nextID int
}

// NewDataStore creates a new instance of DataStore
func NewDataStore() *DataStore {
	return &DataStore{
		users:  make(map[int]User),
		nextID: 1,
	}
}

// Create adds a new user to the data store
func (ds *DataStore) Create(user User) int {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	user.UserID = ds.nextID
	ds.users[user.UserID] = user
	ds.nextID++

	return user.UserID
}

// Read retrieves a user from the data store based on userID
func (ds *DataStore) Read(userID int) (User, bool) {
	ds.mu.RLock()
	defer ds.mu.RUnlock()

	user, ok := ds.users[userID]
	return user, ok
}

// Update updates the information of an existing user
func (ds *DataStore) Update(userID int, newUser User) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.users[userID]; ok {
		newUser.UserID = userID
		ds.users[userID] = newUser
		return true
	}

	return false
}

// Delete removes a user from the data store
func (ds *DataStore) Delete(userID int) bool {
	ds.mu.Lock()
	defer ds.mu.Unlock()

	if _, ok := ds.users[userID]; ok {
		delete(ds.users, userID)
		return true
	}

	return false
}

func main() {
	dataStore := NewDataStore()

	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUsers(w, dataStore)
		case http.MethodPost:
			createUser(w, r, dataStore)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getUser(w, r, dataStore)
		case http.MethodPut:
			updateUser(w, r, dataStore)
		case http.MethodDelete:
			deleteUser(w, r, dataStore)
		default:
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", nil)
}

func getUsers(w http.ResponseWriter, dataStore *DataStore) {
	dataStore.mu.RLock()
	defer dataStore.mu.RUnlock()

	users := make([]User, 0, len(dataStore.users))
	for _, user := range dataStore.users {
		users = append(users, user)
	}

	sendJSONResponse(w, users)
}

func createUser(w http.ResponseWriter, r *http.Request, dataStore *DataStore) {
	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userID := dataStore.Create(newUser)
	sendJSONResponse(w, map[string]int{"user_id": userID})
}

func getUser(w http.ResponseWriter, r *http.Request, dataStore *DataStore) {
	userID := extractIDFromURL(r)
	user, ok := dataStore.Read(userID)
	if !ok {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, user)
}

func updateUser(w http.ResponseWriter, r *http.Request, dataStore *DataStore) {
	userID := extractIDFromURL(r)

	var newUser User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	if !dataStore.Update(userID, newUser) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, map[string]string{"message": "User updated successfully"})
}

func deleteUser(w http.ResponseWriter, r *http.Request, dataStore *DataStore) {
	userID := extractIDFromURL(r)

	if !dataStore.Delete(userID) {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	sendJSONResponse(w, map[string]string{"message": "User deleted successfully"})
}

func extractIDFromURL(r *http.Request) int {
	var id int
	fmt.Sscanf(r.URL.Path, "/users/%d", &id)
	return id
}

func sendJSONResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
