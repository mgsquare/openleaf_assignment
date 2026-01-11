package services

import (
	"errors"
	"math/rand"
	"time"

	"github.com/mgsquare/openleaf_assignment/logger"
	"github.com/mgsquare/openleaf_assignment/models"
)

type RateResult struct {
	Carrier     string
	Rate        float64
	Serviceable bool
	Err         error
}

type RateService struct{}

func NewRateService() *RateService {
	return &RateService{}
}

func (s *RateService) CalculateRates() []models.CarrierRate {
	carriers := []string{"bluedart", "delhivery", "xpressbees"}

	results := make(chan RateResult, len(carriers))

	for _, c := range carriers {
		go func(carrier string) {
			results <- fetchCarrierRate(carrier)
		}(c)
	}

	var rates []models.CarrierRate

	for i := 0; i < len(carriers); i++ {
		res := <-results

		if res.Err != nil {
			continue
		}

		if !res.Serviceable {
			continue
		}

		rates = append(rates, models.CarrierRate{
			Carrier: res.Carrier,
			Rate:    res.Rate,
		})
	}

	return rates
}

func fetchCarrierRate(carrier string) RateResult {
	time.Sleep(time.Duration(rand.Intn(300)) * time.Millisecond)

	n := rand.Intn(3)

	switch n {
	case 0:
		logger.Info("carrier rates received")
		return RateResult{
			Carrier:     carrier,
			Rate:        float64(rand.Intn(200) + 50),
			Serviceable: true,
		}
	case 1:
		logger.Error("no response recieved from carrier api", errors.New("Carrier api down"))
		return RateResult{
			Carrier: carrier,
			Err:     errors.New("carrier api failed"),
		}
	default:
		return RateResult{
			Carrier:     carrier,
			Serviceable: false,
		}
	}
}
