package favorite_service

import (
	"encoding/json"
	favorite_model "gwi/models/favorite"
	"net/http"
	"strconv"

	"gwi/database"
	"gwi/utils"
)

func GetFavorites(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.URL.Query().Get("user_id")
	if userIDStr == "" {
		utils.SendError(w, utils.ErrBadRequest("Missing user_id"))
		return
	}

	userID, err := strconv.Atoi(userIDStr)
	if err != nil {
		utils.SendError(w, utils.ErrBadRequest("Invalid user_id"))
		return
	}

	var userFavorites []favorite_model.Favorite
	for _, f := range database.FavoritesDB {
		if f.UserID == userID {
			userFavorites = append(userFavorites, f)
		}
	}

	utils.SendSuccess(w, userFavorites)
}

func AddFavorite(w http.ResponseWriter, r *http.Request) {
	var newFav favorite_model.Favorite
	if err := json.NewDecoder(r.Body).Decode(&newFav); err != nil {
		utils.SendError(w, utils.ErrBadRequest("Invalid JSON"))
		return
	}

	if err := utils.ValidateFavorite(&newFav); err != nil {
		utils.SendError(w, utils.ErrBadRequest(err.Error()))
		return
	}

	newFav.ID = len(database.FavoritesDB) + 1
	database.FavoritesDB = append(database.FavoritesDB, newFav)

	utils.SendSuccess(w, newFav)
}

func DeleteFavorite(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/favorites/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		utils.SendError(w, utils.ErrBadRequest("Invalid ID"))
		return
	}

	for i, f := range database.FavoritesDB {
		if f.ID == id {
			database.FavoritesDB = append(database.FavoritesDB[:i], database.FavoritesDB[i+1:]...)
			utils.SendSuccess(w, nil)
			return
		}
	}

	utils.SendError(w, utils.ErrNotFound("Favorite not found"))
}
