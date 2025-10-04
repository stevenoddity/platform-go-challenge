package routes

import (
	asset_route "gwi/routes/asset"
	favorite_route "gwi/routes/favorite"
	user_route "gwi/routes/user"

	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	// create one main router
	router := mux.NewRouter()

	user_route.RegisterRoutes(router)
	asset_route.RegisterRoutes(router)
	favorite_route.RegisterRoutes(router)

	return router
}
