package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type GetCartCartResponse struct {
	SKU   int64   `json:"sku"`
	Count int64   `json:"count"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type GetCartResponse struct {
	Items      []GetCartCartResponse `json:"items"`
	TotalPrice float64               `json:"total_price"`
}

func (s *Server) GetGood(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	cartInfo, err := s.cartService.GetCart(r.Context(), userID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("GET /user/{user_id}/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	rawResponse, err := json.Marshal(cartInfo)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("GET /user/{user_id}/cart out failed: %s", errOut.Error())
			return
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	w.Write(rawResponse)
}
