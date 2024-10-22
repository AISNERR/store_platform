package server

import (
	"fmt"
	"log"
	"net/http"
	"route256/cart/internal/pkg/reviews/model"
	"strconv"
)

func (s *Server) DeleteGood(w http.ResponseWriter, r *http.Request) {
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
	if sku < 1 {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", "sku must be more than 0")
		if errOut != nil {
			log.Printf("POST /user/{user_id}/cart/{sku_id} out failed: %s", errOut.Error())
			return
		}
		return
	}

	delete_info := model.RequestDelete{
		SKU:    sku,
		UserID: userID,
	}

	err = s.cartService.DelCart(r.Context(), delete_info.UserID, delete_info.SKU)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("DELETE /user/{user_id}/cart/{sku_id} out failed: %s", errOut.Error())
			return
		}

		return
	}

	log.Println("Goods are deleted", delete_info)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "{}")
}
