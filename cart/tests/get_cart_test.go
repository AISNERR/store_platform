package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestListCartItems(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userID := r.URL.Query().Get("user_id")
		if userID == "" {
			http.Error(w, "Missing user_id", http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, `[{"sku_id": 100, "count": 2}, {"sku_id": 101, "count": 1}]`)
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/cart/list?user_id=1", ts.URL), nil)
	assert.Equal(t, err, nil)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)
	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
	body, err := io.ReadAll(resp.Body)
	assert.Equal(t, err, nil)

	expectedBody := `[{"sku_id": 100, "count": 2}, {"sku_id": 101, "count": 1}]`
	assert.Equal(t, expectedBody, string(body))
}
