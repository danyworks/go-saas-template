package handlers

import (
	"encoding/json"
	"marketplace/internal/user"
	"net/http"
	"strconv"
	"sync"
)

var (
	users     = make(map[uint64]*user.User)
	userID    = uint64(0)
	userMutex sync.Mutex
	waitGroup sync.WaitGroup
)

// usersHandler handles requests to GET and POST /users
func UsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUsers(w, r)
	case http.MethodPost:
		createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// userHandler handles requests to GET, PUT, and DELETE /users/:id
func UserHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getUser(w, r)
	case http.MethodPut:
		updateUser(w, r)
	case http.MethodDelete:
		deleteUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// getUsers handles GET requests to /users
func getUsers(w http.ResponseWriter, r *http.Request) {
	userMutex.Lock()
	defer userMutex.Unlock()

	// create a channel to receive the users
	ch := make(chan *user.User, len(users))

	// loop through the users and retrieve them concurrently
	for _, u := range users {
		waitGroup.Add(1)
		go func(u *user.User) {
			defer waitGroup.Done()
			ch <- u
		}(u)
	}

	// close the channel when all goroutines are done
	go func() {
		waitGroup.Wait()
		close(ch)
	}()

	// retrieve the users from the channel and send the response
	var result []*user.User
	for u := range ch {
		result = append(result, u)
	}
	json.NewEncoder(w).Encode(result)
}

// createUser handles POST requests to /users
func createUser(w http.ResponseWriter, r *http.Request) {
	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	userID++
	user.ID = userID
	users[user.ID] = &user

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
}

// getUser handles GET requests to /users/:id
func getUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	u, ok := users[uint64(id)]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var user user.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	u, ok := users[uint64(id)]
	if !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email

	json.NewEncoder(w).Encode(u)
}

// deleteUser handles DELETE requests to /users/:id
func deleteUser(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Path[len("/users/"):])
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	userMutex.Lock()
	defer userMutex.Unlock()

	if _, ok := users[uint64(id)]; !ok {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	delete(users, uint64(id))
	w.WriteHeader(http.StatusNoContent)
}
