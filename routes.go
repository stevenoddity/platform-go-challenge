package controllers

import (
	"net/http"

	"go-web-app/handlers"
)

func RegisterRoutes() {
	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetItems(w, r)
		} else if r.Method == http.MethodPost {
			handlers.CreateItem(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/items/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			handlers.GetItem(w, r)
		} else if r.Method == http.MethodDelete {
			handlers.DeleteItem(w, r)
		} else {
			http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		}
	})
}
