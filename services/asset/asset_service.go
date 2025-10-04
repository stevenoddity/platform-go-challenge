package asset_service

import (
	asset_model "gwi/models/asset"
	"net/http"
)

var assets = []asset_model.Asset{
	{ID: 1, UserID: 1, Description: "Mock de"},
	{ID: 2, UserID: 1, Description: "Mock de"},
	{ID: 3, UserID: 2, Description: "Mock de"},
}

// GET /assets?user_id=1
func GetAssets(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// POST /assets
func AddAsset(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// DELETE /assets/{id}
func DeleteAsset(w http.ResponseWriter, r *http.Request) {
	// empty body
}

// EDIT /assets/{id}
func EditAsset(w http.ResponseWriter, r *http.Request) {
	// empty body
}
