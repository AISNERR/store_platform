package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"route256/cart/internal/pkg/reviews/model"
	"strconv"
)

type CreateCartRequest struct {
	Count int64 `json:"count"`
}

type CreateCartResponse struct {
	Status string `json:"status"`
}

func (s *Server) CreateGood(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	skuID := r.PathValue("sku_id")
	sku, err := strconv.ParseInt(skuID, 10, 64)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/{id}/cart got error %v", errOut.Error())
			return
		}

		return
	}

	body, err := io.ReadAll(r.Body)

	var createRequest CreateCartRequest

	err = json.Unmarshal(body, &createRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/{user_id}/cart/{sku_id} out failed: %s", errOut.Error())
			return
		}

		return
	}

	inputCart := model.Cart{
		SKU:    sku,
		UserID: userID,
		Count:  createRequest.Count,
	}

	if err := s.validator.Struct(inputCart); err != nil {
		http.Error(w, fmt.Sprintf(`{"message": %s}`, err.Error()), http.StatusBadRequest)
		return
	}

	cartOutput, err := s.cartService.AddCart(r.Context(), inputCart)
	if err == model.ErrNoProductSku {
		w.WriteHeader(http.StatusPreconditionFailed)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/{user_id}/cart/{sku_id} out failed: %s", errOut.Error())
			return
		}
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("POST /user/{user_id}/cart/{sku_id} out failed: %s", errOut.Error())
			return
		}
		return
	}

	w.WriteHeader(http.StatusOK)
	log.Println("Goods are saved", cartOutput)
	fmt.Fprint(w, "{}")
}
