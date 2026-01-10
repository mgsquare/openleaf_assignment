package models

import "time"

type OrderStatus string

const (
	StatusCreated   OrderStatus = "CREATED"
	StatusShipped   OrderStatus = "SHIPPED"
	StatusInTransit OrderStatus = "IN_TRANSIT"
	StatusDelivered OrderStatus = "DELIVERED"
)

type Order struct {
	OrderID          string
	Carrier          string
	TrackingID       string
	Status           OrderStatus
	ExpectedDelivery time.Time
	CreatedAt        time.Time
}
