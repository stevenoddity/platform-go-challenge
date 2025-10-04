package middleware

import (
	"log"
	"net/http"

	"gwi/utils"
)

func RecoverMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("Recovered from panic: %v", err)
				utils.SendError(w, utils.ErrInternalServer("Internal server error"))
			}
		}()

		next.ServeHTTP(w, r)
	})
}
