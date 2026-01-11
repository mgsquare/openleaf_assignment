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

	rateService := services.NewRateService()
	rateController := controllers.NewRateController(rateService)

	trackingRepo := repository.NewTrackingRepository()
	trackingService := services.NewTrackingService(trackingRepo)
	trackingController := controllers.NewTrackingController(trackingService)

	mux := http.NewServeMux()
	mux.HandleFunc("/api/v1/order/create", controller.CreateOrder)
	mux.HandleFunc("/api/v1/order/rates", rateController.CalculateRates)
	mux.HandleFunc("/api/v1/tracking/poll", trackingController.PollTracking)

	return mux
}
