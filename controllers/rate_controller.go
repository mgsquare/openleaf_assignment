package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mgsquare/openleaf_assignment/logger"
	"github.com/mgsquare/openleaf_assignment/models"
	"github.com/mgsquare/openleaf_assignment/services"
)

type RateController struct {
	service *services.RateService
}

func NewRateController(service *services.RateService) *RateController {
	return &RateController{service: service}
}

func (c *RateController) CalculateRates(w http.ResponseWriter, r *http.Request) {
	var req models.RateRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		logger.Error("bad request", http.ErrBodyNotAllowed)
		return
	}

	rates := c.service.CalculateRates()

	resp := models.RateResponse{
		Rates: rates,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
