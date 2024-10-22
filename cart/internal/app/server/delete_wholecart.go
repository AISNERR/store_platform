package server

import (
	"fmt"
	"log"
	"net/http"
	"route256/cart/internal/pkg/reviews/model"
	"strconv"
)

func (s *Server) DeleteCart(w http.ResponseWriter, r *http.Request) {
	rawID := r.PathValue("user_id")
	userID, err := strconv.ParseInt(rawID, 10, 64)
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = s.cartService.DelWholeCart(r.Context(), model.UserID(userID))

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Header().Set("Content-Type", "application/json")
		_, errOut := fmt.Fprintf(w, "{\"message\":\"%s\"}", err)
		if errOut != nil {
			log.Printf("DELETE /user/{user_id}/cart out failed: %s", errOut.Error())
			return
		}

		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "{}")
}
