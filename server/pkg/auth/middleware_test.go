package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/dsypasit/social-clone/server/internal/auth"
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware_MissingAuthorization(t *testing.T) {
	// Define a dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Create a middleware instance with a mock JwtService
	mockService := auth.NewJwtService("secret")
	middleware := AuthMiddleware(mockService)

	// Create a request without authorization header
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nilf(t, err, "Unexpected error: %v", err)

	// Record the response
	rec := httptest.NewRecorder()
	middleware(nextHandler).ServeHTTP(rec, req)

	// Assertions
	assert.Equalf(t, rec.Code, http.StatusUnauthorized, "Expected unauthorized status code, got: %v", rec.Code)

	expected := map[string]string{"message": "Missing authorization token"}
	var response map[string]string
	err = json.NewDecoder(rec.Body).Decode(&response)

	assert.Nil(t, err, fmt.Sprintf("Unexpected error decoding response: %v", err))
	assert.Equal(t, expected, response, fmt.Sprintf("Expected error message: %v, got: %v", expected, response))
}

func TestAuthMiddleware_InvalidFormat(t *testing.T) {
	// Define a dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})

	// Create a middleware instance with a mock JwtService
	mockService := auth.NewJwtService("secret")
	middleware := AuthMiddleware(mockService)

	// Create a request without authorization header
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nilf(t, err, "Unexpected error: %v", err)
	req.Header.Add("Authorization", "Invalid token format")

	// Record the response
	rec := httptest.NewRecorder()
	middleware(nextHandler).ServeHTTP(rec, req)

	// Assertions
	assert.Equalf(t, rec.Code, http.StatusUnauthorized, "Expected unauthorized status code, got: %v", rec.Code)

	expected := map[string]string{"message": "Invalid authorization format"}
	var response map[string]string
	err = json.NewDecoder(rec.Body).Decode(&response)

	assert.Nil(t, err, fmt.Sprintf("Unexpected error decoding response: %v", err))
	assert.Equal(t, expected, response, fmt.Sprintf("Expected error message: %v, got: %v", expected, response))
}

type MockJwtService struct {
	claimToken auth.AuthJWTClaim
}

func (m *MockJwtService) GenerateToken(userUUID string) (string, error) {
	return "", nil
}

func (m *MockJwtService) VerifyToken(token string) (*auth.AuthJWTClaim, error) {
	return &m.claimToken, nil
}

func TestAuthMiddleware_ValidClaims(t *testing.T) {
	claimToken := auth.AuthJWTClaim{UserUUID: "94d67127-78e8-419f-adb8-782d26e4805d", RegisteredClaims: jwt.RegisteredClaims{}}
	// Create a middleware instance with a mock JwtService
	mockService := MockJwtService{
		claimToken: claimToken,
	}

	middleware := AuthMiddleware(&mockService)

	// Create a request without authorization header
	req, err := http.NewRequest(http.MethodGet, "/", nil)
	assert.Nilf(t, err, "Unexpected error: %v", err)
	req.Header.Add("Authorization", "Bearer valid token")

	// Record the response
	rec := httptest.NewRecorder()
	// Define a dummy handler
	nextHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		expected := claimToken.UserUUID

		ctx := r.Context()
		actual, _ := ctx.Value("userUUID").(string)

		assert.Equal(t, expected, actual, fmt.Sprintf("Expected user uuid: %v, got: %v", expected, actual))

		w.WriteHeader(http.StatusOK)
	})

	middleware(nextHandler).ServeHTTP(rec, req)

	assert.Equalf(t, http.StatusOK, rec.Code, "Expected Ok status code, got: %v", rec.Code)
}
