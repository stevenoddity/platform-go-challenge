package user_service

import (
	"net/http"
)

// GET /users?user_id=1
func GetUsers(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// POST /users
func AddUser(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// DELETE /users/{id}
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// EDIT /users/{id}
func EditUser(w http.ResponseWriter, r *http.Request) {
	// empty body
}
