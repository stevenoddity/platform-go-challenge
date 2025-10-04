package user_service

import (
	user_model "gwi/models/user"
	"net/http"
)

var users = []user_model.User{
	{ID: 1, Email: "test@test.com"},
	{ID: 2, Email: "gwi@gwi.com"},
}

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
