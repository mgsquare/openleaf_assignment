package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/mgsquare/openleaf_assignment/logger"
	"github.com/mgsquare/openleaf_assignment/models"
	"github.com/mgsquare/openleaf_assignment/services"
)

type OrderController struct {
	service *services.OrderService
}

func NewOrderController(service *services.OrderService) *OrderController {
	return &OrderController{service: service}
}

func (c *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req models.CreateOrderRequest

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		logger.Error("Error parsing request. try again", err)
		return
	}

	if req.OrderID == "" || req.Carrier == "" {
		http.Error(w, "missing mandatory fields", http.StatusBadRequest)
		logger.Error("Incomplete request. try again", http.ErrBodyNotAllowed)
		return
	}

	order, err := c.service.CreateOrder(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		logger.Error("exception while creating order", err)
		return
	}

	resp := models.CreateOrderResponse{
		OrderID:    order.OrderID,
		TrackingID: order.TrackingID,
		Status:     string(order.Status),
	}
	logger.Info("order created successfully") // make logs more verbose - include order id in the msg.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
