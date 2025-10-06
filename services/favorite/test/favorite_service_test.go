package favorite_service_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"gwi/configuration"
	"gwi/constants"
	favorite_service "gwi/services/favorite"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(userID),
	})
	tokenString, _ := token.SignedString(configuration.JwtSecret)
	return tokenString
}

func TestAddFavorite_ValidRequest(t *testing.T) {
	userID := 1
	authorizationHeader := "Bearer " + GenerateJwtToken(userID)

	body := bytes.NewBufferString(`{"asset_id": 3}`)

	req := httptest.NewRequest(http.MethodPost, "/"+constants.ENDPOINT_FAVORITES, body)
	req.Header.Set("Authorization", authorizationHeader)
	w := httptest.NewRecorder()

	favorite_service.AddFavorite(w, req)
	resp := w.Result()
	defer resp.Body.Close()

	// Assert
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200 Created, got %d", resp.StatusCode)
	}

}

func TestAddFavorite_InvalidToken(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/"+constants.ENDPOINT_FAVORITES, bytes.NewBufferString(`{"asset_id": 1}`))
	req.Header.Set("Authorization", "Bearer invalid-token")
	w := httptest.NewRecorder()

	favorite_service.AddFavorite(w, req)

	// Assert
	if w.Code != http.StatusForbidden {
		t.Errorf("expected 401, got %d", w.Code)
	}
}

func TestAddFavorite_AssetNotFound(t *testing.T) {
	userID := 1
	authorizationHeader := "Bearer " + GenerateJwtToken(userID)
	body := bytes.NewBufferString(`{"asset_id": 999}`)

	req := httptest.NewRequest(http.MethodPost, "/"+constants.ENDPOINT_FAVORITES, body)
	req.Header.Set("Authorization", authorizationHeader)
	w := httptest.NewRecorder()

	favorite_service.AddFavorite(w, req)

	// Assert
	if w.Code != http.StatusNotFound {
		t.Errorf("expected 404, got %d", w.Code)
	}
}
