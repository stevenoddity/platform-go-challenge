package favorite_route

import (
	"gwi/middleware"
	favorite_service "gwi/services/favorite"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	// Favorites routes with JWT middleware
	router.Handle("/favorites", middleware.JWTAuth(http.HandlerFunc(favorite_service.GetFavorites))).Methods("GET")
	router.Handle("/favorites", middleware.JWTAuth(http.HandlerFunc(favorite_service.AddFavorite))).Methods("POST")
	router.Handle("/favorites/{id}", middleware.JWTAuth(http.HandlerFunc(favorite_service.DeleteFavorite))).Methods("DELETE")

}
