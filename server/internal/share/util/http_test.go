package util

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendJson(t *testing.T) {
	testTable := []struct {
		title           string
		input           interface{}
		statusCode      int
		wantErr         error
		wantContentType string
	}{
		{"should return json string", map[string]string{"wowo": "abcd"}, http.StatusOK, nil, "application/json"},
		{"should return json string", struct {
			name string `json:"name"`
		}{"hello"}, http.StatusCreated, nil, "application/json"},
	}

	for _, v := range testTable {
		t.Run(v.title, func(t *testing.T) {
			expected, err := json.Marshal(v.input)
			if err != nil {
				t.Errorf("Unexpected error marshalling input: %v", err)
				return
			}

			rec := httptest.NewRecorder()
			r, _ := http.NewRequest(http.MethodGet, "/", nil)

			h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				SendJson(w, v.input, v.statusCode) // Call SendJson with ResponseWriter argument
			})

			h.ServeHTTP(rec, r) // Use ServeHTTP for proper request handling

			assert.Equal(t, expected, rec.Body.Bytes(), "Want %v but got %v", expected, rec.Body.Bytes())
			assert.Equal(t, v.statusCode, rec.Code, "Want %v but got %v", expected, rec.Body.Bytes())
			assert.Equal(t, v.wantContentType, rec.Header().Get("Content-Type"), "Want %v but got %v", v.wantContentType, rec.Header().Get("Content-Type"))
		})
	}
}