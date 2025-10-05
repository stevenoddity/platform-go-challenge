package asset_route

import (
	"gwi/middleware"
	asset_service "gwi/services/asset"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes sets up the routes for asset management with JWT authentication.
func RegisterRoutes(router *mux.Router) {
	router.Handle("/assets", middleware.JWTAuth(http.HandlerFunc(asset_service.GetAssets))).Methods("GET")
	router.Handle("/assets", middleware.JWTAuth(http.HandlerFunc(asset_service.AddAsset))).Methods("POST")
	router.Handle("/assets/{id}", middleware.JWTAuth(http.HandlerFunc(asset_service.EditAsset))).Methods("PUT")
	router.Handle("/assets/{id}", middleware.JWTAuth(http.HandlerFunc(asset_service.EditAsset))).Methods("PUT")
}
