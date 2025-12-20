package main

import (
	"encoding/json"
	"go-ride/shared/contracts"
	"net/http"
	"time"
)

func HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second * 9)
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "failed to parse JSON data", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	// validation
	if reqBody.UserId == "" {
		http.Error(w, "user ID is required", http.StatusBadRequest)
		return
	}

	// TODO: Call trip service

	response := contracts.APIResponse{Data: "ok"}

	writeJSON(w, http.StatusCreated, response)
}
