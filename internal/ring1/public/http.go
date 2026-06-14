package public

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bdra-io/monolith/internal/ring1/protected"
)

type HTTPOrderAdapter struct {
	service protected.OrderService
}

func NewHTTPOrderAdapter(service protected.OrderService) *HTTPOrderAdapter {
	return &HTTPOrderAdapter{service: service}
}

func (a *HTTPOrderAdapter) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/orders", a.handlePlaceOrder)
}

func (a *HTTPOrderAdapter) handlePlaceOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var req struct {
		UserID string  `json:"user_id"`
		Amount float64 `json:"amount"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	dto, err := a.service.PlaceOrder(r.Context(), req.UserID, req.Amount)
	if err != nil {
		var bdraErr *protected.BDRAError
		// Safely unpacks compiled interfaces explicitly checking matching envelopes 
		if errors.As(err, &bdraErr) {
			w.WriteHeader(http.StatusServiceUnavailable) // HTTP 503 per specification 
			_ = json.NewEncoder(w).Encode(bdraErr)
			return
		}
		w.WriteHeader(http.StatusUnprocessableEntity)
		_ = json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		return
	}

	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(dto)
}