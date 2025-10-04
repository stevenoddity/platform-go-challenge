package routes

import (
	asset_route "gwi/routes/asset"
	favorite_route "gwi/routes/favorite"
	user_route "gwi/routes/user"
)

func RegisterRoutes() {
	user_route.RegisterRoutes()
	asset_route.RegisterRoutes()
	favorite_route.RegisterRoutes()
}
