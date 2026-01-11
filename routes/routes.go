package routes

import (
	"net/http"

	"github.com/mgsquare/openleaf_assignment/controllers"
	"github.com/mgsquare/openleaf_assignment/repository"
	"github.com/mgsquare/openleaf_assignment/services"
)

func RegisterRoutes() http.Handler {
	repo := repository.NewOrderRepository()
	service := services.NewOrderService(repo)
	controller := controllers.NewOrderController(service)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/order/create", controller.CreateOrder)

	return mux
}
