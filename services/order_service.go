package services

import (
	"errors"
	"time"

	"github.com/mgsquare/openleaf_assignment/lib/helpers"
	"github.com/mgsquare/openleaf_assignment/models"
	"github.com/mgsquare/openleaf_assignment/repository"
)

type OrderService struct {
	repo *repository.OrderRepository
}

func NewOrderService(repo *repository.OrderRepository) *OrderService {
	return &OrderService{repo: repo}
}

func (s *OrderService) CreateOrder(req models.CreateOrderRequest) (models.Order, error) {

	switch req.Carrier {
	case "bluedart", "delhivery", "xpressbees":

	default:
		return models.Order{}, errors.New("unsupported carrier")
	}

	order := models.Order{
		OrderID:          req.OrderID,
		Carrier:          req.Carrier,
		TrackingID:       helpers.GenerateTrackingID(),
		Status:           models.StatusCreated,
		ExpectedDelivery: helpers.ExpectedDeliveryDate(),
		CreatedAt:        time.Now(),
	}

	if err := s.repo.Create(order); err != nil {
		return models.Order{}, err
	}

	return order, nil
}
