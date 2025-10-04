package asset_route

import (
	asset_service "gwi/services/asset"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			asset_service.GetAssets(w, r)
		} else if r.Method == http.MethodPost {
			asset_service.AddAsset(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/assets/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			asset_service.DeleteAsset(w, r)
		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})
}
