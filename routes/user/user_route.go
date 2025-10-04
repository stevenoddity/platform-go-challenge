package user_route

import (
	user_service "gwi/services/user"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			user_service.GetUsers(w, r)
		} else if r.Method == http.MethodPost {
			user_service.AddUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			user_service.DeleteUser(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
