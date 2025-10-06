package utils_test

import (
	"errors"
	"gwi/configuration"
	"gwi/utils"
	"testing"

	"github.com/dgrijalva/jwt-go"
)

func GenerateJwtToken(userID int) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": float64(userID),
	})
	tokenString, _ := token.SignedString(configuration.JwtSecret)
	return tokenString
}

func TestExtractUserID_ValidToken(t *testing.T) {
	userID := 1
	authorizationHeader := "Bearer " + GenerateJwtToken(userID)

	extractedID, err := utils.ExtractUserID(authorizationHeader)

	// Assert
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if extractedID != userID {
		t.Errorf("expected user_id %d, got %d", userID, extractedID)
	}
}

func TestExtractUserID_InvalidToken(t *testing.T) {
	authorizationHeader := "Bearer invalid-token"

	extractedID, err := ExtractUserID(authorizationHeader)

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got none")
	}
	if extractedID != 0 {
		t.Errorf("expected user_id 0, got %d", extractedID)
	}
}

func TestExtractUserID_NoUserIDInClaims(t *testing.T) {
	// Arrange
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	tokenString, _ := token.SignedString(configuration.JwtSecret)

	authorizationHeader := "Bearer " + tokenString

	extractedID, err := ExtractUserID(authorizationHeader)

	// Assert
	if err == nil {
		t.Fatalf("expected an error, got none")
	}
	if !errors.Is(err, errors.New("user_id not found in token")) {
		t.Errorf("expected user_id not found error, got %v", err)
	}
	if extractedID != 0 {
		t.Errorf("expected user_id 0, got %d", extractedID)
	}
}
func TestIsUserAuthorized_ValidAuthorization(t *testing.T) {
	userID := 1
	authorizationHeader := "Bearer " + GenerateJwtToken(userID)

	authorized := IsUserAuthorized(userID, authorizationHeader)

	// Assert
	if !authorized {
		t.Errorf("expected user to be authorized, got %v", authorized)
	}
}

func TestIsUserAuthorized_InvalidAuthorization(t *testing.T) {
	userID := 1
	authorizationHeader := "Bearer invalid-token"

	authorized := IsUserAuthorized(userID, authorizationHeader)

	// Assert
	if authorized {
		t.Errorf("expected user to be unauthorized, got %v", authorized)
	}
}

func TestIsUserAuthorized_NoUserIDInClaims(t *testing.T) {
	userID := 1
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	tokenString, _ := token.SignedString(configuration.JwtSecret)
	authorizationHeader := "Bearer " + tokenString

	authorized := IsUserAuthorized(userID, authorizationHeader)

	// Assert
	if authorized {
		t.Errorf("expected user to be unauthorized, got %v", authorized)
	}
}
