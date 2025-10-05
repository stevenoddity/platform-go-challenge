package user_route

import (
	"gwi/middleware"
	user_service "gwi/services/user"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the routes for user management with JWT authentication.
func RegisterRoutes(router *mux.Router) {
	router.Handle("/users", middleware.JWTAuth(http.HandlerFunc(user_service.GetUsers))).Methods("GET")
	router.Handle("/users", middleware.JWTAuth(http.HandlerFunc(user_service.AddUser))).Methods("POST")
	router.Handle("/users/{id}", middleware.JWTAuth(http.HandlerFunc(user_service.DeleteUser))).Methods("DELETE")
}
