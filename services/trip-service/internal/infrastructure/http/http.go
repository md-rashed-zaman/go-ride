package http

import (
	"encoding/json"
	"go-ride/services/trip-service/internal/domain"
	"go-ride/shared/types"
	"log"
	"net/http"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserId      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (s *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	// validation
	if reqBody.UserId == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	ctx := r.Context()
	t, err := s.Service.GetRoute(ctx, &reqBody.Pickup, &reqBody.Destination)
	if err != nil {
		log.Println(err)
	}

	writeJSON(w, http.StatusCreated, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
