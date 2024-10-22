package server

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestDeleteCartItem(t *testing.T) {

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.Method != http.MethodDelete {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		userID := r.URL.Query().Get("user_id")
		skuID := r.URL.Query().Get("sku_id")
		if userID == "" || skuID == "" {
			http.Error(w, "Missing user_id or sku_id", http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "{}")
	}))
	defer ts.Close()

	req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("%s/cart/item/delete?user_id=1&sku_id=100", ts.URL), nil)
	assert.Equal(t, err, nil)

	resp, err := http.DefaultClient.Do(req)
	assert.Equal(t, err, nil)

	defer resp.Body.Close()

	assert.Equal(t, resp.StatusCode, http.StatusOK)
}
