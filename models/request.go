package models

type CreateOrderRequest struct {
	OrderID string `json:"order_id"`
	Carrier string `json:"carrier"`
	Pincode string `json:"pincode"`
}

type RateRequest struct {
	Pincode string `json:"pincode"`
	Weight  int    `json:"weight"`
}
