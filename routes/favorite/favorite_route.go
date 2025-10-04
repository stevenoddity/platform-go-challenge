package favorite_route

import (
	favorite_service "gwi/services/favorite"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/favorites", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			favorite_service.GetFavorites(w, r)
		} else if r.Method == http.MethodPost {
			favorite_service.AddFavorite(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/favorites/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			favorite_service.DeleteFavorite(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
