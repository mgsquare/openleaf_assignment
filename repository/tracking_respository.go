package repository

import (
	"sync"

	"github.com/mgsquare/openleaf_assignment/models/tracking"
)

type TrackingRepository struct {
	mu     sync.RWMutex
	events map[string][]tracking.TrackingEvent
}

func NewTrackingRepository() *TrackingRepository {
	return &TrackingRepository{
		events: make(map[string][]tracking.TrackingEvent),
	}
}

func (r *TrackingRepository) GetEvents(trackingID string) []tracking.TrackingEvent {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.events[trackingID]
}

func (r *TrackingRepository) AddEvents(trackingID string, newEvents []tracking.TrackingEvent) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.events[trackingID] = append(r.events[trackingID], newEvents...)
}
