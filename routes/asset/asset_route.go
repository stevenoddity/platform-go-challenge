package asset_route

import (
	"gwi/middleware"
	asset_service "gwi/services/asset"
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(router *mux.Router) {

	// Assets routes with JWT middleware
	router.Handle("/assets", middleware.JWTAuth(http.HandlerFunc(asset_service.GetAssets))).Methods("GET")
	router.Handle("/assets", middleware.JWTAuth(http.HandlerFunc(asset_service.AddAsset))).Methods("POST")
	router.Handle("/assets/{id}", middleware.JWTAuth(http.HandlerFunc(asset_service.DeleteAsset))).Methods("DELETE")

}
