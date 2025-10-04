package main

import (
	"fmt"
	"gwi/routes"
	"log"
	"net/http"
)

func main() {
	routes.RegisterRoutes()
	fmt.Println("🚀 Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
