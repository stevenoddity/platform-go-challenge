package database

import (
	asset_model "gwi/models/asset"
	favorite_model "gwi/models/favorite"
	user_model "gwi/models/user"
)

var AssetsDB = []asset_model.Asset{
	{ID: 1, UserID: 1, Description: "Mock de"},
	{ID: 2, UserID: 1, Description: "Mock de"},
	{ID: 3, UserID: 2, Description: "Mock de"},
}

var UsersDB = []user_model.User{
	{ID: 1, Email: "test@test.com"},
	{ID: 2, Email: "gwi@gwi.com"},
}

var FavoritesDB = []favorite_model.Favorite{
	{ID: 1, UserID: 1, Asset: "BTC"},
	{ID: 2, UserID: 1, Asset: "AAPL"},
}
