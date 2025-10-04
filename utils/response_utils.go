package utils

import (
	"encoding/json"
	"net/http"
)

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func SendSuccess(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(APIResponse{Success: true, Data: data})
}

func SendError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")

	if apiErr, ok := err.(*APIError); ok {
		w.WriteHeader(apiErr.Status)
		json.NewEncoder(w).Encode(APIResponse{Success: false, Message: apiErr.Message})
		return
	}

	// fallback for unknown errors
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(APIResponse{Success: false, Message: "Internal server error"})
}
