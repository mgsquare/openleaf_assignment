package repository

import (
	"errors"
	"sync"

	"github.com/mgsquare/openleaf_assignment/models"
)

var ErrOrderExists = errors.New("order already exists")

type OrderRepository struct {
	mu     sync.RWMutex
	orders map[string]models.Order
}

func NewOrderRepository() *OrderRepository {
	return &OrderRepository{
		orders: make(map[string]models.Order),
	}
}

func (r *OrderRepository) Create(order models.Order) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if _, exists := r.orders[order.OrderID]; exists {
		return ErrOrderExists
	}

	r.orders[order.OrderID] = order
	return nil
}

func (r *OrderRepository) GetByID(orderID string) (models.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	order, ok := r.orders[orderID]
	return order, ok
}
