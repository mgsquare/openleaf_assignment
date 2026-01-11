package services

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	"github.com/mgsquare/openleaf_assignment/logger"
	"github.com/mgsquare/openleaf_assignment/models/tracking"
	"github.com/mgsquare/openleaf_assignment/repository"
)

type TrackingService struct {
	repo *repository.TrackingRepository
}

func NewTrackingService(repo *repository.TrackingRepository) *TrackingService {
	return &TrackingService{repo: repo}
}

func (s *TrackingService) PollTracking(trackingID string) {
	events, err := fetchCarrierTracking(trackingID)
	if err != nil {
		logger.Error("tracking poll failed for carrier", err)
		return
	}

	logger.Info(
		fmt.Sprintf(
			"tracking events fetched | tracking_id=%s | count=%d",
			trackingID,
			len(events),
		),
	)

	existing := s.repo.GetEvents(trackingID)

	var latest time.Time
	if len(existing) > 0 {
		latest = existing[len(existing)-1].Timestamp
	}

	var newEvents []tracking.TrackingEvent
	for _, e := range events {
		if e.Timestamp.After(latest) {
			newEvents = append(newEvents, e)
		}
	}

	if len(newEvents) > 0 {
		s.repo.AddEvents(trackingID, newEvents)
	}
}

func fetchCarrierTracking(trackingID string) ([]tracking.TrackingEvent, error) {
	if rand.Intn(4) == 0 {
		return nil, errors.New("carrier api failed")
	}

	now := time.Now()

	return []tracking.TrackingEvent{
		{TrackingID: trackingID, Status: "IN_TRANSIT", Timestamp: now.Add(-2 * time.Hour)},
		{TrackingID: trackingID, Status: "OUT_FOR_DELIVERY", Timestamp: now},
	}, nil
}
