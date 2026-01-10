package models

type CreateOrderResponse struct {
	OrderID    string `json:"order_id"`
	TrackingID string `json:"tracking_id"`
	Status     string `json:"status"`
}
