package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mgsquare/openleaf_assignment/services"
)

type TrackingController struct {
	service *services.TrackingService
}

func NewTrackingController(service *services.TrackingService) *TrackingController {
	return &TrackingController{service: service}
}

func (c *TrackingController) PollTracking(w http.ResponseWriter, r *http.Request) {
	var body struct {
		TrackingID string `json:"tracking_id"`
	}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil || body.TrackingID == "" {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	c.service.PollTracking(body.TrackingID)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message":"tracking poll completed"}`))
}
