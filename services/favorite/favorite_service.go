package favorite_service

import (
	"encoding/json"
	asset_model "gwi/models/asset"
	favorite_model "gwi/models/favorite"
	"net/http"
	"strconv"

	"gwi/database"
	"gwi/utils"
)

// GetFavorites handles the HTTP request to retrieve the favorite items of a user.
// It expects a query parameter "user_id" to identify the user. If the user_id is missing
// or invalid, it responds with a bad request error. If valid, it searches the database
// for the user's favorites and returns them in the response.
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
	for _, fav := range database.FavoritesDB {
		if fav.UserID == userID {
			for _, a := range database.AssetsDB {
				if a.ID == fav.AssetID {
					fav.Asset = &a
					break
				}
			}
			userFavorites = append(userFavorites, fav)
		}
	}

	utils.SendSuccess(w, userFavorites)
}

// AddFavorite handles the HTTP request to add a new favorite item.
// It decodes the JSON request body into a Favorite model, validates it,
// assigns a new ID, and appends it to the FavoritesDB. If the JSON is
// invalid or the favorite item fails validation, it sends an error response.
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

	// Check if asset exists
	var assetFound *asset_model.Asset
	for _, a := range database.AssetsDB {
		if a.ID == newFav.AssetID {
			assetFound = &a
			break
		}
	}
	if assetFound == nil {
		utils.SendError(w, utils.ErrNotFound("Asset not found"))
		return
	}
	newFav.ID = len(database.FavoritesDB) + 1
	database.FavoritesDB = append(database.FavoritesDB, newFav)

	utils.SendSuccess(w, newFav)
}

// DeleteFavorite handles the HTTP request to delete a favorite item.
// It extracts the ID from the URL, converts it to an integer, and removes
// the corresponding favorite from the database. If the ID is invalid or
// the favorite is not found, it sends an appropriate error response.
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
