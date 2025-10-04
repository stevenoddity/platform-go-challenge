package favorite_service

import (
	"encoding/json"
	"gwi/database"
	favorite_model "gwi/models/favorite"
	"net/http"
	"strconv"
)

// GET /favorites?user_id=1
func GetFavorites(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")

	if userIDStr == "" {
		http.Error(w, "Missing user_id", http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		http.Error(w, "Invalid user_id", http.StatusBadRequest)
		return
	}

	var userFavorites []favorite_model.Favorite
	for _, f := range database.FavoritesDB {
		if f.UserID == userID {
			userFavorites = append(userFavorites, f)
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(userFavorites)
}

// POST /favorites
func AddFavorite(w http.ResponseWriter, r *http.Request) {
	var newFav favorite_model.Favorite
	if err := json.NewDecoder(r.Body).Decode(&newFav); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Assign new ID
	newFav.ID = len(database.FavoritesDB) + 1
	database.FavoritesDB = append(database.FavoritesDB, newFav)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newFav)
}

// DELETE /favorites/{id}
func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/favorites/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	for i, f := range database.FavoritesDB {
		if f.ID == id {
			database.FavoritesDB = append(database.FavoritesDB[:i], database.FavoritesDB[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	http.Error(w, "Favorite not found", http.StatusNotFound)
}
